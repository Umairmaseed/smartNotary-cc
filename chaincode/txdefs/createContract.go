package txdefs

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	"github.com/hyperledger-labs/cc-tools/events"
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
		id, _ := req["id"].(string)
		discriminator, _ := req["discriminator"].(string)
		notarySubscriptionId, _ := req["notarySubscriptionId"].(string)
		assetRegistry, _ := req["assetRegistry"].(string)
		cns, _ := req["cns"].(string)
		contractType, _ := req["type"].(string)
		realStateAssetTypeCode, _ := req["realStateAssetTypeCode"].(string)
		propertyTypeCode, _ := req["propertyTypeCode"].(string)
		value, _ := req["value"].(float64)
		creationDateUtc, _ := req["creationDateUtc"].(string)
		ccir, _ := req["ccir"].(string)
		nirf, _ := req["nirf"].(string)
		charge, _ := req["charge"].(assets.Asset)
		holders, _ := req["holders"].([]assets.Asset)
		installments, _ := req["installments"].([]assets.Asset)
		notaryData, _ := req["notaryData"].(assets.Asset)
		payment, _ := req["payment"].([]assets.Asset)
		reimbursement, _ := req["reimbursement"].(assets.Asset)
		statusCode, _ := req["statusCode"].(string)
		lastEventDate, _ := req["lastEventDate"].(string)

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

		logMsg, ok := json.Marshal(fmt.Sprintf("New contract created with ID: %s", id))
		if ok != nil {
			return nil, errors.WrapError(nil, "Failed to encode log message to JSON format")
		}

		events.CallEvent(stub, "createContractLog", logMsg)

		return contractJSON, nil
	},
}
