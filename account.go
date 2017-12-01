package bittrex

import (
	"encoding/json"
	"strconv"
)

type JSONBalance struct {
	Currency      string
	Balance       float64
	Available     float64
	Pending       float64
	CryptoAddress string
	Requested     bool
	UUID          string `json:"Uuid"`
}

type JSONDepositAddress struct {
	Currency string
	Address  string
}

type JSONOrderUUID struct {
	AccountID                  string
	OrderUUID                  string `json:"OrderUuid"`
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

type JSONOrderHistory struct {
	OrderUUID         string `json:"OrderUuid"`
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

type JSONAccountHistory struct {
	PaymentUUID    string `json:"PaymentUuid"`
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

//Retrieves balance for defined currencies, if no currencies are defined return all balances
func (b *Bittrex) GetBalance(currencies ...string) ([]JSONBalance, error) {

	jsonBalances := make([]JSONBalance, len(currencies))
	var err error = nil
	if len(currencies) > 0 {
		c := make(chan JSONBalance, len(currencies))
		for _, currency := range currencies {
			endPoint := "account/getbalance?apikey=" + string(b.Key) + "&currency=" + currency
			go func() {
				result, _ := b.Request(endPoint, true)
				jsonBalance := JSONBalance{}
				err = json.Unmarshal([]byte(result), &jsonBalance) //TODO WHAT ARE WE GOING TO DO ABOUT THIS ERROR
				c <- jsonBalance
			}()
		}

		for i := 0; i < len(currencies); i++ {
			res := <-c
			jsonBalances[i] = res
		}
	} else {
		endPoint := "account/getbalances?apikey=" + string(b.Key)
		result, err := b.Request(endPoint, true)

		if err == nil {
			err = json.Unmarshal([]byte(result), &jsonBalances)
		}
	}

	return jsonBalances, err
}

//Retrieves or generates an address for a specific currency. If one does not exist, the call will fail and return ADDRESS_GENERATING until one is available.
func (b *Bittrex) GetDepositAddress(currency string) (JSONDepositAddress, error) {

	endPoint := "account/getdepositaddress?apikey=" + string(b.Key) + "&currency=" + currency
	result, err := b.Request(endPoint, true)
	jSONDepositAddress := JSONDepositAddress{}

	if err == nil {
		err = json.Unmarshal([]byte(result), &jSONDepositAddress)
	}
	return jSONDepositAddress, err
}

//Withdraw fund from account
func (b *Bittrex) Withdraw(currency string, quantity float64, address string, paymentID ...string) (JSONUUID, error) {

	endPoint := "/account/withdraw?apikey=" + string(b.Key) + "&currency=" + currency + "&quantity=" + strconv.FormatFloat(quantity, 'f', -1, 64) + "&address=" + address
	if len(paymentID) > 0 {
		endPoint = endPoint + "&paymentId=" + paymentID[0]
	}

	result, err := b.Request(endPoint, true)
	jSONUUID := JSONUUID{}

	if err == nil {
		err = json.Unmarshal([]byte(result), &jSONUUID)
	}
	return jSONUUID, err
}

//Retrives an order
func (b *Bittrex) GetOrder(UUID string) (JSONOrderUUID, error) {

	endPoint := "account/getorder?apikey=" + string(b.Key) + "&uuid=" + UUID
	result, err := b.Request(endPoint, true)
	jSONOrderUUID := JSONOrderUUID{}

	if err == nil {
		err = json.Unmarshal([]byte(result), &jSONOrderUUID)
	}
	return jSONOrderUUID, err
}

//Retrives order history for defined markets, if no markets are defined retrives for all markets
func (b *Bittrex) GetOrderHistory(markets ...string) ([]JSONOrderHistory, error) {

	//jsonOrderHistories := make([]JSONOrderHistory, len(markets))
	var jsonOrderHistories []JSONOrderHistory

	var err error = nil
	if len(markets) > 0 {
		c := make(chan []JSONOrderHistory, len(markets))
		for _, market := range markets {
			endPoint := "account/getorderhistory?apikey=" + string(b.Key) + "&market=" + market
			go func() {
				result, _ := b.Request(endPoint, true)
				jsonOrderHistory := []JSONOrderHistory{}
				err = json.Unmarshal([]byte(result), &jsonOrderHistory)
				c <- jsonOrderHistory
			}()

		}
		for i := 0; i < len(markets); i++ {
			res := <-c
			jsonOrderHistories = append(jsonOrderHistories, res...)
		}
	} else {
		endPoint := "account/getorderhistory?apikey=" + string(b.Key)
		result, _ := b.Request(endPoint, true)
		if err == nil {
			err = json.Unmarshal([]byte(result), &jsonOrderHistories)
		}
	}

	return jsonOrderHistories, err
}

//Retrives withdrawl history for defined currencies, if no currencies are defined retrives for all currencies
func (b *Bittrex) GetWithdrawHistory(currencies ...string) ([]JSONAccountHistory, error) {

	var jsonAccountHistories []JSONAccountHistory
	var err error = nil
	if len(currencies) > 0 {
		c := make(chan []JSONAccountHistory, len(currencies))
		for _, currency := range currencies {
			endPoint := "account/getwithdrawalhistory?apikey=" + string(b.Key) + "&currency=" + currency
			go func() {
				result, _ := b.Request(endPoint, true)
				jsonAccountHistory := []JSONAccountHistory{}
				err = json.Unmarshal([]byte(result), &jsonAccountHistory)
				c <- jsonAccountHistory
			}()

		}
		for i := 0; i < len(currencies); i++ {
			res := <-c
			jsonAccountHistories = append(jsonAccountHistories, res...)
		}
	} else {
		endPoint := "account/getwithdrawalhistory?apikey=" + string(b.Key)
		result, err := b.Request(endPoint, true)
		if err == nil {
			err = json.Unmarshal([]byte(result), &jsonAccountHistories)
		}
	}

	return jsonAccountHistories, err
}

//Retrives deposit history for defined currencies, if no currencies are defined retrives for all currencies
func (b *Bittrex) GetDepositHistory(currencies ...string) ([]JSONAccountHistory, error) {

	var jsonAccountHistories []JSONAccountHistory
	var err error = nil
	if len(currencies) > 0 {
		c := make(chan []JSONAccountHistory, len(currencies))
		for _, currency := range currencies {
			endPoint := "account/getdeposithistory?apikey=" + string(b.Key) + "&currency=" + currency
			go func() {
				result, _ := b.Request(endPoint, true)
				jsonAccountHistory := []JSONAccountHistory{}
				err = json.Unmarshal([]byte(result), &jsonAccountHistory)
				c <- jsonAccountHistory
			}()
		}

		for i := 0; i < len(currencies); i++ {
			res := <-c
			jsonAccountHistories = append(jsonAccountHistories, res...)

		}
	} else {
		endPoint := "account/getdeposithistory?apikey=" + string(b.Key)
		result, err := b.Request(endPoint, true)
		if err == nil {
			err = json.Unmarshal([]byte(result), &jsonAccountHistories)
		}
	}

	return jsonAccountHistories, err
}
