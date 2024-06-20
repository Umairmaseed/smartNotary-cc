package main

import (
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/smartescritura-cc/chaincode/assettypes"
)

var assetTypeList = []assets.AssetType{
	assettypes.Secret,
	assettypes.Contract,
	assettypes.Charge,
	assettypes.Holder,
	assettypes.Instalment,
	assettypes.NotaryData,
	assettypes.Payment,
	assettypes.Reimbursement,
	assettypes.Address,
}
