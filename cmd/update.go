package cmd

import (
	"fmt"
	"os"

	"github.com/abdfnx/gomo/core/pipe/update"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func UpdateCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update all packages.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := tea.NewProgram(update.Update()).Start(); err != nil {
				fmt.Printf("could not start program: %s\n", err)
				os.Exit(3)
			}
		},
	}

	return cmd
}
