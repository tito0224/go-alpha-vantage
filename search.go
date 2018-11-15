package alphago

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SearchResponse struct {
	Results []Result `json:"bestMatches"`
}

type Result struct {
	Symbol      string `json:"1. symbol"`
	Name        string `json:"2. name"`
	Type        string `json:"3. type"`
	Region      string `json:"4. region"`
	MarketOpen  string `json:"5. marketOpen"`
	MarketClose string `json:"6. marketClose"`
	TimeZone    string `json:"7. timezone"`
	Currency    string `json:"8. currency"`
	MatchScore  string `json:"9. matchScore"`
}

func (client *Client) Search(query string) ([]Result, error) {
	var r []Result

	params := map[string]string{
		"keywords": query,
	}

	res, err := client.ExecuteRequest("SYMBOL_SEARCH", params)
	if err != nil {
		return r, err
	}

	return parseResults(res)
}

func parseResults(res *http.Response) ([]Result, error) {
	var rr SearchResponse
	var r []Result

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(body, &rr)
	if err != nil {
		return r, err
	}

	return rr.Results, nil
}
