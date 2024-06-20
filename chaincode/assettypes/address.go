package assettypes

import "github.com/hyperledger-labs/cc-tools/assets"

var Address = assets.AssetType{
	Tag:         "address",
	Label:       "Address",
	Description: "Address",

	Props: []assets.AssetProp{
		{
			IsKey:    true,
			Tag:      "id",
			Label:    "Id",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "ufCode",
			Label:    "ufCode",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "city",
			Label:    "City",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "neighborhood",
			Label:    "Neighborhood",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "zipCode",
			Label:    "ZipCode",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "number",
			Label:    "Number",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "additionalInfo",
			Label:    "AdditionalInfo",
			DataType: "string",
		},
	},
}
