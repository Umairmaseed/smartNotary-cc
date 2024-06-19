package txdefs

import (
	"encoding/json"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var PayInstallment = tx.Transaction{
	Tag:         "payInstallment",
	Label:       "Installment payment",
	Description: "Installment payment",
	Method:      "POST",

	Args: []tx.Argument{
		{
			Required: true,
			Tag:      "assetId",
			Label:    "Contract Id",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "installment",
			Label:    "Installment",
			DataType: "@object",
		},
		{
			Required: true,
			Tag:      "payment",
			Label:    "Payment",
			DataType: "@object",
		},
	},

	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		assetId, ok := req["assetId"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'assetId'")
		}

		contractKey, err := assets.NewKey(map[string]interface{}{
			"@assetType": "contract",
			"id":         assetId,
		})
		if err != nil {
			return nil, errors.WrapError(err, "failed to create key for contract")
		}

		contractAsset, err := contractKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to get contract asset from the ledger")
		}

		// Extract key values from installment
		installment, ok := req["installment"].(map[string]interface{})
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment'")
		}

		payment, ok := req["payment"].(map[string]interface{})
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'payment'")
		}

		installmentId, ok := installment["id"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment.id'")
		}

		installmentNumber, ok := installment["number"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment.number'")
		}

		dueDateUtc, ok := installment["dueDateUtc"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment.dueDateUtc'")
		}

		amortization, ok := installment["amortization"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment.amortization'")
		}

		financingModalityInterestValue, ok := installment["financingModalityInterestValue"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment.financingModalityInterestValue'")
		}

		installmentValue, ok := installment["installmentValue"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment.installmentValue'")
		}

		fineInCaseOfDelayValue, ok := installment["fineInCaseOfDelayValue"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment.fineInCaseOfDelayValue'")
		}

		interestInCaseOfDelayValue, ok := installment["interestInCaseOfDelayValue"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment.interestInCaseOfDelayValue'")
		}

		monetaryCorrectionValue, ok := installment["monetaryCorrectionValue"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment.monetaryCorrectionValue'")
		}

		reimbursement, ok := installment["reimbursement"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment.reimbursement'")
		}

		// Extract key values from payment
		paymentValue, ok := payment["value"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'payment.value'")
		}

		dateCompletedUtc, ok := payment["dateCompletedUtc"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'payment.dateCompletedUtc'")
		}

		discriminator, ok := payment["discriminator"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'payment.discriminator'")
		}

		gatewaySaleId, ok := payment["gatewaySaleId"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'payment.gatewaySaleId'")
		}

		gatewaySellerId, ok := payment["gatewaySellerId"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'payment.gatewaySellerId'")
		}

		gatewayFee, ok := payment["gatewayFee"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'payment.gatewayFee'")
		}

		keyInstallment, err := assets.NewKey(map[string]interface{}{
			"@assetType": "installment",
			"id":         installmentId,
		})
		if err != nil {
			return nil, errors.WrapError(err, "failed to create key")
		}

		installmentAsset, err := keyInstallment.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to get installment asset from the ledger")
		}

		// Registering payment asset

		paymentMap := map[string]interface{}{
			"@assetType":       "payment",
			"installment":      keyInstallment,
			"value":            paymentValue,
			"dateCompletedUtc": dateCompletedUtc,
			"discriminator":    discriminator,
			"gatewaySaleId":    gatewaySaleId,
			"gatewaySellerId":  gatewaySellerId,
			"gatewayFee":       gatewayFee,
			"reimbursement":    reimbursement,
		}
		paymentAsset, err := assets.NewAsset(paymentMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new payment asset")
		}

		_, err = paymentAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Updating installment asset
		updateInstallmentMap := (map[string]interface{})(*installmentAsset)
		updateInstallmentMap["number"] = installmentNumber
		updateInstallmentMap["dueDateUtc"] = dueDateUtc
		updateInstallmentMap["amortization"] = amortization
		updateInstallmentMap["financingModalityInterestValue"] = financingModalityInterestValue
		updateInstallmentMap["installmentValue"] = installmentValue
		updateInstallmentMap["fineInCaseOfDelayValue"] = fineInCaseOfDelayValue
		updateInstallmentMap["interestInCaseOfDelayValue"] = interestInCaseOfDelayValue
		updateInstallmentMap["monetaryCorrectionValue"] = monetaryCorrectionValue
		updateInstallmentMap["reimbursement"] = reimbursement

		_, err = installmentAsset.Update(stub, updateInstallmentMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to update installment asset")
		}

		// update contract asset
		var paymentKeyMap []interface{}
		if existingPayments, exists := (*contractAsset)["payment"].([]interface{}); exists {
			paymentKeyMap = append(paymentKeyMap, existingPayments...)
		}
		paymentKeyMap = append(paymentKeyMap, map[string]interface{}{"@key": paymentAsset["@key"]})
		updateContractMap := (map[string]interface{})(*contractAsset)
		updateContractMap["payment"] = paymentKeyMap

		_, err = contractAsset.Update(stub, updateContractMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to update contract asset")
		}

		paymentJSON, nerr := json.Marshal(paymentAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "Failed to encode asset to JSON format")
		}

		return paymentJSON, nil

	},
}
