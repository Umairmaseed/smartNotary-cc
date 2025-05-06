package main

import (
	tx "github.com/hyperledger-labs/cc-tools/transactions"
	"github.com/hyperledger-labs/smartNotary-cc/chaincode/txdefs"
)

var txList = []tx.Transaction{
	tx.CreateAsset,
	tx.UpdateAsset,
	tx.DeleteAsset,
	txdefs.CreateContract,
	txdefs.UpdateContractEvent,
	txdefs.PayInstallment,
	txdefs.Reimbursement,
	txdefs.GetContract,
	txdefs.GetContractHistory,
	txdefs.GetAddressHistory,
	txdefs.GetChargeHistory,
	txdefs.GetHolderHistory,
	txdefs.GetInstallmentHistory,
	txdefs.GetNotaryDataHistory,
	txdefs.GetPaymentHistory,
	txdefs.GetReimbursementHistory,
}
