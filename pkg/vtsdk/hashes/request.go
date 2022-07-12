package hashes

import (
	"errors"
	"fmt"
	"time"

	"github.com/christianrang/hackerfinder/pkg/vtsdk"
	resty "github.com/go-resty/resty/v2"
)

var _hashesPath = "/api/v3/files/"

func Query(client vtsdk.Client, hash string, response *Response) (*resty.Response, error) {
	var (
		resp *resty.Response
		err  error
	)

queryLoop:
	for {
		resp, err = client.Resty.R().
			SetResult(&response).
			Get(_hashesPath + hash)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error in sending request %s: ", err))
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
