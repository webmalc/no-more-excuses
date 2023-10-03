package ui

import (
	"strings"
	"webmalc/no-more-excuses/internal/dto"
	"webmalc/no-more-excuses/internal/serializers"
	"webmalc/no-more-excuses/internal/utils"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// cmdModel is the model required by the bubbletea.
type cmdModel struct {
	table table.Model
}

// Init is method required by the bubbletea.
func (m cmdModel) Init() tea.Cmd { return nil } //nolint:gocritic

// Update is method required by the bubbletea.
func (m cmdModel) Update(_ tea.Msg) ( //nolint:gocritic,ireturn
	tea.Model, tea.Cmd,
) {
	return m, tea.Quit
}

// View views the table.
func (m cmdModel) View() string { //nolint:gocritic
	return m.table.View()
}

// The cmd ui.
type Cmd struct {
	model *cmdModel
	apps  map[string]dto.App
}

// Show the configuration.
func (s *Cmd) ShowConfig() {
	pathLen := 20
	cellLen := 11
	s.model.table.SetColumns([]table.Column{
		{Title: "Path", Width: pathLen},
		{Title: "Monday", Width: cellLen},
		{Title: "Tuesday", Width: cellLen},
		{Title: "Wednesday", Width: cellLen},
		{Title: "Thursday", Width: cellLen},
		{Title: "Friday", Width: cellLen},
		{Title: "Saturday", Width: cellLen},
		{Title: "Sunday", Width: cellLen},
	})
	rows := []table.Row{}
	for _, app := range s.apps {
		separator := table.Row{
			strings.Repeat("-", pathLen),
			strings.Repeat("-", cellLen),
			strings.Repeat("-", cellLen),
			strings.Repeat("-", cellLen),
			strings.Repeat("-", cellLen),
			strings.Repeat("-", cellLen),
			strings.Repeat("-", cellLen),
			strings.Repeat("-", cellLen),
		}
		rows = append(rows,
			table.Row{app.Name},
			separator,
		)
		row := table.Row{
			app.Path,
		}
		for _, day := range utils.GetWeek() {
			duration := ""
			if _, ok := app.Weekdays[day]; ok {
				serializer := serializers.NewDurationRange(
					app.Weekdays[day].StartTime, app.Weekdays[day].EndTime, "",
				)
				duration = serializer.Serialize()
			}
			row = append(row, duration)
		}

		rows = append(rows, row, separator, table.Row{""})
	}
	s.model.table.SetRows(rows)
	if _, err := tea.NewProgram(s.model).Run(); err != nil {
		panic(err)
	}
}

func (s *Cmd) initTableStyles() {
	style := table.DefaultStyles()
	style.Header = style.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	style.Selected = style.Selected.
		Foreground(style.Cell.GetForeground()).
		Bold(false)
	s.model.table.SetStyles(style)
}

func NewCmd(apps map[string]dto.App) *Cmd {
	t := table.New(
		table.WithFocused(true),
	)
	cmd := &Cmd{model: &cmdModel{t}, apps: apps}
	cmd.initTableStyles()

	return cmd
}
