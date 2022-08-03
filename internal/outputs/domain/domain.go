package domain

import (
	"io"

	table "github.com/calyptia/go-bubble-table"
	outputTypes "github.com/christianrang/hackerfinder/internal/outputs/types"
	"github.com/christianrang/hackerfinder/pkg/vtsdk/domain"
	"github.com/pkg/browser"
)

var _tableHeaders = []string{
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

func (d Domain) CreateTableRow() table.SimpleRow {
	return table.SimpleRow{
		d.VirusTotalDomain.Data.Id,
		d.VirusTotalDomain.Data.Attributes.LastAnalysisStats.Malicious,
		d.VirusTotalDomain.Data.Attributes.LastAnalysisStats.Suspicious,
		d.VirusTotalDomain.Data.Attributes.LastAnalysisStats.Harmless,
		d.VirusTotalDomain.Data.Attributes.LastAnalysisStats.Undetected,
	}
}

func (d Domain) OpenGui() {
	// We don't care about these and the goof the UI
	browser.Stderr = io.Discard
	browser.Stdout = io.Discard

	browser.OpenURL(domain.CreateGuiUrl(d.VirusTotalDomain.Data.Id))
}
