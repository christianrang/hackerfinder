package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/christianrang/hackerfinder/pkg/vtsdk"
	resty "github.com/go-resty/resty/v2"
)

var DomainReportUrlPath = "/api/v3/domains/%s"

func Query(client vtsdk.Client, domain string, response *Response) (*resty.Response, error) {
	var (
		resp *resty.Response
		err  error
	)

queryLoop:
	for {
		resp, err = client.Resty.R().
			SetResult(&response).
			Get(fmt.Sprintf(DomainReportUrlPath, domain))
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error querying for %s on VirusTotal: %s", domain, err))
		}

		switch resp.StatusCode() {
		case 200:
			break queryLoop
		case 429:
			time.Sleep(time.Minute)
		default:
			break queryLoop
		}
	}
	return resp, err
}
