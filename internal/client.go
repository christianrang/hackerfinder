package internal

import (
	"fmt"

	"github.com/christianrang/find-bad-ip/internal/outputs/ip"
	"github.com/christianrang/find-bad-ip/pkg/abuseipdbsdk"
	"github.com/christianrang/find-bad-ip/pkg/abuseipdbsdk/check"
	"github.com/christianrang/find-bad-ip/pkg/vtsdk"
	"github.com/christianrang/find-bad-ip/pkg/vtsdk/ipaddress"
)

type Client struct {
	VirusTotalClient *vtsdk.Client
	AbuseipdbClient  *abuseipdbsdk.Client
}

func (client *Client) QueryIp(_ip string) (*ip.Ip, error) {
	var (
		response ip.Ip
		ok       bool
	)

	_, ok = client.VirusTotalClient.Resty.Header["X-Apikey"]
	if ok {
		_, err := ipaddress.QueryIp(*client.VirusTotalClient, _ip, &response.VirusTotalIp)
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
