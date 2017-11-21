package bittrex

import (
	"encoding/json"
)

//JSONCurrency : JSON currency result structure
type JSONCurrency struct {
	Currency        string
	CurrencyLong    string
	MinConfirmation int
	TxFee           float64
	IsActive        bool
	CoinType        string
	BaseAddress     string
}

//JSONMarket : JSON market result structure
type JSONMarket struct {
	MarketCurrency     string
	BaseCurrency       string
	MarketCurrencyLong string
	BaseCurrencyLong   string
	MinTradeSize       float64
	MarketName         string
	IsActive           bool
	Created            string
}

//JSONTicker : JSON ticker result structure
type JSONTicker struct {
	Bid  float64
	Ask  float64
	Last float64
}

//JSONMarketSummary : JSON market summary result structure
type JSONMarketSummary struct {
	MarketName        string
	High              float64
	Low               float64
	Volume            float64
	Last              float64
	BaseVolume        float64
	TimeStamp         string
	Bid               float64
	Ask               float64
	OpenBuyOrders     int
	OpenSellOrders    int
	PrevDay           float64
	Created           string
	DisplayMarketName string
}

//JSONOrderBookRate : JSON orderbook result structure
type JSONOrderBookRate struct {
	Quantity float64
	Rate     float64
}

//JSONOrderBook : JSON orderbook structure (buy, sell)
type JSONOrderBook struct {
	Buy  []JSONOrderBookRate `json:"buy,omitempty"`
	Sell []JSONOrderBookRate `json:"sell,omitempty"`
}

type JSONMarketHistory struct {
	ID        int `json:"Id"`
	TimeStamp string
	Quantity  float64
	Price     float64
	Total     float64
	FillType  string
	OrderType string
}

type OrderType string

const (
	BUY  OrderType = "buy"
	SELL OrderType = "sell"
	BOTH OrderType = "both"
)

//Retreives all supported currencies with other meta data
func (b *Bittrex) GetCurrencies() ([]JSONCurrency, error) {

	endPoint := "public/getcurrencies"
	result, err := b.Request(endPoint, false)
	jsonCurrencies := []JSONCurrency{}
	if err == nil {
		err = json.Unmarshal([]byte(result), &jsonCurrencies)
	}
	return jsonCurrencies, err
}

//Retreives all open and available trading markets with other meta data
func (b *Bittrex) GetMarkets() ([]JSONMarket, error) {

	endPoint := "public/getmarkets"
	result, err := b.Request(endPoint, false)
	jsonMarkets := []JSONMarket{}
	if err == nil {
		err = json.Unmarshal([]byte(result), &jsonMarkets)
	}
	return jsonMarkets, err
}

//Retreives current tick values for defined market
func (b *Bittrex) GetTicker(market string) (JSONTicker, error) {

	endPoint := "public/getticker?market=" + market
	result, err := b.Request(endPoint, false)
	jsonTicker := JSONTicker{}

	if err == nil {
		err = json.Unmarshal([]byte(result), &jsonTicker)
	}
	return jsonTicker, err
}

//Retrives last 24 hour summary of defined markets.  If no markets are passed retrives summaries for all markets
func (b *Bittrex) GetMarketSummary(markets ...string) ([]JSONMarketSummary, error) {

	jsonMarketSummaries := make([]JSONMarketSummary, len(markets))
	var err error = nil
	if len(markets) > 0 {
		c := make(chan []JSONMarketSummary, len(markets))
		for _, market := range markets {
			endPoint := "public/getmarketsummary?market=" + market
			go func() {
				result, _ := b.Request(endPoint, false)
				jsonMarketSummary := []JSONMarketSummary{}
				err = json.Unmarshal([]byte(result), &jsonMarketSummary) //TODO WHAT ARE WE GOING TO DO ABOUT THIS ERROR
				c <- jsonMarketSummary
			}()
		}

		for i := 0; i < len(markets); i++ {
			res := <-c
			jsonMarketSummaries = append(jsonMarketSummaries, res...)
		}

	} else {
		endPoint := "public/getmarketsummaries"
		result, err := b.Request(endPoint, false)

		if err == nil {
			err = json.Unmarshal([]byte(result), &jsonMarketSummaries)
		}
	}

	return jsonMarketSummaries, err
}

//Retrives the order book for a given market (buy, sell, or both)
func (b *Bittrex) GetOrderBook(market string, orderType OrderType) (JSONOrderBook, error) {

	endPoint := "public/getorderbook?market=" + market + "&type=" + string(orderType)
	result, err := b.Request(endPoint, false)
	jsonOrderBook := JSONOrderBook{}

	if err == nil {
		switch orderType {
		case BUY:
			err = json.Unmarshal([]byte(result), &jsonOrderBook.Buy)

		case SELL:
			err = json.Unmarshal([]byte(result), &jsonOrderBook.Sell)

		case BOTH:
			err = json.Unmarshal([]byte(result), &jsonOrderBook)
		}
	}
	return jsonOrderBook, err
}

//Retrives the latest trades for a specific market
func (b *Bittrex) GetMarketHistory(market string) ([]JSONMarketHistory, error) {
	endPoint := "public/getmarkethistory?market=" + market
	result, err := b.Request(endPoint, false)
	jsonMarketHistories := []JSONMarketHistory{}

	if err == nil {
		err = json.Unmarshal([]byte(result), &jsonMarketHistories)
	}

	return jsonMarketHistories, err
}
