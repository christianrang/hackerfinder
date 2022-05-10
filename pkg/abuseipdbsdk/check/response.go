package check

type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	IpAddress            string    `json:"ipAddress"`
	IsPublic             bool      `json:"isPublic"`
	IpVersion            int       `json:"IpVersion"`
	IsWhitelisted        bool      `json:"isWhitelisted"`
	AbuseConfidenceScore int       `json:"abuseConfidenceScore"`
	CountryCode          string    `json:"countryCode"`
	CountryName          string    `json:"countryName"`
	UsageType            string    `json:"usageType"`
	Isp                  string    `json:"isp"`
	Domain               string    `json:"domain"`
	Hostnames            []string  `json:"hostnames"`
	ToalReports          int       `json:"toalReports"`
	NumDistinctUsers     int       `json:"numDistinctUsers"`
	LastReportedAt       string    `json:"lastReportedAt"`
	Reports              []Reports `json:"reports"`
}

type Reports struct {
	ReportedAt          string `json:"reportedAt"`
	Comment             string `json:"comment"`
	Categories          []int  `json:"categories"`
	ReporterId          int    `json:"reporterId"`
	ReporterCountryCode string `json:"reporterCountryCode"`
	ReporterCountryName string `json:"reporterCountryName"`
}
