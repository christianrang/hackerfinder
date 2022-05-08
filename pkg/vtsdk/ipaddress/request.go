package ipaddress

import (
	"errors"
	"fmt"
	"time"

	"github.com/christianrang/find-bad-ip/pkg/vtsdk"
	resty "github.com/go-resty/resty/v2"
)

var _vtIpAddressUrlPath = "/api/v3/ip_addresses/%s"

func QueryIp(client vtsdk.Client, ip string) (*resty.Response, *Response, error) {
	var result *Response

	resp, err := client.Resty.R().
		SetResult(&result).
		Get(fmt.Sprintf(_vtIpAddressUrlPath, ip))
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("error in sending request %s\n", err))
	}

	if resp.StatusCode() == 429 {
		time.Sleep(time.Minute)
		resp, err := client.Resty.R().
			SetResult(&result).
			Get(fmt.Sprintf(_vtIpAddressUrlPath, ip))
		if err != nil {
			return nil, nil, errors.New(fmt.Sprintf("error in sending request %s\n", err))
		}

		return resp, result, err
	}

	return resp, result, err
}
