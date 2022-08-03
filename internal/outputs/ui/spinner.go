package ui

import (
	"fmt"
	"os"

	table "github.com/calyptia/go-bubble-table"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pkg/browser"
	"golang.org/x/term"
)

type model struct {
	spinner  spinner.Model
	quitting bool
	err      error
	value    string
}

type tableModel struct {
	table table.Model
	rows  [][]string
}

type QueryMsg struct {
	Target string
}

func InitialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{spinner: s}
}

var (
	styleDoc = lipgloss.NewStyle().Padding(1)
)

func InitTableModel(rows []table.Row, rowsString [][]string) tableModel {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		w = 80
		h = 24
	}
	top, right, bottom, left := styleDoc.GetPadding()
	w = w - left - right
	h = h - top - bottom
	t := table.New([]string{"IP", "VT Mal", "VT Sus", "VT Hrmls", "VT Unkn", "VT Rep", "VT Country", "VT Continent", "AbuseIp Conf Score", "AbuseIp Reports", "AbuseIp Users", "AbuseIp Hostnames"}, w, h)

	t.SetRows(rows)

	return tableModel{table: t, rows: rowsString}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case tea.KeyEnter.String():
			target := m.rows[m.table.Cursor()][0]
			browser.OpenURL("https://virustotal.com/gui/ip-address/" + target)
			browser.OpenURL("https://abuseipdb.com/check/" + target)
		}
	}
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m tableModel) View() string {
	return styleDoc.Render(
		m.table.View(),
	)
}

func (m tableModel) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	case QueryMsg:
		m.value = msg.Target
	}
	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	str := fmt.Sprintf("%s Querying for %s... press CTRL+C to cancel", m.spinner.View(), m.value)
	if m.quitting {
		return str + "\n"
	}

	return str
}
