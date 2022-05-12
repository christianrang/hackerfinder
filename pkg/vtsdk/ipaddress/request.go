package ipaddress

import (
	"errors"
	"fmt"

	"github.com/christianrang/find-bad-ip/pkg/vtsdk"
	resty "github.com/go-resty/resty/v2"
)

var _vtIpAddressUrlPath = "/api/v3/ip_addresses/%s"

func QueryIp(client vtsdk.Client, ip string, response *Response) (*resty.Response, error) {
	resp, err := client.Resty.R().
		SetResult(&response).
		Get(fmt.Sprintf(_vtIpAddressUrlPath, ip))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in sending request %s\n", err))
	}

	return resp, err
}
