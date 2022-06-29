package ipaddress

import (
	"github.com/christianrang/hackerfinder/pkg/vtsdk/responses"
)

type Response struct {
	Data  Data            `json:"data,omitempty"`
	Error responses.Error `json:"error,omitempty"`
}

type Data struct {
	Attributes Attributes        `json:"attributes"`
	Type       string            `json:"type"`
	Id         string            `json:"id"`
	Links      map[string]string `json:"links"`
}

type Attributes struct {
	RegionalInternetRegistery string                                   `json:"regional_internet_registery"`
	Jarm                      string                                   `json:"jarm"`
	Network                   string                                   `json:"network"`
	LastHttpsCertificateDate  int                                      `json:"last_https_certificate_date"`
	Tags                      []string                                 `json:"tags"`
	CrowdsourcedContext       []CrowdsourcedContext                    `json:"crowdsourced_context"`
	Country                   string                                   `json:"country"`
	AsOwner                   string                                   `json:"as_owner"`
	LastAnalysisStats         responses.LastAnalysisStats              `json:"last_analysis_stats"`
	Asn                       int                                      `json:"asn"`
	WhoIsDate                 int                                      `json:"who_is_date"`
	LastAnalysisResults       map[string]responses.LastAnalysisResults `json:"last_analysis_results"`
	Reputation                int                                      `json:"reputation"`
	LastModificationDate      int                                      `json:"last_modification_date"`
	TotalVotes                responses.LastAnalysisStats              `json:"total_votes"`
	LastHttpsCertificate      responses.LastHttpsCertificate           `json:"last_https_certificate"`
	Continent                 string                                   `json:"continent"`
	Whois                     string                                   `json:"whois"`
}

type CrowdsourcedContext struct {
	Source    string `json:"source"`
	Title     string `json:"title"`
	Detail    string `json:"detail"`
	Severity  string `json:"severity"`
	Timestamp int    `json:"timestamp"`
}
