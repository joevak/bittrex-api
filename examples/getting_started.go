package main

import (
	"fmt"
	"github.com/joevak/bittrex-api"
)

const (
	apiKey    = "XXX"
	apiSecret = "XXX"
)

func main() {

	bittrex := bittrex.NewBittrex(apiKey, apiSecret)
	markets, err := bittrex.GetMarkets()

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("%v", markets)
	}

}
