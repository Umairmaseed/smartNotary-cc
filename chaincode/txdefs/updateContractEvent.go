package txdefs

import (
	"encoding/json"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var UpdateContractEvent = tx.Transaction{
	Tag:         "updateContractEvent",
	Label:       "Update Contract Event",
	Description: "Register a event change on the contract",
	Method:      "POST",

	Args: []tx.Argument{
		{
			Tag:      "assetId",
			Label:    "Contract ID",
			DataType: "string",
			Required: true,
		},
		{
			Tag:      "statusCode",
			Label:    "Status Code",
			DataType: "string",
			Required: true,
		},
		{
			Tag:      "eventDate",
			Label:    "Event Date",
			DataType: "string",
			Required: true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		assetId, ok := req["assetId"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter assetId must be a string")
		}
		statusCode, ok := req["statusCode"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter statusCode must be a string")
		}
		eventDate, ok := req["eventDate"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter eventDate must be a string")
		}

		key, err := assets.NewKey(map[string]interface{}{
			"@assetType": "contract",
			"id":         assetId,
		})
		if err != nil {
			return nil, errors.WrapError(err, "failed to create key")
		}

		contractAsset, err := key.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to get contract asset from the ledger")
		}
		contractMap := (map[string]interface{})(*contractAsset)

		contractMap["statusCode"] = statusCode
		contractMap["lastEventDate"] = eventDate

		contractMap, err = contractAsset.Update(stub, contractMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to update contract asset")
		}

		contractJSON, nerr := json.Marshal(contractMap)
		if nerr != nil {
			return nil, errors.WrapError(err, "Failed to marshal updated contract to JSON format")
		}

		return contractJSON, nil
	},
}
