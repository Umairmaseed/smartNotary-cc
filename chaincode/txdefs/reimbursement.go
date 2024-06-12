package txdefs

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var Reimbursement = tx.Transaction{
	Tag:         "reimbursement",
	Label:       "Create Reimbursement",
	Description: "Create Reimbursement",
	Method:      "POST",

	Args: []tx.Argument{
		{
			Required: true,
			Tag:      "assetId",
			Label:    "Contract ID",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "cancelledInstallments",
			Label:    "Cancelled installments",
			DataType: "[]@object",
		},
		{
			Required: true,
			Tag:      "reimbursementInstallment",
			Label:    "Reimbursement installment",
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

		cancelledInstallments, ok := req["cancelledInstallments"].([]interface{})
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'cancelledInstallments'")
		}

		reimbursementInstallment, ok := req["reimbursementInstallment"].(map[string]interface{})
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'reimbursementInstallment'")
		}

		// extracting reimbursement keys

		reimbursementId, ok := reimbursementInstallment["id"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment.id'")
		}
		reimbursementNumber, ok := reimbursementInstallment["number"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment.number'")
		}
		reimbursementDueDateUtc, ok := reimbursementInstallment["dueDateUtc"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment.dueDateUtc'")
		}
		reimbursementAmortization, ok := reimbursementInstallment["amortization"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment.amortization'")
		}

		reimbursementFinancingModalityInterestValue, ok := reimbursementInstallment["financingModalityInterestValue"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment.financingModalityInterestValue'")
		}

		reimbursementInstallmentValue, ok := reimbursementInstallment["installmentValue"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment.installmentValue'")
		}

		reimbursementMap := map[string]interface{}{
			"@assetType":                     "installment",
			"id":                             reimbursementId,
			"number":                         reimbursementNumber,
			"dueDateUtc":                     reimbursementDueDateUtc,
			"amortization":                   reimbursementAmortization,
			"financingModalityInterestValue": reimbursementFinancingModalityInterestValue,
			"installmentValue":               reimbursementInstallmentValue,
		}
		reimbursementInstallmentAsset, err := assets.NewAsset(reimbursementMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new reimbursement installment asset")
		}

		RIAsset, err := reimbursementInstallmentAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		var cancelledInstallmentAssets []interface{}

		for _, installment := range cancelledInstallments {
			cancelledInstallment, ok := installment.(map[string]interface{})
			if !ok {
				return nil, errors.WrapError(nil, "Invalid type for cancelledInstallment")
			}

			cancelledInstallmentId, ok := cancelledInstallment["id"].(string)
			if !ok {
				return nil, errors.WrapError(nil, "Invalid type for parameter 'cancelledInstallment.id'")
			}

			keyInstallment, err := assets.NewKey(map[string]interface{}{
				"@assetType": "installment",
				"id":         cancelledInstallmentId,
			})
			if err != nil {
				return nil, errors.WrapError(err, "failed to create key")
			}

			installmentAsset, err := keyInstallment.Get(stub)
			if err != nil {
				return nil, errors.WrapError(err, "Failed to get installment asset from the ledger")
			}

			key := (*installmentAsset)["@key"].(string)
			cancelledInstallmentAssets = append(cancelledInstallmentAssets, map[string]interface{}{"@key": key})
		}

		fmt.Println("---------------", cancelledInstallmentAssets)

		reimburmentAssetMap := map[string]interface{}{
			"@assetType":               "reimbursement",
			"cancelledInstallments":    cancelledInstallmentAssets,
			"reimbursementInstallment": RIAsset,
		}

		reimbursementAsset, err := assets.NewAsset(reimburmentAssetMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new reimbursement asset")
		}

		RAsset, err := reimbursementAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}
		updateContractMap := (map[string]interface{})(*contractAsset)
		updateContractMap["reimbursement"] = RAsset

		updateContract, err := contractAsset.Update(stub, updateContractMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to update contract asset")
		}

		paymentJSON, nerr := json.Marshal(updateContract)
		if nerr != nil {
			return nil, errors.WrapError(nil, "Failed to encode asset to JSON format")
		}

		return paymentJSON, nil

	},
}
