package vm

import "os"

type Client struct {
	AuthorizationToken string
	UserAgent          string
	BaseUrl            string
}

// NewDefaultClient creates and returns a new API Client with default properties set.
func NewDefaultClient() Client {
	var url = os.Getenv("API_URL")
	if url == "" {
		url = "https://www.alphavantage.co/query"
	}

	return Client{
		UserAgent: "Alpha Vantage Go Client (https://github.com/tito0224/go-alpha-vantage)",
		BaseUrl:   url,
	}
}

// NewClient builds an API Client with the provided URL and user agent.
func NewClient(url string, userAgent string) Client {
	return Client{
		UserAgent: userAgent,
		BaseUrl:   url,
	}
}
