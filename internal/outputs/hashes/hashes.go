package hashes

import (
	"io"

	table "github.com/calyptia/go-bubble-table"
	outputTypes "github.com/christianrang/hackerfinder/internal/outputs/types"
	"github.com/christianrang/hackerfinder/pkg/vtsdk/hashes"
	"github.com/pkg/browser"
)

var _tableHeaders = []string{
	"Hashes",
	"VT M",
	"VT S",
	"VT H",
	"VT U",
}

type Hashes struct {
	VirusTotalHashes hashes.Response
}

var _ outputTypes.Output = (*Hashes)(nil)

func (h Hashes) CreateTableRow() table.SimpleRow {
	return table.SimpleRow{
		h.VirusTotalHashes.Data.Id,                                      // IP
		h.VirusTotalHashes.Data.Attributes.LastAnalysisStats.Malicious,  // VT M
		h.VirusTotalHashes.Data.Attributes.LastAnalysisStats.Suspicious, // VT S
		h.VirusTotalHashes.Data.Attributes.LastAnalysisStats.Harmless,   // VT H
		h.VirusTotalHashes.Data.Attributes.LastAnalysisStats.Undetected, // VT U
	}
}

func (h Hashes) OpenGui() {
	// We don't care about these and the goof the UI
	browser.Stderr = io.Discard
	browser.Stdout = io.Discard

	browser.OpenURL(hashes.CreateGuiUrl(h.VirusTotalHashes.Data.Id))
}
