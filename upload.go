package EDDNClient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"io"
	"io/ioutil"
	"log"
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

// Current schema URI's
const bmSchemaURI = "https://raw.githubusercontent.com/jamesremuscat/EDDN/master/schemas/blackmarket-v1.0.json"
const comSchemaURI = "https://raw.githubusercontent.com/jamesremuscat/EDDN/master/schemas/commodity-v3.0.json"
const jSchemaURI = "https://raw.githubusercontent.com/jamesremuscat/EDDN/master/schemas/journal-v1.0.json"
const outfSchemaURI = "https://raw.githubusercontent.com/jamesremuscat/EDDN/master/schemas/outfitting-v2.0.json"
const shipSchemaURI = "https://raw.githubusercontent.com/jamesremuscat/EDDN/master/schemas/shipyard-v2.0.json"

// Uploader is a helper type (required) that keeps track of the header, and
// other potential portions of data that don't need to be regenerated after
// each message.  It also updates its timestamp internally on each message
// so it shouldn't ever be off time.
type Uploader struct {
	header            Header               // header sent with each message.
	blackmarketSchema *gojsonschema.Schema // JSON validation for blackmarket messages
	commoditySchema   *gojsonschema.Schema // JSON validation for commodity messages
	journalSchema     *gojsonschema.Schema // JSON validation for journal messages
	outfittingSchema  *gojsonschema.Schema // JSON validation for outfitting messages
	shipyardSchema    *gojsonschema.Schema // JSON validation for shipyard messages
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

	var e error

	schemaLoad := func(uri string) *gojsonschema.Schema {
		if e != nil {
			return nil
		}

		loader := gojsonschema.NewReferenceLoader(uri)
		var schema *gojsonschema.Schema
		schema, e = gojsonschema.NewSchema(loader)

		return schema
	}

	// Prepare various schemas for validation.
	bmSchema := schemaLoad(bmSchemaURI)
	comSchema := schemaLoad(comSchemaURI)
	jSchema := schemaLoad(jSchemaURI)
	outSchema := schemaLoad(outfSchemaURI)
	shipSchema := schemaLoad(shipSchemaURI)

	if e != nil {
		return nil, e
	}

	return &Uploader{header, bmSchema, comSchema, jSchema, outSchema,
		shipSchema}, nil
}

func generateSchema(schemaType int) (schema string, err error) {

	switch schemaType {
	case blackmarketSchema:
		return "http://schemas.elite-markets.net/eddn/blackmarket/1", nil
	case commoditySchema:
		return "http://schemas.elite-markets.net/eddn/commodity/3", nil
	case journalSchema:
		return "http://schemas.elite-markets.net/eddn/journal/1", nil
	case outfittingSchema:
		return "http://schemas.elite-markets.net/eddn/outfitting/2", nil
	case shipyardSchema:
		return "http://schemas.elite-markets.net/eddn/shipyard/2", nil
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
	newHeader.GatewayTimestamp = GenerateUTCDateTime()

	return newHeader, nil
}

// Updates the header to the current time.  Nothing else really needs to
// change.
func (uploader *Uploader) updateHeader() {
	uploader.header.GatewayTimestamp = GenerateUTCDateTime()
}

// GenerateUTCDateTime is a helper function for generating RFC3339Nano time
// strings.
func GenerateUTCDateTime() (timeString string) {
	UTCTime := time.Now().UTC()

	return UTCTime.Format(time.RFC3339)
}

func checkResponse(body io.ReadCloser) (err error) {
	output, _ := ioutil.ReadAll(body)

	if string(output) != "OK" {
		errStr := fmt.Sprintf("Error sending data: %s\n", output)
		return errors.New(errStr)
	}

	return nil
}

func (uploader *Uploader) sendMessage(msg interface{}) (err error) {
	jsonData, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(jsonData)

	resp, err := http.Post(EDDNUploadAddress, "application/json; charset=utf-8",
		buf)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return checkResponse(resp.Body)
}

func validateMessage(schema *gojsonschema.Schema, data interface{}) (err error) {
	loader := gojsonschema.NewGoLoader(data)
	result, err := schema.Validate(loader)

	if err != nil {
		return err
	}

	if !result.Valid() {
		log.Printf("The document is not valid. see errors :\n")
		for _, err := range result.Errors() {
			// Err implements the ResultError interface
			log.Printf("- %s\n", err)
		}

		return errors.New("error validating message")
	}

	return nil
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

	if err = validateMessage(uploader.blackmarketSchema, data); err != nil {
		return err
	}

	return uploader.sendMessage(data)
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

	if err = validateMessage(uploader.commoditySchema, data); err != nil {
		return err
	}

	return uploader.sendMessage(data)
}

// SendJournalDocked sends a Docked message to the EDDN servers.  The
// message should be filled (especially the required fields).  The required
// fields are marked in the journal.go source file.
func (uploader *Uploader) SendJournalDocked(msg *JournalDocked) (err error) {
	schema, err := generateSchema(journalSchema)

	if err != nil {
		return err
	}

	uploader.updateHeader()

	data := &Journal{schema, uploader.header, *msg}

	if err = validateMessage(uploader.journalSchema, data); err != nil {
		return err
	}

	return uploader.sendMessage(data)
}

// SendJournalFSDJump sends a FSDJump message to the EDDN servers.  The
// message should be filled (especially the required fields).  The required
// fields are marked in the journal.go source file.
func (uploader *Uploader) SendJournalFSDJump(msg *JournalFSDJump) (err error) {
	schema, err := generateSchema(journalSchema)

	if err != nil {
		return err
	}

	uploader.updateHeader()

	data := &Journal{schema, uploader.header, *msg}

	if err = validateMessage(uploader.journalSchema, data); err != nil {
		return err
	}

	return uploader.sendMessage(data)
}

// SendJournalScanStar sends a star Scan message to the EDDN servers.  The
// message should be filled (especially the required fields).  The required
// fields are marked in the journal.go source file.
func (uploader *Uploader) SendJournalScanStar(msg *JournalScanStar) (err error) {
	schema, err := generateSchema(journalSchema)

	if err != nil {
		return err
	}

	uploader.updateHeader()

	data := &Journal{schema, uploader.header, *msg}

	if err = validateMessage(uploader.journalSchema, data); err != nil {
		return err
	}

	return uploader.sendMessage(data)
}

// SendJournalScanPlanet sends a planet Scan message to the EDDN servers.  The
// message should be filled (especially the required fields).  The required
// fields are marked in the journal.go source file.
func (uploader *Uploader) SendJournalScanPlanet(msg *JournalScanPlanet) (err error) {
	schema, err := generateSchema(journalSchema)

	if err != nil {
		return err
	}

	uploader.updateHeader()

	data := &Journal{schema, uploader.header, *msg}

	if err = validateMessage(uploader.journalSchema, data); err != nil {
		return err
	}

	return uploader.sendMessage(data)
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

	if err = validateMessage(uploader.outfittingSchema, data); err != nil {
		return err
	}

	return uploader.sendMessage(data)
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

	if err = validateMessage(uploader.shipyardSchema, data); err != nil {
		return err
	}

	return uploader.sendMessage(data)
}
