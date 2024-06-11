package txdefs

import (
	"encoding/json"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var CreateContract = tx.Transaction{
	Tag:         "createContract",
	Label:       "Create Contract",
	Description: "Create Contract",
	Method:      "POST",

	Args: []tx.Argument{
		{
			Required: true,
			Tag:      "id",
			Label:    "Id",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "discriminator",
			Label:    "Discriminator",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "notarySubscriptionId",
			Label:    "Notary Subscription Id",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "assetRegistry",
			Label:    "Asset Registry",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "cns",
			Label:    "CNS",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "type",
			Label:    "Type",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "realStateAssetTypeCode",
			Label:    "Real StateAsset Type Code",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "propertyTypeCode",
			Label:    "PropertyTypeCode",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "value",
			Label:    "Value",
			DataType: "number",
		},
		{
			Required: true,
			Tag:      "creationDateUtc",
			Label:    "Creation Date Utc",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "ccir",
			Label:    "CCIR",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "nirf",
			Label:    "NIRF",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "charge",
			Label:    "Charge",
			DataType: "->charge",
		},
		{
			Required: true,
			Tag:      "holders",
			Label:    "Holders",
			DataType: "[]->holder",
		},
		{
			Required: true,
			Tag:      "installments",
			Label:    "Installments",
			DataType: "[]->installment",
		},
		{
			Required: true,
			Tag:      "notaryData",
			Label:    "Notary Data",
			DataType: "->notaryData",
		},
		{
			Tag:      "payment",
			Label:    "Payment",
			DataType: "[]->payment",
		},
		{
			Tag:      "reimbursement",
			Label:    "Reimbursement",
			DataType: "->reimbursement",
		},
		{
			Tag:      "statusCode",
			Label:    "StatusCode",
			DataType: "string",
		},
		{
			Tag:      "lastEventDate",
			Label:    "Last Event Date",
			DataType: "string",
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		id, ok := req["id"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'id'")
		}

		discriminator, ok := req["discriminator"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'discriminator'")
		}

		notarySubscriptionId, ok := req["notarySubscriptionId"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'notarySubscriptionId'")
		}

		assetRegistry, ok := req["assetRegistry"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'assetRegistry'")
		}

		cns, ok := req["cns"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'cns'")
		}

		contractType, ok := req["type"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'type'")
		}

		realStateAssetTypeCode, ok := req["realStateAssetTypeCode"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'realStateAssetTypeCode'")
		}

		propertyTypeCode, ok := req["propertyTypeCode"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'propertyTypeCode'")
		}

		value, ok := req["value"].(float64)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'value'")
		}

		creationDateUtc, ok := req["creationDateUtc"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'creationDateUtc'")
		}

		ccir, ok := req["ccir"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'ccir'")
		}

		nirf, ok := req["nirf"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'nirf'")
		}

		charge, ok := req["charge"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'charge'")
		}

		holders, ok := req["holders"].([]interface{})
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'holders'")
		}

		installments, ok := req["installments"].([]interface{})
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installments'")
		}

		notaryData, ok := req["notaryData"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'notaryData'")
		}

		payment, ok := req["payment"].([]interface{})
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'payment'")
		}

		reimbursement, ok := req["reimbursement"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'reimbursement'")
		}

		statusCode, ok := req["statusCode"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'statusCode'")
		}

		lastEventDate, ok := req["lastEventDate"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'lastEventDate'")
		}

		contractMap := map[string]interface{}{
			"@assetType":             "contract",
			"id":                     id,
			"discriminator":          discriminator,
			"notarySubscriptionId":   notarySubscriptionId,
			"assetRegistry":          assetRegistry,
			"cns":                    cns,
			"type":                   contractType,
			"realStateAssetTypeCode": realStateAssetTypeCode,
			"propertyTypeCode":       propertyTypeCode,
			"value":                  value,
			"creationDateUtc":        creationDateUtc,
			"ccir":                   ccir,
			"nirf":                   nirf,
			"charge":                 charge,
			"holders":                holders,
			"installments":           installments,
			"notaryData":             notaryData,
			"payment":                payment,
			"reimbursement":          reimbursement,
			"statusCode":             statusCode,
			"lastEventDate":          lastEventDate,
		}

		contractAsset, err := assets.NewAsset(contractMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		_, err = contractAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		contractJSON, nerr := json.Marshal(contractAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "Failed to encode asset to JSON format")
		}

		return contractJSON, nil
	},
}
