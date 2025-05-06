package main

import (
	"github.com/hyperledger-labs/cc-tools/events"
	"github.com/hyperledger-labs/smartNotary-cc/chaincode/eventtypes"
)

var eventTypeList = []events.Event{
	eventtypes.CreateLibraryLog,
}
