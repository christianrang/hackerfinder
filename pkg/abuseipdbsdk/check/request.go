package check

import (
	"errors"
	"fmt"
	"time"

	"github.com/christianrang/find-bad-ip/pkg/abuseipdbsdk"
	resty "github.com/go-resty/resty/v2"
)

const _abuseipdbCheckUrlPath = "/api/v2/check/?ipAddress=%s"

func QueryCheck(client abuseipdbsdk.Client, ip string) (*resty.Response, *Response, error) {
	var result *Response

	resp, err := client.Resty.R().
		SetResult(&result).
		Get(fmt.Sprintf(_abuseipdbCheckUrlPath, ip))

	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("error in sending request %s\n", err))
	}

	// TODO: remove this
	if resp.StatusCode() == 429 {
		time.Sleep(time.Minute)
		resp, err = client.Resty.R().
			SetResult(&result).
			Get(fmt.Sprintf(_abuseipdbCheckUrlPath, ip))

		if err != nil {
			return nil, nil, errors.New(fmt.Sprintf("error in sending request %s\n", err))
		}
	}

	return resp, result, nil
}
