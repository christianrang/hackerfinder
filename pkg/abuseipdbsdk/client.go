package abuseipdbsdk

import (
	"net/url"

	"github.com/go-resty/resty/v2"
)

var (
	BaseApiUrl = url.URL{
		Scheme: "https",
		Host:   "api.abuseipdb.com",
	}
	BaseUrl = url.URL{
		Scheme: "https",
		Host:   "abuseipdb.com",
	}
)

type Client struct {
	Resty *resty.Client
}

func CreateClient(c Configuration) *Client {
	client := &Client{
		Resty: resty.New().
			SetHeader("Key", c.ApiKey).
			SetBaseURL(BaseApiUrl.String()),
	}

	return client
}
