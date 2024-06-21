package assettypes

import "github.com/hyperledger-labs/cc-tools/assets"

var Contract = assets.AssetType{
	Tag:         "contract",
	Label:       "Contract",
	Description: "Contract",

	Props: []assets.AssetProp{
		{
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "Id",
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
			Tag:      "notarySubscriptionId",
			Label:    "Notary Subscription Id",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "assetRegistry",
			Label:    "Asset Registry",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "cns",
			Label:    "CNS",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "type",
			Label:    "Type",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "realStateAssetTypeCode",
			Label:    "Real StateAsset Type Code",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "propertyTypeCode",
			Label:    "PropertyTypeCode",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "value",
			Label:    "Value",
			DataType: "number",
		},
		{
			Required: true,
			Tag:      "creationDateUtc",
			Label:    "Creation Date Utc",
			DataType: "datetime",
		},
		{
			Required: true,
			Tag:      "ccir",
			Label:    "CCIR",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "nirf",
			Label:    "NIRF",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "charge",
			Label:    "Charge",
			DataType: "->charge",
		},
		{
			Required: true,
			Tag:      "holders",
			Label:    "Holders",
			DataType: "[]->holder",
		},
		{
			Required: true,
			Tag:      "installments",
			Label:    "Installments",
			DataType: "[]->installment",
		},
		{
			Required: true,
			Tag:      "notaryData",
			Label:    "Notary Data",
			DataType: "->notaryData",
		},
		{
			Tag:      "payment",
			Label:    "Payment",
			DataType: "[]->payment",
		},
		{
			Tag:      "reimbursement",
			Label:    "Reimbursement",
			DataType: "->reimbursement",
		},
		{
			Tag:      "statusCode",
			Label:    "StatusCode",
			DataType: "string",
		},
		{
			Tag:      "lastEventDate",
			Label:    "Last Event Date",
			DataType: "datetime",
		},
		{
			Required: true,
			Tag:      "address",
			Label:    "Address",
			DataType: "->address",
		},
	},
}
