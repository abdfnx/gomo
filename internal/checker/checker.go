package checker

import (
	"fmt"
	"strings"

	"github.com/mgutz/ansi"
	"github.com/abdfnx/looker"
	"github.com/abdfnx/gomo/api"
	"github.com/abdfnx/gomo/cmd/factory"
)

func Check(buildVersion string) {
	cmdFactory := factory.New()
	stderr := cmdFactory.IOStreams.ErrOut

	latestVersion := api.GetLatest()
	isFromHomebrewTap := isUnderHomebrew()
	isFromUsrBinDir := isUnderUsr()
	isFromAppData := isUnderAppData()

	var command = func() string {
		if isFromHomebrewTap {
			return "brew upgrade gomo"
		} else if isFromUsrBinDir {
			return "curl -fsSL https://bit.ly/gomo-cli | bash"
		} else if isFromAppData {
			return "iwr -useb https://bit.ly/gomo-win | iex"
		}

		return ""
	}

	if buildVersion != latestVersion {
		fmt.Fprintf(stderr, "%s %s → %s\n",
		ansi.Color("There's a new version of ", "yellow") + ansi.Color("gomo", "cyan") + ansi.Color(" is avalaible:", "yellow"),
		ansi.Color(buildVersion, "cyan"),
		ansi.Color(latestVersion, "cyan"))

		if command() != "" {
			fmt.Fprintf(stderr, ansi.Color("To upgrade, run: %s\n", "yellow"), ansi.Color(command(), "black:white"))
		}
	}
}

var gomoExe, _ = looker.LookPath("gomo")

func isUnderHomebrew() bool {
	return strings.Contains(gomoExe, "brew")
}

func isUnderUsr() bool {
	return strings.Contains(gomoExe, "usr")
}

func isUnderAppData() bool {
	return strings.Contains(gomoExe, "AppData")
}
