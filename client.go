package alphago

import (
	"net/http"
	"net/url"
	"os"
)

type Client struct {
	APIKey    string
	UserAgent string
	BaseURL   string
}

// NewDefaultClient creates and returns a new API Client with default properties set.
func NewDefaultClient(apiKey string) Client {
	var url = os.Getenv("API_URL")
	if url == "" {
		url = "https://www.alphavantage.co/query"
	}

	return NewClient(
		url,
		"Alpha Vantage Go Client (https://github.com/tito0224/go-alpha-vantage)",
		apiKey,
	)
}

// NewClient builds an API Client with the provided URL and user agent.
func NewClient(url string, userAgent string, apiKey string) Client {
	return Client{
		UserAgent: userAgent,
		BaseURL:   url,
		APIKey:    apiKey,
	}
}

func (client *Client) ExecuteRequest(function string, params map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", client.BaseURL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", client.UserAgent)

	q := url.Values{}
	q.Add("function", function)
	q.Add("apikey", client.APIKey)
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	c := http.Client{}
	return c.Do(req)
}
