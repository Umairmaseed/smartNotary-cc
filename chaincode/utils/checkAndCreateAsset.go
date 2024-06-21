package utils

import (
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
)

func CheckAndCreateAsset(stub *sw.StubWrapper, assetMap map[string]interface{}, assetType string) (map[string]interface{}, errors.ICCError) {
	assetMap["@assetType"] = assetType
	assetKey, err := assets.NewKey(assetMap)
	if err != nil {
		return nil, errors.WrapError(err, "failed to create key for asset")
	}
	_, err = assetKey.Get(stub)
	if err == nil {
		return assetKey, nil
	} else if err.Status() == 404 {
		newAsset, err := assets.NewAsset(assetMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		_, err = newAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}
		return map[string]interface{}{
			"@assetType": assetType,
			"@key":       newAsset["@key"],
		}, nil
	} else {
		return nil, errors.WrapError(err, "Failed to check asset existence")
	}
}
