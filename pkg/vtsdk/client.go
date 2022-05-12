package vtsdk

import (
	"net/url"
	"time"

	resty "github.com/go-resty/resty/v2"
)

var VtIpAddressUrl = url.URL{
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
			SetBaseURL(VtIpAddressUrl.String()).
			AddRetryCondition(func(r *resty.Response, err error) bool {
				if r.StatusCode() == 429 {
					time.Sleep(time.Minute)
				}
				return r.StatusCode() == 429
			}),
	}

	return client
}
