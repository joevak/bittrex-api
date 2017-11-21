# bitterex-api
A Go API wrapper implementation of the Bitterex API v1.1

# Import
~~~ go
  import "github.com/joevak/bittrex-api"
~~~
  
# Usage
~~~ go
package main

import (
	"fmt"
	"github.com/joevak/bittrex-api""
)

const (
	apiKey    = "API_KEY"
	apiSecret = "API_SECRET"
)

func main() {
    bittrex := bittrex.NewBittrex(apiKey, apiSecret)

    markets, err := bittrex.GetMarketSummary()

    if err != nil {
        fmt.Printf("%s", err)
    } else {
        fmt.Printf("%s", markets)
    }
}
~~~

# Functions

## Public
  ~~~ go 
  func (b *Bittrex)  GetCurrencies() ([]JSONCurrency, error) 
  ~~~
    Retrives the latest trades for a specific market
  ~~~ go
  func (b *Bittrex)  GetMarkets() ([]JSONMarket, error)
  ~~~
    Retreives all open and available trading markets with other meta data
  ~~~ go
  func (b *Bittrex) GetTicker(market string) (JSONTicker, error)
  ~~~
    Retreives current tick values for defined market
  ~~~ go
  func (b *Bittrex) GetMarketSummary(markets ...string) ([]JSONMarketSummary, error)
  ~~~
    Retrives last 24 hour summary of defined markets.  If no markets are passed retrives summaries for all markets
  ~~~ go
  func (b *Bittrex) GetOrderBook(market string, orderType OrderType) (JSONOrderBook, error)
  ~~~
    Retrives the order book for a given market (buy, sell, or both)
  ~~~ go
  func (b *Bittrex) GetMarketHistory(market string) ([]JSONMarketHistory, error)
  ~~~
    Retrives the latest trades for a specific market
  
## Market
##### requires API key

  ~~~ go
  func (b *Bittrex) PlaceBuyOrder(market string, quantity float64, rate float64) (JSONUUID, error)
  ~~~
    Places buy limit order for one market
  ~~~ go
  func (b *Bittrex) PlaceSellOrder(market string, quantity float64, rate float64) (JSONUUID, error)
  ~~~
    Places sell limit order for a specific market
  ~~~ go
  func (b *Bittrex) CancelOrder(UUID string) error
  ~~~
    Cancels buy or sell order
  ~~~ go
  func (b *Bittrex) GetOrders(market string) ([]JSONOrder, error)
  ~~~
    Retrives all open order for specified market.

## Account
##### requires API key

  ~~~ go
  func (b *Bittrex) GetBalance(currencies ...string) ([]JSONBalance, error)
  ~~~
    Retrieves balance for defined currencies, if no currencies are defined return all balances
  ~~~ go
  func (b *Bittrex) GetDepositAddress(currency string) (JSONDepositAddress, error)
  ~~~
    Retrieves or generates an address for a specific currency. If one does not exist, the call will fail and return ADDRESS_GENERATING until one is available.
  ~~~ go
  func (b *Bittrex) Withdraw(currency string, quantity float64, address string, paymentID ...string) (JSONUUID, error)
  ~~~
    Withdraw fund from account
  ~~~ go
  func (b *Bittrex) GetWithdrawHistory(currencies ...string) ([]JSONAccountHistory, error)
  ~~~
    Retrives withdrawl history for defined currencies, if no currencies are defined retrives for all currencies
  ~~~ go
  func (b *Bittrex) GetDepositHistory(currencies ...string) ([]JSONAccountHistory, error)
  ~~~
    Retrives deposit history for defined currencies, if no currencies are defined retrives for all currencies
  ~~~ go
  func (b *Bittrex) GetOrder(UUID string) (JSONUUIDOrder, error)
  ~~~
    Retrives an order
  ~~~ go
  func (b *Bittrex) GetOrderHistory(markets ...string) ([]JSONOrderHistory, error)
  ~~~
    Retrives order history for defined markets, if no markets are defined retrives for all markets

# Types

#### JSONBalance
~~~ go
type JSONBalance struct {
    Currency      string
    Balance       float64
    Available     float64
    Pending       float64
    CryptoAddress string
    Requested     bool
    UUID          string
}
~~~
#### JSONDepositAddress
~~~ go
type JSONDepositAddress struct {
    Currency string
    Address  string
}
~~~
#### JSONOrderUUID
~~~ go
type JSONOrderUUID struct {
    AccountID                  string
    OrderUUID                  string 
    Exchange                   string
    Type                       string
    Quantity                   float64
    QuantityRemaining          float64
    Limit                      float64
    Reserved                   float64
    ReserveRemaining           float64
    CommissionReserver         float64
    CommissionReserveRemaining float64
    CommissionPaid             float64
    Price                      float64
    PricePerUnit               float64
    Opened                     string
    Closed                     string
    IsOpen                     bool
    Sentinel                   string
    CancelInitiated            bool
    ImmediateOrCancel          bool
    IsConditional              bool
    Condition                  string
    ConditionTarget            string
}
~~~
#### JSONOrderHistory
~~~ go
type JSONOrderHistory struct {
    OrderUUID         string
    Exchange          string
    TimeStamp         string
    OrderType         string
    Quantity          float64
    QuantityRemaining float64
    Limit             float64
    Commission        float64
    Price             float64
    PricePerUnit      float64
    ImmediateOrCancel bool
    IsConditional     bool
    Condition         string
    ConditionTarget   string
}
~~~
#### JSONAccountHistory
~~~ go
type JSONAccountHistory struct {
    PaymentUUID    string 
    Currency       string
    Amount         float64
    Address        string
    Opened         string
    Authorized     bool
    PendingPayment bool
    TxCost         float64
    TxID           string
    Canceled       bool
    InvalidAddress bool
}
~~~
#### JSONCurrency
~~~ go
type JSONCurrency struct {
    Currency        string
    CurrencyLong    string
    MinConfirmation int
    TxFee           float64
    IsActive        bool
    CoinType        string
    BaseAddress     string
} 
~~~
#### JSONMarket
~~~ go
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
~~~
#### JSONTicker
~~~ go
type JSONTicker struct {
    Bid  float64
    Ask  float64
    Last float64
}
~~~
#### JSONMarketSummary
~~~ go
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
~~~
#### JSONOrderBookRate
~~~ go
type JSONOrderBookRate struct {
    Quantity float64
    Rate     float64
}
~~~
#### JSONOrderBook
~~~ go
type JSONOrderBook struct {
    Buy  []JSONOrderBookRate
    Sell []JSONOrderBookRate
}
~~~
#### JSONMarketHistory
~~~ go
type JSONMarketHistory struct {
    ID        int 
    TimeStamp string
    Quantity  float64
    Price     float64
    Total     float64
    FillType  string
    OrderType string
}
~~~
#### OrderType
~~~ go
type OrderType string

const (
    BUY  OrderType = "buy"
    SELL OrderType = "sell"
    BOTH OrderType = "both"
)
~~~
#### JSONOrder
~~~ go
type JSONOrder struct {
    UUID              string
    OrderUUID         string 
    Exchange          string
    OrderType         string
    Quantity          float64
    QuantityRemaning  float64
    Limit             float64
    CommissionPaid    float64
    Price             float64
    PricePerUnit      float64
    Opened            string
    closed            string
    CancelInitiated   bool
    ImmediateOrCancel bool
    IsConditional     bool
    Condition         string
    ConditionTarget   string
}
~~~
  
