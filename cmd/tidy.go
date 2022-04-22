package cmd

import (
	"fmt"
	"os"

	"github.com/abdfnx/gomo/core/pipe/download"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func TidyCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tidy",
		Short: "Add any missing packages necessary to build your modules.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := tea.NewProgram(download.Download(true)).Start(); err != nil {
				fmt.Printf("could not start program: %s\n", err)
				os.Exit(3)
			}
		},
	}

	return cmd
}
