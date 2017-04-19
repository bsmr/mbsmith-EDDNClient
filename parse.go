package EDDNClient

import (
	"compress/zlib"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mbsmith/EDDNClient/blackmarket"
	"github.com/mbsmith/EDDNClient/commodity"
	"github.com/mbsmith/EDDNClient/journal"
	"github.com/mbsmith/EDDNClient/outfitting"
	"github.com/mbsmith/EDDNClient/shipyard"
	"io/ioutil"
	"strings"
)

// Root is the root of every JSON message received from EDDN.  This should
// not be used directly as this is lazily parsed to find the schema first.
type Root struct {
	SchemaRef string          `json:"$schemaRef"` // The schema of the message
	Header    Header          `json:"header"`     // The message header
	Message   json.RawMessage `json:"message"`    // The message unparsed until later
}

// Header type that is common to all messages.  This bit is only used by the parser
// however.  The types sent by the ChannelInterface will have their own
// Root/Header types that the receiver should use.
type Header struct {
	GatewayTimestamp string `json:"gatewayTimestamp,omitempty"` // Timestamp
	SoftwareName     string `json:"softwareName"`               // Software that sent the data
	SoftwareVersion  string `json:"softwareVersion"`            // Software version
	UploaderID       string `json:"uploaderID"`                 // ID of the uploader
}

func parseJSON(data string) (parsed interface{}, err error) {
	r, _ := zlib.NewReader(strings.NewReader(data))
	defer r.Close()

	output, err := ioutil.ReadAll(r)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	// Parse the schema to find out what kind of message we're going to be
	// handling.
	var jsonData Root

	err = json.Unmarshal(output, &jsonData)

	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	switch jsonData.SchemaRef {
	case "http://schemas.elite-markets.net/eddn/commodity/1":
		fallthrough
	case "http://schemas.elite-markets.net/eddn/commodity/2":
		err := errors.New("commodity versions 1 and 2 not currently supported")
		return nil, err

	case "http://schemas.elite-markets.net/eddn/commodity/3":
		var commodityData commodity.Root
		json.Unmarshal(output, &commodityData)
		return commodityData, nil

	case "http://schemas.elite-markets.net/eddn/journal/1":
		var journalData journal.Root
		json.Unmarshal(output, &journalData)
		return journalData, nil

	case "http://schemas.elite-markets.net/eddn/outfitting/1":
		err := errors.New("outfitting version 1 is not currently supported")
		return nil, err

	case "http://schemas.elite-markets.net/eddn/outfitting/2":
		var outfittingData outfitting.Root
		json.Unmarshal(output, &outfittingData)
		return outfittingData, nil

	case "http://schemas.elite-markets.net/eddn/blackmarket/1":
		var blackmarketData blackmarket.Root
		json.Unmarshal(output, &blackmarketData)
		return blackmarketData, nil

	case "http://schemas.elite-markets.net/eddn/shipyard/1":
		err := errors.New("shipyard version 1 is not currently supported")
		return nil, err

	case "http://schemas.elite-markets.net/eddn/shipyard/2":
		var shipyardData shipyard.Root
		json.Unmarshal(output, &shipyardData)
		return shipyardData, nil

	default:
		err := fmt.Errorf("unhandled schema: '%s'", jsonData.SchemaRef)
		return nil, err
	}

}
