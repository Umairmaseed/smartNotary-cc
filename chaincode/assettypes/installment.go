package assettypes

import "github.com/hyperledger-labs/cc-tools/assets"

var Instalment = assets.AssetType{
	Tag:         "installment",
	Label:       "Installment",
	Description: "Installment",

	Props: []assets.AssetProp{
		{
			IsKey:    true,
			Tag:      "id",
			Label:    "Id",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "number",
			Label:    "number",
			DataType: "integer",
		},
		{
			Required: true,
			Tag:      "dueDateUtc",
			Label:    "Due Date Utc",
			DataType: "datetime",
		},
		{
			Required: true,
			Tag:      "amortization",
			Label:    "Amortization",
			DataType: "number",
		},
		{
			Required: true,
			Tag:      "financingModalityInterestValue",
			Label:    "Financing Modality Interest Value",
			DataType: "number",
		},
		{
			Required: true,
			Tag:      "installmentValue",
			Label:    "Installment Value",
			DataType: "number",
		},
		{
			Tag:      "fineInCaseOfDelayValue",
			Label:    "Fine InCase Of Delay Value",
			DataType: "number",
		},
		{
			Tag:      "interestInCaseOfDelayValue",
			Label:    "Interest InCase Of Delay Value",
			DataType: "number",
		},
		{
			Tag:      "monetaryCorrectionValue",
			Label:    "Monetary Correction Value",
			DataType: "number",
		},
		{
			Tag:      "reimbursement",
			Label:    "reimbursement",
			DataType: "number",
		},
	},
}
