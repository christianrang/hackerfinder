package ip

import (
	"encoding/csv"
	"fmt"
	"strconv"
)

func CreateHeaders() []string {
	// Number of Columns: 12
	return []string{
		"ID",
		"VirusTotal Analysis Malicious Votes",
		"VirusTotal Analysis Suspicious Votes",
		"VirusTotal Last Analysis Harmless Votes",
		"VirusTotal Last Analysis Undetected Votes",
		"VirusTotal Reputation",
		"VirusTotal Country",
		"VirusTotal Continent",
		"Abuseipdb Confidence Score",
		"Abuseipdb Total Reports",
		"Abuseipdb Number of Distinct Users",
		"Abuseipdb Hostnames",
	}
}

func (ip Ip) CreateRecord() []string {
	// Number of Columns: 12
	return []string{
		ip.VtIpAddress.Data.Id,
		strconv.Itoa(ip.VtIpAddress.Data.Attributes.LastAnalysisStats.Malicious),
		strconv.Itoa(ip.VtIpAddress.Data.Attributes.LastAnalysisStats.Suspicious),
		strconv.Itoa(ip.VtIpAddress.Data.Attributes.LastAnalysisStats.Harmless),
		strconv.Itoa(ip.VtIpAddress.Data.Attributes.LastAnalysisStats.Undetected),
		strconv.Itoa(ip.VtIpAddress.Data.Attributes.Reputation),
		ip.VtIpAddress.Data.Attributes.Country,
		ip.VtIpAddress.Data.Attributes.Continent,
		fmt.Sprint(ip.AbuseipdbCheck.Data.AbuseConfidenceScore, "%"),
		strconv.Itoa(ip.AbuseipdbCheck.Data.TotalReports),
		strconv.Itoa(ip.AbuseipdbCheck.Data.NumDistinctUsers),
		ip.AbuseipdbCheck.Data.Hostnames.String(),
	}
}

func WriteRow(w *csv.Writer, row []string) error {
	defer w.Flush()

	if err := w.Write(row); err != nil {
		return err
	}

	return nil
}
