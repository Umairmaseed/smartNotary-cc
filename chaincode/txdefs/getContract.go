package txdefs

import (
	"encoding/json"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	"github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var GetContract = tx.Transaction{
	Tag:         "getContract",
	Label:       "Get Contract",
	Description: "Get Contract",
	Method:      "GET",

	Args: []tx.Argument{
		{
			Required: true,
			Tag:      "id",
			Label:    "Id",
			DataType: "string",
		},
	},
	Routine: func(stub *stubwrapper.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		id, ok := req["id"].(string)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'id'")
		}

		contractId, err := assets.NewKey(map[string]interface{}{
			"@assetType": "contract",
			"id":         id,
		})
		if err != nil {
			return nil, errors.WrapError(err, "failed to create key for contract")
		}

		contractAsset, err := contractId.GetRecursive(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to get contract asset from the ledger")
		}

		contractJSON, nerr := json.Marshal(contractAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "Failed to encode asset to JSON format")
		}

		return contractJSON, nil

	},
}
