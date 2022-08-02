package domain

import (
	"fmt"
	"time"

	"github.com/christianrang/hackerfinder/pkg/vtsdk"
	virusTotalErrors "github.com/christianrang/hackerfinder/pkg/vtsdk/errors"
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
			return nil, virusTotalErrors.NewQueryError(domain).Wrap(err)
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
