package bittrex

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

//JSONResponseBody : JSON high level template for Bittrex API
type JSONResponseBody struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

type JSONUUID struct {
	UUID string `json:"uuid"`
}

const (
	BASE_URL        = "https://bittrex.com/api"
	API_VERSION     = "v1.1"
	DEFAULT_TIMEOUT = 10
)

type Bittrex struct {
	Key    string
	Secret string
	client *http.Client
	config *BittrexConfig
}

type BittrexConfig struct {
	Debug   bool
	Timeout time.Duration
}

func NewBittrex(key string, secret string, config ...BittrexConfig) *Bittrex {

	bittrexConfig := BittrexConfig{}
	if len(config) > 0 {
		bittrexConfig = config[0]
	}
	bittrexConfig.Debug = false
	bittrexConfig.Timeout = time.Second * DEFAULT_TIMEOUT

	client := &http.Client{Timeout: bittrexConfig.Timeout}

	return &Bittrex{Key: key, Secret: secret, config: &bittrexConfig, client: client}
}

func (b *Bittrex) Request(endPoint string, authenticate bool) (json.RawMessage, error) {

	uri := fmt.Sprintf("%s/%s/%s", BASE_URL, API_VERSION, endPoint)

	req, err := http.NewRequest("GET", uri, nil)

	if authenticate {
		nonce := fmt.Sprintf("%d", time.Now().UnixNano())
		q := req.URL.Query()
		q.Add("nonce", nonce)
		req.URL.RawQuery = q.Encode()

		mac := hmac.New(sha512.New, []byte(b.Secret))
		mac.Write([]byte(req.URL.String()))
		apiSign := mac.Sum(nil)
		req.Header.Add("apisign", hex.EncodeToString(apiSign))
	}

	resp, err := b.client.Do(req)
	body := JSONResponseBody{}

	if err == nil {
		defer resp.Body.Close()

		if err == nil {
			decoded := json.NewDecoder(resp.Body)
			err = decoded.Decode(&body)
		}

		if err == nil && !body.Success {
			err = errors.New("Request error: " + body.Message)
		}
	}

	return body.Result, err
}
