package check

import (
	"errors"
	"fmt"
	"time"

	"github.com/christianrang/hackerfinder/pkg/abuseipdbsdk"
	resty "github.com/go-resty/resty/v2"
)

const _abuseipdbCheckUrlPath = "/api/v2/check/?ipAddress=%s"

func QueryCheck(client abuseipdbsdk.Client, ip string, response *Response) (*resty.Response, error) {

	resp, err := client.Resty.R().
		SetResult(&response).
		Get(fmt.Sprintf(_abuseipdbCheckUrlPath, ip))

	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in sending request %s\n", err))
	}

	// TODO: remove this
	if resp.StatusCode() == 429 {
		time.Sleep(time.Minute)
		resp, err = client.Resty.R().
			SetResult(&response).
			Get(fmt.Sprintf(_abuseipdbCheckUrlPath, ip))

		if err != nil {
			return nil, errors.New(fmt.Sprintf("error in sending request %s\n", err))
		}
	}

	return resp, nil
}
