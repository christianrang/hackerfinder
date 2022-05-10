package check_test

import (
	"log"
	"os"
	"testing"

	"github.com/christianrang/find-bad-ip/pkg/abuseipdbsdk"
	"github.com/christianrang/find-bad-ip/pkg/abuseipdbsdk/check"
	"github.com/stretchr/testify/assert"
)

var configuration = abuseipdbsdk.Configuration{}

func createConfiguration() {
	configuration.SetApiKey(getApiKey())
}

func getApiKey() string {
	apiKey, ok := os.LookupEnv("ABUSEIPDB_API_KEY")
	if !ok {
		log.Fatalln("Please export a valid abuseaipdb API Key\n\texport ABUSEIPDB_API_KEY='API key here'")
	}

	return apiKey
}

func TestQueryCheck(t *testing.T) {
	createConfiguration()

	tests := []struct {
		expected int
		client   *abuseipdbsdk.Client
		param    string
	}{
		{
			200,
			abuseipdbsdk.CreateClient(configuration),
			"8.8.8.8",
		},
	}
	for _, test := range tests {
		resp, res, err := check.QueryCheck(*test.client, test.param)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", res)
		t.Logf("%d", resp.StatusCode())
		assert.Equal(t, test.expected, resp.StatusCode(), "they should be equal")
	}
}
