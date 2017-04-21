package EDDNClient_test

import (
	eddn "github.com/mbsmith/EDDNClient"
	"log"
)

func ExampleSendJournalFSDJump() {
	uploader, err := eddn.NewUploader("me", "mysoftware", "1.0")

	if err != nil {
		log.Fatalln(err)
	}

	journalMsg := &eddn.JournalFSDJump{
		Event:      "FSDJump",
		StarPos:    []float64{33.3, 33.4, 33.5},
		StarSystem: "none",
		Timestamp:  eddn.GenerateUTCDateTime()}

	err = uploader.SendJournalFSDJump(journalMsg)

	if err != nil {
		log.Fatalln(err)
	}

	// Output:
}

func ExampleSendBlackmarket() {
	uploader, err := eddn.NewUploader("me", "mysoftware", "1.0")

	if err != nil {
		log.Fatalln(err)
	}

	blackmarketMessage := &eddn.BlackmarketMessage{
		Name:        "usscargoblackbox",
		Prohibited:  false,
		SellPrice:   1806,
		SystemName:  "Pleione",
		StationName: "Stargazer",
		Timestamp:   eddn.GenerateUTCDateTime()}

	err = uploader.SendBlackmarket(blackmarketMessage)

	if err != nil {
		log.Fatalln(err)
	}

	// Output:
}
