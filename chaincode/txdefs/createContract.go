package txdefs

import (
	"encoding/json"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
	"github.com/hyperledger-labs/smartNotary-cc/chaincode/utils"
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
			DataType: "@object",
		},
		{
			Required: true,
			Tag:      "holders",
			Label:    "Holders",
			DataType: "[]@object",
		},
		{
			Required: true,
			Tag:      "installments",
			Label:    "Installments",
			DataType: "[]@object",
		},
		{
			Required: true,
			Tag:      "notaryData",
			Label:    "Notary Data",
			DataType: "@object",
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
		{
			Required: true,
			Tag:      "address",
			Label:    "Address",
			DataType: "@object",
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

		statusCode, ok := req["statusCode"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'statusCode'")
		}

		lastEventDate, ok := req["lastEventDate"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'lastEventDate'")
		}

		//------------------------checking and creating charge asset-----------------------
		chargeMap, ok := req["charge"].(map[string]interface{})
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'charge'")
		}
		charge, err := utils.CheckAndCreateAsset(stub, chargeMap, "charge")
		if err != nil {
			return nil, errors.WrapError(err, "Failed to process 'charge'")
		}

		//------------------------checking and creating holder asset-----------------------
		holderMaps, ok := req["holders"].([]interface{})
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'holders'")
		}
		var holders []map[string]interface{}
		for _, holderMap := range holderMaps {
			holder, ok := holderMap.(map[string]interface{})
			if !ok {
				return nil, errors.WrapError(nil, "Invalid type in 'holders' array")
			}
			asset, err := utils.CheckAndCreateAsset(stub, holder, "holder")
			if err != nil {
				return nil, errors.WrapError(err, "Failed to process 'holders'")
			}
			holders = append(holders, asset)
		}

		//------------------------checking and creating installment asset-----------------------
		installmentMaps, ok := req["installments"].([]interface{})
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installments'")
		}
		var installments []map[string]interface{}
		for _, installmentMap := range installmentMaps {
			installment, ok := installmentMap.(map[string]interface{})
			if !ok {
				return nil, errors.WrapError(nil, "Invalid type in 'installments' array")
			}
			asset, err := utils.CheckAndCreateAsset(stub, installment, "installment")
			if err != nil {
				return nil, errors.WrapError(err, "Failed to process 'installments'")
			}
			installments = append(installments, asset)
		}

		//------------------------checking and creating notary data asset-----------------------
		notaryDataMap, ok := req["notaryData"].(map[string]interface{})
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'notaryData'")
		}
		notaryData, err := utils.CheckAndCreateAsset(stub, notaryDataMap, "notaryData")
		if err != nil {
			return nil, errors.WrapError(err, "Failed to process 'notaryData'")
		}
		//------------------------checking and creating address asset-----------------------
		addressMap, ok := req["address"].(map[string]interface{})
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'address'")
		}
		address, err := utils.CheckAndCreateAsset(stub, addressMap, "address")
		if err != nil {
			return nil, errors.WrapError(err, "Failed to process 'address'")
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
			"address":                address,
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
