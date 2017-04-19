package EDDNClient

import (
	"fmt"
	"github.com/mbsmith/EDDNClient/blackmarket"
	"github.com/mbsmith/EDDNClient/commodity"
	"github.com/mbsmith/EDDNClient/journal"
	"github.com/mbsmith/EDDNClient/outfitting"
	"github.com/mbsmith/EDDNClient/shipyard"
	zmq "github.com/pebbe/zmq4"
)

// Constants used internally by control methods to control the ChannelInterface
// goroutine.
const (
	channelInterfaceClose = iota
)

// An enumeration of filters used to tell the ChannelInterface what data the
// receiver is interested in.  These can be OR'd together to build any filter
// the receiver wishes.
const (
	FilterNone        = 1 << iota // Filter nothing
	FilterJournal     = 1 << iota // Filter journal messages
	FilterShipyard    = 1 << iota // Filter shipyard messages
	FilterCommodity   = 1 << iota // Filter commodity messages
	FilterBlackmarket = 1 << iota // Filter blackmarket messages
	FilterOutfitting  = 1 << iota // Filter outfitting messages.
)

// A ChannelInterface provides an interface to a group of channels that
// each provide various types of EDDN data separated into their respective
// channels.  JournalChan, ShipyardChan, CommodityChan, BlackmarketChan,
// and OutfittingChan each only send messages pertaining to their
// respective types.  When a Done message is received all processing on the
// receiver should halt as this means that the ChannelInterface is closed.
type ChannelInterface struct {
	Socket          *zmq.Socket             // Underlying ZeroMQ socket
	JournalChan     <-chan journal.Root     // Channel for reading journal messages
	ShipyardChan    <-chan shipyard.Root    // Channel for reading shipyard messages
	CommodityChan   <-chan commodity.Root   // Channel for reading commodity messages
	BlackmarketChan <-chan blackmarket.Root // Channel for reading blackmarket messages
	OutfittingChan  <-chan outfitting.Root  // Channel for reading outfitting messages
	ControlChan     chan<- int              // Channel providing control of the goroutine
	Done            chan bool               // Sent when the ChannelInterface is finished.
}

// NewChannelInterface creates an active ChannelInterface using the provided
// filter.  If one wishes no filters then 0, or FilterNone can be passed here.
// If an error is found then err will not be nil and shall be returned.  The
// ChannelInterface will be active immediately and must be stopped with Close()
// if the receiver wishes to stop receiving messages.
//
// Should the receiver wish to begin receiving messages again then a new
// ChannelInterface must be created.
func NewChannelInterface(filter int) (channels *ChannelInterface, err error) {

	subscriber, err := zmq.NewSocket(zmq.SUB)

	if err != nil {
		return nil, err
	}

	subscriber.Connect(EddnAddress)
	subscriber.SetSubscribe("")

	journalChan := make(chan journal.Root)
	shipyardChan := make(chan shipyard.Root)
	commodityChan := make(chan commodity.Root)
	blackmarketChan := make(chan blackmarket.Root)
	outfittingChan := make(chan outfitting.Root)
	controlChan := make(chan int, 1)
	Done := make(chan bool)

	go func() {
		defer close(journalChan)
		defer close(shipyardChan)
		defer close(commodityChan)
		defer close(blackmarketChan)
		defer close(outfittingChan)
		defer close(controlChan)
		defer close(Done)

		for {
			// Check if we have any control messages first.
			select {
			case control := <-controlChan:
				switch control {
				case channelInterfaceClose:
					Done <- true
					return
				}
			default:
				// NOOP
			}

			eddnData, err := subscriber.Recv(0)

			if err != nil {
				fmt.Printf("Error: %v", err)
				continue
			}

			Message, err := parseJSON(eddnData)

			if err != nil {
				fmt.Printf("Error: %v", err)
				continue
			}

			switch Message.(type) {
			case journal.Root:

				if filter&FilterJournal == 0 {
					journalChan <- Message.(journal.Root)
				}

			case shipyard.Root:

				if filter&FilterShipyard == 0 {
					shipyardChan <- Message.(shipyard.Root)
				}

			case commodity.Root:

				if filter&FilterCommodity == 0 {
					commodityChan <- Message.(commodity.Root)
				}

			case blackmarket.Root:

				if filter&FilterBlackmarket == 0 {
					blackmarketChan <- Message.(blackmarket.Root)
				}

			case outfitting.Root:

				if filter&FilterOutfitting == 0 {
					outfittingChan <- Message.(outfitting.Root)
				}

			default:
				continue
			}
		}
	}()

	return &ChannelInterface{subscriber, journalChan, shipyardChan,
		commodityChan, blackmarketChan,
		outfittingChan, controlChan, Done}, nil
}

// Close closes the given ChannelInterface ci.
func (ci *ChannelInterface) Close() {
	ci.ControlChan <- channelInterfaceClose
}
