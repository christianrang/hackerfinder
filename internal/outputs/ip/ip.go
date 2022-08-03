package ip

import (
	"fmt"
	"io"

	table "github.com/calyptia/go-bubble-table"
	outputTypes "github.com/christianrang/hackerfinder/internal/outputs/types"
	"github.com/christianrang/hackerfinder/pkg/abuseipdbsdk/check"
	"github.com/christianrang/hackerfinder/pkg/vtsdk/ipaddress"
	"github.com/pkg/browser"
)

var _tableHeaders = []string{
	"IP",
	"VT M",
	"VT S",
	"VT H",
	"VT U",
	"AbuseIp Conf Score",
	"AbuseIp Report Count",
	"AbuseIp Users",
	"AbuseIp Hostnames",
}

type Ip struct {
	AbuseipdbCheck check.Response     `json:"abuseipdb_check"`
	VirusTotalIp   ipaddress.Response `json:"vt_ip_address"`
}

var _ outputTypes.Output = (*Ip)(nil)

func (ip Ip) CreateTableRow() table.SimpleRow {
	return table.SimpleRow{
		ip.VirusTotalIp.Data.Id, // IP
		ip.VirusTotalIp.Data.Attributes.LastAnalysisStats.Malicious,  // VT M
		ip.VirusTotalIp.Data.Attributes.LastAnalysisStats.Suspicious, // VT S
		ip.VirusTotalIp.Data.Attributes.LastAnalysisStats.Harmless,   // VT H
		ip.VirusTotalIp.Data.Attributes.LastAnalysisStats.Undetected, // VT U
		fmt.Sprint(ip.AbuseipdbCheck.Data.AbuseConfidenceScore, "%"), // AbuseIp Conf Score
		ip.AbuseipdbCheck.Data.TotalReports,                          // AbuseIp Report Count
		ip.AbuseipdbCheck.Data.NumDistinctUsers,                      // AbuseIp Users
		ip.AbuseipdbCheck.Data.Hostnames.String(),                    // AbuseIp Hostnames
	}
}

func (ip Ip) OpenGui() {
	// We don't care about these and the goof the UI
	browser.Stderr = io.Discard
	browser.Stdout = io.Discard

	browser.OpenURL(ipaddress.CreateGuiUrl(ip.VirusTotalIp.Data.Id))
	browser.OpenURL(check.CreateGuiUrl(ip.AbuseipdbCheck.Data.IpAddress))
}
