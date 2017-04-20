// Simple client that echos any data it receives via EDDN.  This is NOT how
// I see this library being used.  This does show some of the flexibility in
// using the interface even for something it's not intended for.  As well as
// how easy it is to convert to and from go types to json and back.
//
// The channels send data from a goroutine (concurrent routine), convert the
// given JSON into native Go types, and then send them to the receiver, (in
// this case, our echo_client).  We then convert the types BACK to json, and
// echo them to stdout pretty printed, or otherwise.  This should still
// display some basic usage however, and could even be of some use to some
// odd duck out there. :)

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	eddn "github.com/mbsmith/EDDNClient"
	"log"
	"os"
	"strings"
)

// Handle some flags
type filter []string

// Diagnostic method for flags.
func (f *filter) String() string {
	return fmt.Sprint(*f)
}

// Handles the filter flag which contains comma-separated values.
func (f *filter) Set(value string) error {
	if len(*f) > 0 {
		return errors.New("filter flag already set")
	}

	for _, ft := range strings.Split(value, ",") {
		filter := ft
		*f = append(*f, strings.ToLower(filter))
	}

	return nil
}

var filterFlag filter
var prettyPrint = flag.Bool("pretty-print", false, "Pretty print JSON we receive.")

func init() {
	// Tie the filter flag to filter
	flag.Var(&filterFlag, "filters",
		"comma-separated values of results to filter. [outfitting, journal, shipyard, commodity, and blackmarket]")
}

func output(data []byte) {
	if *prettyPrint {
		var out bytes.Buffer
		json.Indent(&out, data, "=", "\t")
		out.WriteTo(os.Stdout)
		return
	}

	buf := bytes.NewBuffer(data)
	buf.WriteTo(os.Stdout)
}

func handleConnections(channelInterface *eddn.ChannelInterface) {
	// Loop forever, outputting whatever we get.  Obviously the error checking
	// is very lax here so I wouldn't recommend actually doing something this
	// way, but it does show the flexibility of the types, and the basic usage
	// of the channels.
	for {
		select {
		case journalMsg := <-channelInterface.JournalChan:
			b, _ := json.Marshal(journalMsg)
			output(b)

		case bmMessage := <-channelInterface.BlackmarketChan:
			b, _ := json.Marshal(bmMessage)
			output(b)

		case commMessage := <-channelInterface.CommodityChan:
			b, _ := json.Marshal(commMessage)
			output(b)

		case outfMessage := <-channelInterface.OutfittingChan:
			b, _ := json.Marshal(outfMessage)
			output(b)

		case shipMessage := <-channelInterface.ShipyardChan:
			b, _ := json.Marshal(shipMessage)
			output(b)

		default:
			// NOOP
		}
	}
}

func main() {
	flag.Parse()

	var filters int

	// Handle filters if any.
	for _, filter := range filterFlag {

		switch filter {
		case "outfitting":
			filters |= eddn.FilterOutfitting
		case "journal":
			filters |= eddn.FilterJournal
		case "shipyard":
			filters |= eddn.FilterShipyard
		case "commodity":
			filters |= eddn.FilterCommodity
		case "blackmarket":
			filters |= eddn.FilterBlackmarket
		default:
			log.Printf("%s is not a valid filter", filter)
			continue
		}

	}

	// Now on to the good stuff.
	channelInterface, err := eddn.NewChannelInterface(filters)

	if err != nil {
		log.Fatalln(err)
		return
	}

	handleConnections(channelInterface)
}
