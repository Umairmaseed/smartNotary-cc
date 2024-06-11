package assettypes

import "github.com/hyperledger-labs/cc-tools/assets"

var Payment = assets.AssetType{
	Tag:         "payment",
	Label:       "Payment",
	Description: "Payment",

	Props: []assets.AssetProp{
		{
			IsKey:    true,
			Tag:      "installment",
			Label:    "Installment",
			DataType: "->installment",
		},
		{
			Required: true,
			Tag:      "value",
			Label:    "Value",
			DataType: "number",
		},
		{
			Required: true,
			Tag:      "dateCompletedUtc",
			Label:    "Date Completed Utc",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "discriminator",
			Label:    "Discriminator",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "gatewaySaleId",
			Label:    "Gateway Sale Id",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "gatewaySellerId",
			Label:    "Gateway Seller Id",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "gatewayFee",
			Label:    "Gateway Fee",
			DataType: "number",
		},
		{
			Required: true,
			Tag:      "reimbursement",
			Label:    "Reimbursement",
			DataType: "number",
		},
	},
}
