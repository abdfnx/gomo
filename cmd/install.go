package cmd

import (
	"fmt"
	"os"

	"github.com/abdfnx/gomo/core/pipe/download"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func DownloadCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download",
		Short: "Download go packages through all your modules.",
		Aliases: []string{"load"},
		Run: func(cmd *cobra.Command, args []string) {
			if err := tea.NewProgram(download.Download(false)).Start(); err != nil {
				fmt.Printf("could not start program: %s\n", err)
				os.Exit(3)
			}
		},
	}

	return cmd
}
