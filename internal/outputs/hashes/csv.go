package hashes

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

func (_hashes Hashes) CreateRecord() []string {
	// Number of Columns: 6
	return []string{
		_hashes.VirusTotalHashes.Data.Id,
		strconv.Itoa(_hashes.VirusTotalHashes.Data.Attributes.LastAnalysisStats.Malicious),
		strconv.Itoa(_hashes.VirusTotalHashes.Data.Attributes.LastAnalysisStats.Suspicious),
		strconv.Itoa(_hashes.VirusTotalHashes.Data.Attributes.LastAnalysisStats.Harmless),
		strconv.Itoa(_hashes.VirusTotalHashes.Data.Attributes.LastAnalysisStats.Undetected),
		strconv.Itoa(_hashes.VirusTotalHashes.Data.Attributes.Reputation),
	}
}

func (_hashes Hashes) WriteRow(w *csv.Writer, createRecord func() []string) error {
	return WriteRow(w, createRecord)
}

func WriteRow(w *csv.Writer, createRecord func() []string) error {
	defer w.Flush()

	if err := w.Write(createRecord()); err != nil {
		return err
	}

	return nil
}
