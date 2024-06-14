package assettypes

import "github.com/hyperledger-labs/cc-tools/assets"

var Holder = assets.AssetType{
	Tag:         "holder",
	Label:       "Holder",
	Description: "Holder",

	Props: []assets.AssetProp{
		{
			IsKey:    true,
			Tag:      "holderId",
			Label:    "Holder Id",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "role",
			Label:    "Role",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "ownedPercentage",
			Label:    "Owned Percentage",
			DataType: "number",
		},
		{
			Required: true,
			Tag:      "percentage",
			Label:    "Percentage",
			DataType: "number",
		},
	},
}
