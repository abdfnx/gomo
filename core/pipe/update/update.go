package update

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/abdfnx/gomo/internal/shared"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tidwall/gjson"
)

type model struct {
	styles   shared.Styles
	state    shared.State
	spinner  spinner.Model
	message  string
	errOut   string
	err   	 error
}

func Update() model {
	st := shared.DefaultStyles()

	return model{
		styles:   st,
		state:    shared.Ready,
		message:  "",
		spinner:  shared.NewSpinner(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case shared.SuccessMsg:
			m.state = shared.Ready
			head := m.styles.Success.Render("SUCCESS")
			body := m.styles.Bold.Render(" Packages updated successfully")
			m.message = m.styles.Wrap.Render(head + body)

			return m, tea.Quit

		case shared.ErrorMsg:
			m.state = shared.Ready
			head := m.styles.Error.Render("ERROR")
			body := m.styles.Subtle.Render(" " + m.errOut)
			m.message = m.styles.Wrap.Render(head + body)

			return m, tea.Quit

		case spinner.TickMsg:
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(msg)

			return m, cmd

		default:
			m.state = shared.Loading
			m.message = ""

			return m, tea.Batch(
				run(m),
				spinner.Tick,
			)
	}
}

func (m model) View() string {
	s := ""

	if m.state == shared.Loading {
		s += spinnerView(m)
	} else {
		if m.message != "" {
			fmt.Println(lipgloss.NewStyle().Padding(0, 2).SetString(m.message).String())
			os.Exit(3)
		}
	}

	return lipgloss.NewStyle().Padding(0, 2).SetString(s).String()
}

func spinnerView(m model) string {
	return m.spinner.View() + "🚧 Upgrading ..."
}

func run(m model) tea.Cmd {
	return func() tea.Msg {
		cmdOut := ""
		errOut := ""

		gomoFile, err := os.ReadFile("gomo.json")

		if err != nil {
			fmt.Println(err)
		}

		modules := gjson.Get(string(gomoFile), "modules.#")

		for i := 0; i < int(modules.Int()); i++ {
			mod := gjson.Get(string(gomoFile), "modules." + fmt.Sprint(i)).String()

			var stdout bytes.Buffer
			var stderr bytes.Buffer

			cmd := exec.Command("")
			updateCmd := gjson.Get(string(gomoFile), "cmds.update").String()

			if runtime.GOOS == "windows" {
				cmd = exec.Command("powershell.exe", updateCmd)
			} else {
				cmd = exec.Command("bash", "-c", updateCmd)
			}

			cmd.Dir = mod
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err := cmd.Run()

			if err != nil {
				errOut = stderr.String()
				m.errOut = errOut
			}

			fmt.Print(stdout.String())
		}

		if m.errOut != "" {
			cmdOut = strings.TrimSuffix(m.errOut, "\n")
			return shared.ErrorMsg{}
		} else if m.errOut == "" {
			return shared.SuccessMsg{}
		}

		return shared.SetMsg(cmdOut)
	}
}
