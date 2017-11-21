package bittrex

import (
	"encoding/json"
	"strconv"
)

type JSONOrder struct {
	UUID              string `json:"uuid"`
	OrderUUID         string `json:"OrderUuid"`
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

//Places buy limit order for one market
func (b *Bittrex) PlaceBuyOrder(market string, quantity float64, rate float64) (JSONUUID, error) {
	endPoint := "market/buylimit?apikey=" + string(b.Key) + "&market=" + market + "&quantity=" + strconv.FormatFloat(quantity, 'f', -1, 64) + "&rate=" + strconv.FormatFloat(rate, 'f', -1, 64)
	result, err := b.Request(endPoint, true)
	jsonUUID := JSONUUID{}

	if err == nil {
		err = json.Unmarshal([]byte(result), &jsonUUID)
	}

	return jsonUUID, err
}

//Places sell limit order for a specific market
func (b *Bittrex) PlaceSellOrder(market string, quantity float64, rate float64) (JSONUUID, error) {
	endPoint := "market/selllimit?apikey=" + string(b.Key) + "&market=" + market + "&quantity=" + strconv.FormatFloat(quantity, 'f', -1, 64) + "&rate=" + strconv.FormatFloat(rate, 'f', -1, 64)
	result, err := b.Request(endPoint, true)
	jsonUUID := JSONUUID{}

	if err == nil {
		err = json.Unmarshal([]byte(result), &jsonUUID)
	}
	return jsonUUID, err
}

//Cancels buy or sell order by UUID
func (b *Bittrex) CancelOrder(UUID string) error {
	endPoint := "market/cancel?apikey=" + string(b.Key) + "&uuid=" + UUID
	_, err := b.Request(endPoint, true)

	return err
}

//Retrives all open order for specified market.
func (b *Bittrex) GetOrders(market string) ([]JSONOrder, error) {

	endPoint := "market/getopenorders?apikey=" + string(b.Key) + "&market=" + market
	result, err := b.Request(endPoint, true)
	jSONOrders := []JSONOrder{}

	if err == nil {
		err = json.Unmarshal([]byte(result), &jSONOrders)
	}
	return jSONOrders, err
}
