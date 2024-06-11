package main

import (
	"github.com/hyperledger-labs/smartescritura-cc/chaincode/eventtypes"
	"github.com/hyperledger-labs/cc-tools/events"
)

var eventTypeList = []events.Event{
	eventtypes.CreateLibraryLog,
}
