package internal

import (
	"github.com/christianrang/hackerfinder/pkg/abuseipdbsdk"
	"github.com/christianrang/hackerfinder/pkg/vtsdk"
)

// TODO: move this somewhere better
type Configuration struct {
	Api Api `mapstructure:"api"`
}

type Api struct {
	VTConfig  vtsdk.Configuration        `mapstructure:"virustotal"`
	Abuseipdb abuseipdbsdk.Configuration `mapstructure:"abuseaipdb"`
}

func (api *Api) HasApiKey() bool {
	return api.VTConfig.ApiKey != "" || api.Abuseipdb.ApiKey != ""
}
