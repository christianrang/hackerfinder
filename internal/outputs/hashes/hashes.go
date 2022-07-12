package hashes

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

var _tableHeaders = table.Row{
	"Hashes",
	"VT M",
	"VT S",
	"VT H",
	"VT U",
}

func InitializeTable() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(_tableHeaders)
	t.AppendSeparator()
	return t
}

func (_hashes Hashes) CreateTableRow(t table.Writer) {
	t.AppendRow([]interface{}{
		_hashes.VirusTotalHashes.Data.Id,                                      // IP
		_hashes.VirusTotalHashes.Data.Attributes.LastAnalysisStats.Malicious,  // VT M
		_hashes.VirusTotalHashes.Data.Attributes.LastAnalysisStats.Suspicious, // VT S
		_hashes.VirusTotalHashes.Data.Attributes.LastAnalysisStats.Harmless,   // VT H
		_hashes.VirusTotalHashes.Data.Attributes.LastAnalysisStats.Undetected, // VT U
	})
	t.AppendSeparator()
}
