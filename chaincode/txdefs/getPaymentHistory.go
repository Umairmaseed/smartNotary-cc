package txdefs

import (
	"encoding/json"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	"github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

type PaymentHistoryRecord struct {
	TxID      string          `json:"txId"`
	Timestamp string          `json:"timestamp"`
	Value     json.RawMessage `json:"value"`
	IsDeleted bool            `json:"isDeleted"`
}

var GetPaymentHistory = tx.Transaction{
	Tag:         "getPaymentHistory",
	Label:       "Get Payment History",
	Description: "Get Payment History",
	Method:      "GET",

	Args: []tx.Argument{
		{
			Required: true,
			Tag:      "installment",
			Label:    "Installment",
			DataType: "->installment",
		},
	},
	Routine: func(stub *stubwrapper.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		installmentAsset, ok := req["installment"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Invalid type for parameter 'installment'")
		}

		contractId, err := assets.NewKey(map[string]interface{}{
			"@assetType":  "payment",
			"installment": installmentAsset,
		})
		if err != nil {
			return nil, errors.WrapError(err, "failed to create key for payment")
		}

		_, err = contractId.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to get payment asset from the ledger")
		}

		historyIterator, err := stub.GetHistoryForKey(contractId.Key())
		if err != nil {
			return nil, errors.WrapError(err, "failed to read payment history from blockchain")
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
