package assettypes

import "github.com/hyperledger-labs/cc-tools/assets"

var Reimbursement = assets.AssetType{
	Tag:         "reimbursement",
	Label:       "Reimbursement",
	Description: "Reimbursement",

	Props: []assets.AssetProp{
		{
			Tag:      "cancelledInstallments",
			Label:    "Cancelled Installments",
			DataType: "[]->installment",
		},
		{
			IsKey:    true,
			Tag:      "reimbursementInstallment",
			Label:    "Reimbursement Installment",
			DataType: "->installment",
		},
	},
}
