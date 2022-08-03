package hashes

import (
	"fmt"
	"time"

	"github.com/christianrang/hackerfinder/pkg/vtsdk"
	virusTotalErrors "github.com/christianrang/hackerfinder/pkg/vtsdk/errors"
	resty "github.com/go-resty/resty/v2"
)

var (
	_hashesPath = "/api/v3/files/"
	hashGuiPath = "/gui/file/%s"
)

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
			return nil, virusTotalErrors.NewQueryError(hash).Wrap(err)
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

func CreateGuiUrl(target string) string {
	return vtsdk.VirusTotalBaseUrl.String() + fmt.Sprintf(hashGuiPath, target)
}
