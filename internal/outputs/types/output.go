package outputTypes

import (
	"encoding/csv"

	table "github.com/calyptia/go-bubble-table"
)

type Output interface {
	CreateRecord() []string
	WriteRow(*csv.Writer, func() []string) error
	CreateTableRow() table.SimpleRow
	OpenGui()
}
