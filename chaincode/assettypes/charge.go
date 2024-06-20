package assettypes

import "github.com/hyperledger-labs/cc-tools/assets"

var Charge = assets.AssetType{
	Tag:         "charge",
	Label:       "Charge",
	Description: "Charge",

	Props: []assets.AssetProp{
		{
			IsKey:    true,
			Tag:      "id",
			Label:    "Id",
			DataType: "string",
		},
		// id should be generate automatically and should be unique
		{
			Required: true,
			Tag:      "financingModalityCode",
			Label:    "Financing Modality Code",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "financingModalityInterest",
			Label:    "Financing Modality Interest",
			DataType: "number",
		},
		{
			Required: true,
			Tag:      "monetaryCorrectionIndexCode",
			Label:    "Monetary Correction Index Code",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "monetaryCorrectionIndexPeriodicityCode",
			Label:    "Monetary Correction Index Periodicity Code",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "fineInCaseOfDelay",
			Label:    "Fine InCase Of Delay",
			DataType: "number",
		},
		{
			Required: true,
			Tag:      "interestInCaseOfDelay",
			Label:    "Interest InCase Of Delay",
			DataType: "number",
		},
		{
			Required: true,
			Tag:      "delayGracePeriodDays",
			Label:    "Delay Grace Period Days",
			DataType: "integer",
		},
		{
			Required: true,
			Tag:      "absoluteDefaultFineOverSignalPayment",
			Label:    "Absolute Default Fine Over Signal Payment",
			DataType: "number",
		},
		{
			Required: true,
			Tag:      "absoluteDefaultFineOverRemaining",
			Label:    "Absolute Default Fine Over Remaining",
			DataType: "number",
		},
		{
			Required: true,
			Tag:      "absoluteDefaultGracePeriodDays",
			Label:    "Absolute Default Grace Period Days",
			DataType: "integer",
		},
		{
			Required: true,
			Tag:      "hasPossession",
			Label:    "Has Possession",
			DataType: "boolean",
		},
		{
			Required: true,
			Tag:      "possessionDate",
			Label:    "Possession Date",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "exitDate",
			Label:    "Exit Date",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "fineInCaseOfOverstay",
			Label:    "Fine InCase Of Overstay",
			DataType: "number",
		},
	},
}
