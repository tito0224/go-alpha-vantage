package alphago

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type QuoteResponse struct {
	Quote Quote `json:"Global Quote"`
}

type Quote struct {
	Symbol        string `json:"01. symbol"`
	Open          string `json:"02. open"`
	High          string `json:"03. high"`
	Low           string `json:"04. low"`
	Price         string `json:"05. price"`
	Volume        string `json:"06. volume"`
	LastDate      string `json:"07. latest trading day"`
	PrevClose     string `json:"08. previous close"`
	Change        string `json:"09. change"`
	ChangePercent string `json:"10. change percent"`
}

func (client *Client) GetQuote(symbol string) (Quote, error) {
	var q Quote

	params := map[string]string{
		"symbol": symbol,
	}

	res, err := client.ExecuteRequest("GLOBAL_QUOTE", params)
	if err != nil {
		return q, err
	}

	return parseQuote(res)
}

func parseQuote(res *http.Response) (Quote, error) {
	var qr QuoteResponse
	var q Quote

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return q, err
	}

	err = json.Unmarshal(body, &qr)
	if err != nil {
		return q, err
	}

	return qr.Quote, nil
}
