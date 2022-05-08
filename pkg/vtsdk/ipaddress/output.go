package ipaddress

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
)

func (r Response) String() string {
	return fmt.Sprintf("%s: %d, %d, %d",
		r.Data.Id,
		r.Data.Attributes.LastAnalysisStats.Malicious,
		r.Data.Attributes.LastAnalysisStats.Suspicious,
		r.Data.Attributes.LastAnalysisStats.Harmless,
	)
}

func (r Response) Table(t table.Writer) {
	t.AppendRow([]interface{}{
		r.Data.Id,
		r.Data.Attributes.LastAnalysisStats.Malicious,
		r.Data.Attributes.LastAnalysisStats.Suspicious,
		r.Data.Attributes.LastAnalysisStats.Harmless,
		r.ThreatLevel(),
	})

	t.AppendSeparator()
}

func (r Response) ThreatLevel() string {
	malicious := r.Data.Attributes.LastAnalysisStats.Malicious
	suspicious := r.Data.Attributes.LastAnalysisStats.Suspicious
	harmless := r.Data.Attributes.LastAnalysisStats.Harmless
	undetected := r.Data.Attributes.LastAnalysisStats.Undetected

	if malicious > suspicious && malicious > harmless && malicious > undetected {
		return "malicious"
	}
	if suspicious > malicious && suspicious > harmless && suspicious > undetected {
		return "suspicious"
	}
	if harmless > malicious && harmless > suspicious && harmless > undetected {
		return "harmless"
	}
	if undetected > malicious && undetected > suspicious && undetected > harmless {
		return "undetected"
	}
	return ""
}
