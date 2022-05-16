package domain

import (
	vtdomain "github.com/christianrang/find-bad-ip/pkg/vtsdk/domain"
)

type Domain struct {
	VirusTotal vtdomain.Response `virustotal`
}
