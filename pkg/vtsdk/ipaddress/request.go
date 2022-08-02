package ipaddress

import (
	"fmt"
	"time"

	"github.com/christianrang/hackerfinder/pkg/vtsdk"
	virusTotalErrors "github.com/christianrang/hackerfinder/pkg/vtsdk/errors"
	resty "github.com/go-resty/resty/v2"
)

var _ipAddressPath = "/api/v3/ip_addresses/%s"

func Query(client vtsdk.Client, ip string, response *Response) (*resty.Response, error) {
	var (
		resp *resty.Response
		err  error
	)

queryLoop:
	for {
		resp, err = client.Resty.R().
			SetResult(&response).
			Get(fmt.Sprintf(_ipAddressPath, ip))
		if err != nil {
			return nil, virusTotalErrors.NewQueryError(ip).Wrap(err)
		}

		switch resp.StatusCode() {
		case 200:
			break queryLoop
		case 429:
			time.Sleep(time.Minute)
		}
	}

	return resp, err
}
