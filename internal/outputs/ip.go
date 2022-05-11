package outputs

import (
	"fmt"
	"os"

	"github.com/christianrang/find-bad-ip/pkg/abuseipdbsdk/check"
	"github.com/christianrang/find-bad-ip/pkg/vtsdk/ipaddress"
	"github.com/jedib0t/go-pretty/v6/table"
)

var _tableHeaders = table.Row{
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
	VtIpAddress    ipaddress.Response `json:"vt_ip_address"`
}

func InitializeTable() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(_tableHeaders)
	t.AppendSeparator()
	return t
}

func (ip Ip) CreateTableRow(t table.Writer) {
	t.AppendRow([]interface{}{
		ip.VtIpAddress.Data.Id, // IP
		ip.VtIpAddress.Data.Attributes.LastAnalysisStats.Malicious,   // VT M
		ip.VtIpAddress.Data.Attributes.LastAnalysisStats.Suspicious,  // VT S
		ip.VtIpAddress.Data.Attributes.LastAnalysisStats.Harmless,    // VT H
		ip.VtIpAddress.Data.Attributes.LastAnalysisStats.Undetected,  // VT U
		fmt.Sprint(ip.AbuseipdbCheck.Data.AbuseConfidenceScore, "%"), // AbuseIp Conf Score
		ip.AbuseipdbCheck.Data.TotalReports,                          // AbuseIp Report Count
		ip.AbuseipdbCheck.Data.NumDistinctUsers,                      // AbuseIp Users
		ip.AbuseipdbCheck.Data.Hostnames.String(),                    // AbuseIp Hostnames
	})
	t.AppendSeparator()
}
