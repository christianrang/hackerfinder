package vtsdk

import (
	"net/url"

	resty "github.com/go-resty/resty/v2"
)

var VirusTotalBaseUrl = url.URL{
	Scheme: "https",
	Host:   "www.virustotal.com",
}

type Client struct {
	Resty *resty.Client
}

func CreateClient(c Configuration) *Client {
	client := &Client{
		Resty: resty.New().
			SetHeader("x-apikey", c.ApiKey).
			SetBaseURL(VirusTotalBaseUrl.String()),
	}

	return client
}
