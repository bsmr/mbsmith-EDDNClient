package EDDNClient_test

import (
	"fmt"
	eddn "github.com/mbsmith/EDDNClient"
	"log"
)

func ExampleChannelInterface() {
	// Create a new channel interface that filters everything but journal
	// messages.
	channelInterface, err := eddn.NewChannelInterface(eddn.FilterShipyard |
		eddn.FilterCommodity | eddn.FilterOutfitting | eddn.FilterBlackmarket)

	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

loop:
	for {

		// Get a single journal message and then close/exit.
		select {
		case journalMsg := <-channelInterface.JournalChan:
			fmt.Printf("Schema: %v", journalMsg.SchemaRef)
			break loop
		}

	}

	channelInterface.Close()

	// Output:
	// Schema: http://schemas.elite-markets.net/eddn/journal/1
}
