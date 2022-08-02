package internal

import (
	"fmt"

	outputsDomain "github.com/christianrang/hackerfinder/internal/outputs/domain"
	outputsHashes "github.com/christianrang/hackerfinder/internal/outputs/hashes"
	"github.com/christianrang/hackerfinder/internal/outputs/ip"
	outputTypes "github.com/christianrang/hackerfinder/internal/outputs/types"
	"github.com/christianrang/hackerfinder/pkg/abuseipdbsdk"
	"github.com/christianrang/hackerfinder/pkg/abuseipdbsdk/check"
	"github.com/christianrang/hackerfinder/pkg/vtsdk"
	"github.com/christianrang/hackerfinder/pkg/vtsdk/domain"
	"github.com/christianrang/hackerfinder/pkg/vtsdk/hashes"
	"github.com/christianrang/hackerfinder/pkg/vtsdk/ipaddress"
)

type Client struct {
	VirusTotalClient *vtsdk.Client
	AbuseipdbClient  *abuseipdbsdk.Client
}

type QueryFunc func(string) (outputTypes.Output, error)

func (client *Client) Query(value string, query QueryFunc) (outputTypes.Output, error) {
	return query(value)
}

func (client *Client) QueryIp(_ip string) (outputTypes.Output, error) {
	var (
		response ip.Ip
		ok       bool
	)

	_, ok = client.VirusTotalClient.Resty.Header["X-Apikey"]
	if ok {
		_, err := ipaddress.Query(*client.VirusTotalClient, _ip, &response.VirusTotalIp)
		if err != nil {
			return nil, err
		}
	} else {
		fmt.Println("warning: No API key was set for VirusTotal. The output may be incomplete.")
	}

	_, ok = client.AbuseipdbClient.Resty.Header["Key"]
	if ok {
		_, err := check.QueryCheck(*client.AbuseipdbClient, _ip, &response.AbuseipdbCheck)
		if err != nil {
			return nil, err
		}
	} else {
		fmt.Println("warning: No API key was set for Abuseipdb. The output may be incomplete.")
	}

	return &response, nil
}

func (client *Client) QueryDomain(_domain string) (outputTypes.Output, error) {
	var (
		response outputsDomain.Domain
		ok       bool
	)

	_, ok = client.VirusTotalClient.Resty.Header["X-Apikey"]
	if ok {
		_, err := domain.Query(*client.VirusTotalClient, _domain, &response.VirusTotalDomain)
		if err != nil {
			return outputsDomain.Domain{}, err
		}
	} else {
		fmt.Println("warning: No API key was set for VirusTotal. The output may be incomplete.")
	}

	return response, nil
}

func (client *Client) QueryHashes(_hash string) (outputTypes.Output, error) {
	var (
		response outputsHashes.Hashes
		ok       bool
	)

	_, ok = client.VirusTotalClient.Resty.Header["X-Apikey"]
	if ok {
		_, err := hashes.Query(*client.VirusTotalClient, _hash, &response.VirusTotalHashes)
		if err != nil {
			return outputsHashes.Hashes{}, err
		}
	} else {
		fmt.Println("warning: No API key was set for VirusTotal. The output may be incomplete.")
	}

	return response, nil
}
