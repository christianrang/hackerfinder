package internal

import (
	"github.com/christianrang/find-bad-ip/pkg/abuseipdbsdk"
	"github.com/christianrang/find-bad-ip/pkg/vtsdk"
)

// TODO: move this somewhere better
type Configuration struct {
	Api Api `mapstructure:"api"`
}

type Api struct {
	VTConfig  vtsdk.Configuration        `mapstructure:"virustotal"`
	Abuseipdb abuseipdbsdk.Configuration `mapstructure:"abuseaipdb"`
}
