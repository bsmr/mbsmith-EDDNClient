package EDDNClient

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"log"
	"time"
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
//
// It should be noted that the Journal channel can send several types that
// must be asserted by the caller.  While this may be a bit tedious it
// does provide type correctness, and allows the caller to know precisely
// what data was provided by EDDN.
type ChannelInterface struct {
	Socket          *zmq.Socket        // Underlying ZeroMQ socket
	JournalChan     <-chan Journal     // Channel for journal messages. (Provides many message										  // types.
	ShipyardChan    <-chan Shipyard    // Channel for reading shipyard messages
	CommodityChan   <-chan Commodity   // Channel for reading commodity messages
	BlackmarketChan <-chan Blackmarket // Channel for reading blackmarket messages
	OutfittingChan  <-chan Outfitting  // Channel for reading outfitting messages
	ControlChan     chan<- int         // Channel providing goroutine control
	Done            chan bool          // Sent when the ChannelInterface is done.
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

	subscriber.Connect(EDDNSubAddress)
	subscriber.SetSubscribe("")
	subscriber.SetConnectTimeout(time.Duration(600000))
	subscriber.SetHeartbeatIvl(500 * time.Millisecond)

	journalChan := make(chan Journal)
	shipyardChan := make(chan Shipyard)
	commodityChan := make(chan Commodity)
	blackmarketChan := make(chan Blackmarket)
	outfittingChan := make(chan Outfitting)
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
				log.Fatalln(err)
				continue
			}

			Message, err := parseJSON(eddnData)

			if err != nil {
				fmt.Printf("Error: %v", err)
				continue
			}

			switch Message.(type) {
			case Journal:

				if filter&FilterJournal == 0 {
					journalChan <- Message.(Journal)
				}

			case Shipyard:

				if filter&FilterShipyard == 0 {
					shipyardChan <- Message.(Shipyard)
				}

			case Commodity:

				if filter&FilterCommodity == 0 {
					commodityChan <- Message.(Commodity)
				}

			case Blackmarket:

				if filter&FilterBlackmarket == 0 {
					blackmarketChan <- Message.(Blackmarket)
				}

			case Outfitting:

				if filter&FilterOutfitting == 0 {
					outfittingChan <- Message.(Outfitting)
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
