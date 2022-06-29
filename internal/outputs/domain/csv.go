package domain

import (
	"encoding/csv"
	"strconv"

	"github.com/christianrang/hackerfinder/pkg/vtsdk/domain"
)

type Domain struct {
	VirusTotalDomain domain.Response
}

func CreateHeaders() []string {
	// Number of Columns: 12
	return []string{
		"ID",
		"VirusTotal Analysis Malicious Votes",
		"VirusTotal Analysis Suspicious Votes",
		"VirusTotal Last Analysis Harmless Votes",
		"VirusTotal Last Analysis Undetected Votes",
		"VirusTotal Reputation",
	}
}

func (_domain Domain) CreateRecord() []string {
	// Number of Columns: 12
	return []string{
		_domain.VirusTotalDomain.Data.Id,
		strconv.Itoa(_domain.VirusTotalDomain.Data.Attributes.LastAnalysisStats.Malicious),
		strconv.Itoa(_domain.VirusTotalDomain.Data.Attributes.LastAnalysisStats.Suspicious),
		strconv.Itoa(_domain.VirusTotalDomain.Data.Attributes.LastAnalysisStats.Harmless),
		strconv.Itoa(_domain.VirusTotalDomain.Data.Attributes.LastAnalysisStats.Undetected),
		strconv.Itoa(_domain.VirusTotalDomain.Data.Attributes.Reputation),
	}
}

func WriteRow(w *csv.Writer, row []string) error {
	defer w.Flush()

	if err := w.Write(row); err != nil {
		return err
	}

	return nil
}
