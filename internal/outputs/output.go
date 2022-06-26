package output

import (
	"github.com/christianrang/hackerfinder/pkg/abuseipdbsdk/check"
	"github.com/christianrang/hackerfinder/pkg/vtsdk/ipaddress"
)

type Ip struct {
	AbuseipdbCheck check.Response
	VtIpAddress    ipaddress.Response
}
