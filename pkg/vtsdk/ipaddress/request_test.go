package ipaddress_test

import (
	"log"
	"os"
	"testing"

	"github.com/christianrang/find-bad-ip/pkg/vtsdk"
	"github.com/christianrang/find-bad-ip/pkg/vtsdk/ipaddress"
	"github.com/stretchr/testify/assert"
)

var configuration = vtsdk.Configuration{}

func createConfiguration() {
	configuration.SetApiKey(getApiKey())
}

func getApiKey() string {
	apiKey, ok := os.LookupEnv("VT_API_KEY")
	if !ok {
		log.Fatalln("Please export a valid virustotal API Key\n\texport VT_API_KEY='API key here'")
	}

	return apiKey
}

func TestQueryIp(t *testing.T) {
	createConfiguration()

	tests := []struct {
		expected int
		client   *vtsdk.Client
		param    string
	}{
		{
			200,
			vtsdk.CreateClient(configuration),
			"8.8.8.8",
		},
	}
	for _, test := range tests {
		resp, res, err := ipaddress.QueryIp(*test.client, test.param)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", res)
		assert.Equal(t, test.expected, resp.StatusCode(), "they should be equal")
	}
}
