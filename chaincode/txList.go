package main

import (
	tx "github.com/hyperledger-labs/cc-tools/transactions"
	"github.com/hyperledger-labs/smartescritura-cc/chaincode/txdefs"
)

var txList = []tx.Transaction{
	tx.CreateAsset,
	tx.UpdateAsset,
	tx.DeleteAsset,
	txdefs.CreateContract,
	txdefs.UpdateContractEvent,
}
