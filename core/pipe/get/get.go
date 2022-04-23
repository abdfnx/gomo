package get

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
	pkg      string
	isDelete bool
}

func Get(pkg string, isDelete bool) model {
	st := shared.DefaultStyles()

	return model{
		styles:   st,
		state:    shared.Ready,
		message:  "",
		spinner:  shared.NewSpinner(),
		pkg:      pkg,
		isDelete: isDelete,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case shared.SuccessMsg:
			_msg := " Package added successfully"

			if m.isDelete {
				_msg = " Package deleted successfully"
			}

			m.state = shared.Ready
			head := m.styles.Success.Render("SUCCESS")
			body := m.styles.Subtle.Render(" " + m.styles.Bold.Render(m.pkg) + _msg)
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
			fmt.Println(m.message)
		}
	}

	return lipgloss.NewStyle().SetString(s).String()
}

func spinnerView(m model) string {
	msg := "☄️ Getting "

	if m.isDelete {
		msg = "🧹 Deleting "
	}

	return m.spinner.View() + msg + m.styles.Bold.Render(m.pkg) + " ..."
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
			getCmd := "go get " + m.pkg

			if m.isDelete {
				getCmd += "@none"
			}

			if runtime.GOOS == "windows" {
				cmd = exec.Command("powershell.exe", getCmd)
			} else {
				cmd = exec.Command("bash", "-c", getCmd)
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
