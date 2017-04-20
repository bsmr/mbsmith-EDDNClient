package EDDNClient_test

import (
	eddn "github.com/mbsmith/EDDNClient"
	"log"
)

var blackmarketMessage *eddn.BlackmarketMessage = &eddn.BlackmarketMessage{"foo", true, 500, "nelder", "wangal", "2017-04-20T14:17:57.975641Z"}

func ExampleSendJournal() {
	uploader, err := eddn.NewUploader("me", "mysoftware", "1.0")

	if err != nil {
		log.Fatalln(err)
	}

	journalMsg := &eddn.JournalMessage{
		Event:      "FSDJump",
		StarPos:    []float64{33.3, 33.4, 33.5},
		StarSystem: "none",
		Timestamp:  "2017-04-20T14:17:57.975641Z"}

	err = uploader.SendJournal(journalMsg)

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

	err = uploader.SendBlackmarket(blackmarketMessage)

	if err != nil {
		log.Fatalln(err)
	}

	// Output:
}
