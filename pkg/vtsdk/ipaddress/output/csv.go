package output

import (
	"encoding/csv"
	"strconv"

	"github.com/christianrang/find-bad-ip/pkg/vtsdk/ipaddress"
)

func CreateHeaders() []string {
	// Number of Columns: 9
	return []string{
		"ID",
		"Last Analysis Malicious Votes",
		"Last Analysis Suspicious Votes",
		"Last Analysis Harmless Votes",
		"Last Analysis Undetected Votes",
		"Threat Level",
		"Reputation",
		"Country",
		"Continent",
	}
}

func CreateRecord(r *ipaddress.Response) []string {
	// Number of Columns: 9
	return []string{
		r.Data.Id,
		strconv.Itoa(r.Data.Attributes.LastAnalysisStats.Malicious),
		strconv.Itoa(r.Data.Attributes.LastAnalysisStats.Suspicious),
		strconv.Itoa(r.Data.Attributes.LastAnalysisStats.Harmless),
		strconv.Itoa(r.Data.Attributes.LastAnalysisStats.Undetected),
		r.ThreatLevel(),
		strconv.Itoa(r.Data.Attributes.Reputation),
		r.Data.Attributes.Country,
		r.Data.Attributes.Continent,
	}
}

func WriteRow(w *csv.Writer, row []string) error {
	defer w.Flush()

	if err := w.Write(row); err != nil {
		return err
	}

	return nil
}
