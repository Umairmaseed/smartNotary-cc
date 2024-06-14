package assettypes

import "github.com/hyperledger-labs/cc-tools/assets"

var NotaryData = assets.AssetType{
	Tag:         "notaryData",
	Label:       "NotaryData",
	Description: "NotaryData",

	Props: []assets.AssetProp{
		{
			IsKey:    true,
			Tag:      "signerDocumentId",
			Label:    "Signer Document Id",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "mne",
			Label:    "MNE",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "notarizationDate",
			Label:    "Notarization Date",
			DataType: "string",
		},
	},
}
