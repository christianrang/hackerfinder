package internal

import (
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
	var response ip.Ip
	_, err := ipaddress.QueryIp(*client.VirusTotalClient, _ip, &response.VirusTotalIp)
	if err != nil {
		return nil, err
	}

	_, err = check.QueryCheck(*client.AbuseipdbClient, _ip, &response.AbuseipdbCheck)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
