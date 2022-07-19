package outputTypes

import (
	"encoding/csv"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Output interface {
	CreateRecord() []string
	WriteRow(*csv.Writer, func() []string) error
	CreateTableRow(table.Writer)
}
