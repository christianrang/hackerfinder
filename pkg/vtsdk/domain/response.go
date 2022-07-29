package domain

import (
	"github.com/christianrang/hackerfinder/pkg/vtsdk/responses"
)

type Response struct {
	Data  Data  `json:"data,omitempty"`
	Error Error `json:"error,omitempty"`
}

type Data struct {
	Attributes Attributes `json:"attributes"`
	Id         string     `json:"id"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Attributes struct {
	Categories           Categories                     `json:"categories"`
	CreationDate         int64                          `json:"creation_date"`
	Favicon              Favicon                        `json:"favicon"`
	LastAnalysisResults  responses.LastAnalysisResults  `json:"last_analysis_results"`
	LastAnalysisStats    responses.LastAnalysisStats    `json:"last_analysis_stats"`
	LastDnsRecords       []LastDnsRecords               `json:"last_dns_records"`
	LastDnsRecordsDate   int                            `json:"last_dns_records_date"`
	LastHttpsCertificate responses.LastHttpsCertificate `json:"last_https_certificate"`
	LastModificationDate int                            `json:"last_modification_date"`
	LastUpdateDate       int                            `json:"last_update_date"`
	PopularityRank       map[string]PopularityRank      `json:"popularity_rank"`
	Registrar            string                         `json:"registrar"`
	Reputation           int                            `json:"reputation"`
	Tags                 []string                       `json:"tags"`
	TotalVotes           responses.LastAnalysisStats    `json:"total_votes"`
	WhoIs                string                         `json:"whois"`
	WhoIsDate            int                            `json:"whois_date"`
	Links                Links                          `json:"links"`
	Type                 string                         `json:"type"`
}

type Links struct {
	self string
}

type PopularityRank struct {
	Rank      int `json:"rank"`
	Timestamp int `json:"timestamp"`
}

type Categories map[string]string

type Favicon struct {
	Dhash  string `json:"dhash"`
	RawMd5 string `json:"raw_md5"`
}

type LastDnsRecords struct {
	Expire   int    `json:"expire"`
	Flag     int    `json:"flag"`
	Minimum  int    `json:"minimum"`
	Priority int    `json:"priority"`
	Refresh  int    `json:"refresh"`
	RName    string `json:"rname"`
	Retry    int    `json:"retry"`
	Serial   int    `json:"serial"`
	Tag      string `json:"tag"`
	TTL      int    `json:"ttl"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}
