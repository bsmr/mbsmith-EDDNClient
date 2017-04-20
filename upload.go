package EDDNClient

import (
	"encoding/json"
	//"github.com/xeipuuv/gojsonschema"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	blackmarketSchema = iota
	commoditySchema   = iota
	journalSchema     = iota
	outfittingSchema  = iota
	shipyardSchema    = iota
)

// Uploader is a helper type (required) that keeps track of the header, and
// other potential portions of data that don't need to be regenerated after
// each message.  It also updates its timestamp internally on each message
// so it shouldn't ever be off time.
type Uploader struct {
	header Header // header represents the header portion of the JSON in each message.
}

// NewUploader creates a new Uploader that will be used to send various types
// of messages.  uploaderID, softwareName, and softwareVersion should be set
// to values that you want represented in the header of every message you send.
func NewUploader(uploaderID string, softwareName string,
	softwareVersion string) (uploader *Uploader, err error) {
	header, err := generateHeader(uploaderID, softwareName, softwareVersion)

	if err != nil {
		return nil, err
	}

	return &Uploader{header}, nil
}

func generateSchema(schemaType int) (schema string, err error) {

	switch schemaType {
	case blackmarketSchema:
		return "http://schemas.elite-markets.net/eddn/blackmarket/1/test", nil
	case commoditySchema:
		return "http://schemas.elite-markets.net/eddn/commodity/3/test", nil
	case journalSchema:
		return "http://schemas.elite-markets.net/eddn/journal/1/test", nil
	case outfittingSchema:
		return "http://schemas.elite-markets.net/eddn/outfitting/2/test", nil
	case shipyardSchema:
		return "http://schemas.elite-markets.net/eddn/shipyard/2/test", nil
	default:
		return "", errors.New("Invalid schema provided")
	}

}

func generateHeader(uploaderID string, softwareName string,
	softwareVersion string) (header Header, err error) {

	var newHeader Header
	newHeader.UploaderID = uploaderID
	newHeader.SoftwareName = softwareName
	newHeader.SoftwareVersion = softwareVersion

	// Unsure if this is valid.
	UTCTime := time.Now().UTC()
	newHeader.GatewayTimestamp = UTCTime.String()

	return newHeader, nil
}

// Updates the header to the current time.  Nothing else really needs to
// change.
func (uploader *Uploader) updateHeader() {
	UTCTime := time.Now().UTC()
	uploader.header.GatewayTimestamp = UTCTime.String()
}

func checkResponse(body io.ReadCloser) (err error) {
	output, _ := ioutil.ReadAll(body)

	if string(output) != "OK" {
		errStr := fmt.Sprintf("Error sending blackmarket data: %s\n", output)
		return errors.New(errStr)
	}

	return nil
}

func sendMessage(msg interface{}) (err error) {
	jsonData, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	fmt.Printf("JSON: %s\n", string(jsonData))

	buf := bytes.NewBuffer(jsonData)

	resp, err := http.Post(EDDNUploadAddress, "application/json; charset=utf-8",
		buf)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return checkResponse(resp.Body)
}

// SendBlackmarket sends a blackmarket message to the EDDN servers.  The
// message should be filled (especially the required fields).  The required
// fields are marked in the blackmarket.go source file.
func (uploader *Uploader) SendBlackmarket(msg *BlackmarketMessage) (err error) {
	schema, err := generateSchema(blackmarketSchema)

	if err != nil {
		return err
	}

	uploader.updateHeader()

	data := &Blackmarket{schema, uploader.header, *msg}

	return sendMessage(data)
}

// SendCommodity sends a commodity message to the EDDN servers.  The
// message should be filled (especially the required fields).  The required
// fields are marked in the commodity.go source file.
func (uploader *Uploader) SendCommodity(msg *CommodityMessage) (err error) {
	schema, err := generateSchema(commoditySchema)

	if err != nil {
		return err
	}

	uploader.updateHeader()

	data := &Commodity{schema, uploader.header, *msg}

	return sendMessage(data)
}

// SendJournal sends a journal message to the EDDN servers.  The
// message should be filled (especially the required fields).  The required
// fields are marked in the journal.go source file.
func (uploader *Uploader) SendJournal(msg *JournalMessage) (err error) {
	schema, err := generateSchema(journalSchema)

	if err != nil {
		return err
	}

	uploader.updateHeader()

	data := &Journal{schema, uploader.header, *msg}

	return sendMessage(data)
}

// SendOutfitting sends a outfitting message to the EDDN servers.  The
// message should be filled (especially the required fields).  The required
// fields are marked in the outfitting.go source file.
func (uploader *Uploader) SendOutfitting(msg *OutfittingMessage) (err error) {
	schema, err := generateSchema(commoditySchema)

	if err != nil {
		return err
	}

	uploader.updateHeader()

	data := &Outfitting{schema, uploader.header, *msg}

	return sendMessage(data)
}

// SendShipyard sends a shipyard message to the EDDN servers.  The
// message should be filled (especially the required fields).  The required
// fields are marked in the shipyard.go source file.
func (uploader *Uploader) SendShipyard(msg *ShipyardMessage) (err error) {
	schema, err := generateSchema(commoditySchema)

	if err != nil {
		return err
	}

	uploader.updateHeader()

	data := &Shipyard{schema, uploader.header, *msg}

	return sendMessage(data)
}
