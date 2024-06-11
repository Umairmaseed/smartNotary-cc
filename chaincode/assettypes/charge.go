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
			Tag:      "absoluteDefautlFineOverSignalPayment",
			Label:    "Absolute Defautl Fine Over Signal Payment",
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
			Tag:      "hasOwnership",
			Label:    "Has Ownership",
			DataType: "boolean",
		},
		{
			Required: true,
			Tag:      "ownershipDate",
			Label:    "Ownership Date",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "deadlineToLeavePropertyInTermination",
			Label:    "Deadline To Leave Property In Termination",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "fineInCaseOfStayingInProperty",
			Label:    "Fine InCase Of Staying In Property",
			DataType: "number",
		},
	},
}
