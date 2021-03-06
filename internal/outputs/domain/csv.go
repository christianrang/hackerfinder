package domain

import (
	"encoding/csv"
	"strconv"
)

func CreateHeaders() []string {
	// Number of Columns: 6
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
	// Number of Columns: 6
	return []string{
		_domain.VirusTotalDomain.Data.Id,
		strconv.Itoa(_domain.VirusTotalDomain.Data.Attributes.LastAnalysisStats.Malicious),
		strconv.Itoa(_domain.VirusTotalDomain.Data.Attributes.LastAnalysisStats.Suspicious),
		strconv.Itoa(_domain.VirusTotalDomain.Data.Attributes.LastAnalysisStats.Harmless),
		strconv.Itoa(_domain.VirusTotalDomain.Data.Attributes.LastAnalysisStats.Undetected),
		strconv.Itoa(_domain.VirusTotalDomain.Data.Attributes.Reputation),
	}
}

func (_domain Domain) WriteRow(w *csv.Writer, createRecord func() []string) error {
	return WriteRow(w, createRecord)
}

func WriteRow(w *csv.Writer, createRecord func() []string) error {
	defer w.Flush()

	if err := w.Write(createRecord()); err != nil {
		return err
	}

	return nil
}
