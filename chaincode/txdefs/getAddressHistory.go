package txdefs

import (
	"encoding/json"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	"github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

type AddressHistoryRecord struct {
	TxID      string          `json:"txId"`
	Timestamp string          `json:"timestamp"`
	Value     json.RawMessage `json:"value"`
	IsDeleted bool            `json:"isDeleted"`
}

var GetAddressHistory = tx.Transaction{
	Tag:         "getAddressHistory",
	Label:       "Get Address History",
	Description: "Get Address History",
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
			"@assetType": "address",
			"id":         id,
		})
		if err != nil {
			return nil, errors.WrapError(err, "failed to create key for address")
		}

		_, err = contractId.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to get address asset from the ledger")
		}

		historyIterator, err := stub.GetHistoryForKey(contractId.Key())
		if err != nil {
			return nil, errors.WrapError(err, "failed to read address history from blockchain")
		}
		defer historyIterator.Close()

		var history []ContractHistoryRecord
		for historyIterator.HasNext() {
			queryResponse, err := historyIterator.Next()
			if err != nil {
				return nil, errors.WrapError(err, "error iterating response")
			}

			var record ContractHistoryRecord
			record.TxID = queryResponse.TxId
			record.Timestamp = queryResponse.Timestamp.String()
			record.Value = queryResponse.Value
			record.IsDeleted = queryResponse.IsDelete

			history = append(history, record)
		}

		responseJSON, nerr := json.Marshal(history)
		if nerr != nil {
			return nil, errors.WrapError(err, "error marshaling history response")
		}
		return responseJSON, nil
	},
}
