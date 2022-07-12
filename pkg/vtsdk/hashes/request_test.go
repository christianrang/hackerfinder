package hashes_test

import (
	"log"
	"os"
	"testing"

	"github.com/christianrang/hackerfinder/pkg/vtsdk"
	"github.com/christianrang/hackerfinder/pkg/vtsdk/hashes"
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

func TestQuery(t *testing.T) {
	createConfiguration()

	tests := []struct {
		expected int
		client   *vtsdk.Client
		param    string
		response *hashes.Response
	}{
		{
			200,
			vtsdk.CreateClient(configuration),
			"74768564ea2ac673e57e937f80c895c81d015e99a72544efa5a679d729c46d5f",
			&hashes.Response{},
		},
	}
	for _, test := range tests {
		resp, err := hashes.Query(*test.client, test.param, test.response)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, test.expected, resp.StatusCode(), "should get 200")
	}
}
