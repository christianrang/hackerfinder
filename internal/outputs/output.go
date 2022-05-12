package output

import (
	"github.com/christianrang/find-bad-ip/pkg/abuseipdbsdk/check"
	"github.com/christianrang/find-bad-ip/pkg/vtsdk/ipaddress"
)

type Ip struct {
	AbuseipdbCheck check.Response
	VtIpAddress    ipaddress.Response
}
