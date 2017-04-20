package EDDNClient_test

import (
	"fmt"
	eddn "github.com/mbsmith/EDDNClient"
	"log"
)

var blackmarketMessage *eddn.BlackmarketMessage = &eddn.BlackmarketMessage{"foo", true, 500, "nelder", "wangal", "sometimestamp"}

func ExampleSendBlackmarket() {
	uploader, err := eddn.NewUploader("me", "mysoftware", "1.0")

	if err != nil {
		log.Fatalln(err)
	}

	uploader.SendBlackmarket(blackmarketMessage)
	fmt.Println("done")
	// Output:
	// ahhh
}
