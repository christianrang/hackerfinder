package domain

import (
	"os"

	outputTypes "github.com/christianrang/hackerfinder/internal/outputs/types"
	"github.com/christianrang/hackerfinder/pkg/vtsdk/domain"
	"github.com/jedib0t/go-pretty/v6/table"
)

var _tableHeaders = table.Row{
	"Domain",
	"VT M",
	"VT S",
	"VT H",
	"VT U",
}

type Domain struct {
	VirusTotalDomain domain.Response
}

var _ outputTypes.Output = (*Domain)(nil)

func InitializeTable() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(_tableHeaders)
	t.AppendSeparator()
	return t
}

func (_domain Domain) CreateTableRow(t table.Writer) {
	t.AppendRow([]interface{}{
		_domain.VirusTotalDomain.Data.Id,                                      // IP
		_domain.VirusTotalDomain.Data.Attributes.LastAnalysisStats.Malicious,  // VT M
		_domain.VirusTotalDomain.Data.Attributes.LastAnalysisStats.Suspicious, // VT S
		_domain.VirusTotalDomain.Data.Attributes.LastAnalysisStats.Harmless,   // VT H
		_domain.VirusTotalDomain.Data.Attributes.LastAnalysisStats.Undetected, // VT U
	})
	t.AppendSeparator()
}
