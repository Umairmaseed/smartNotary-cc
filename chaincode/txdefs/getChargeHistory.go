package txdefs

import (
	"encoding/json"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	"github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

type ChargeHistoryRecord struct {
	TxID      string          `json:"txId"`
	Timestamp string          `json:"timestamp"`
	Value     json.RawMessage `json:"value"`
	IsDeleted bool            `json:"isDeleted"`
}

var GetChargeHistory = tx.Transaction{
	Tag:         "getChargeHistory",
	Label:       "Get Charge History",
	Description: "Get Charge History",
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
			"@assetType": "charge",
			"id":         id,
		})
		if err != nil {
			return nil, errors.WrapError(err, "failed to create key for charge")
		}

		_, err = contractId.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to get charge asset from the ledger")
		}

		historyIterator, err := stub.GetHistoryForKey(contractId.Key())
		if err != nil {
			return nil, errors.WrapError(err, "failed to read charge history from blockchain")
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
