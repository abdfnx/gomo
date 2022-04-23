package shared

import (
	"github.com/abdfnx/gomo/constants"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	BaseStyle,
	Wrap,
	Doc,
	Subtle,
	Success,
	Error,
	Bold lipgloss.Style
}

func DefaultStyles() Styles {
	s := Styles{}

	s.Wrap = lipgloss.NewStyle().Width(75)
	s.Subtle = lipgloss.NewStyle().
		Foreground(constants.SUBTITLE_COLOR)
	s.Success = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFF")).Background(lipgloss.Color(constants.GREEN_COLOR)).PaddingLeft(1).PaddingRight(1)
	s.Error = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFF")).Background(lipgloss.Color(constants.RED_COLOR)).PaddingLeft(1).PaddingRight(1)
	s.Bold = lipgloss.NewStyle().Bold(true)
	s.BaseStyle = lipgloss.NewStyle()
	s.Doc = lipgloss.NewStyle().Align(lipgloss.Center)

	return s
}

var (
	spinnerStyle = lipgloss.NewStyle().Foreground(constants.GRAY_COLOR)
)

func NewSpinner() spinner.Model {
	s := spinner.NewModel()

	s.Spinner = spinner.Dot
	s.Style = spinnerStyle

	return s
}
