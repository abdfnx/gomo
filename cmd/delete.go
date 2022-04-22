package cmd

import (
	"fmt"
	"os"

	"github.com/abdfnx/gomo/core/pipe/get"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func DeleteCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a go package through all modules.",
		Aliases: []string{"remove"},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				if err := tea.NewProgram(get.Get(args[0], true)).Start(); err != nil {
					fmt.Printf("could not start program: %s\n", err)
					os.Exit(3)
				}
			} else {
				fmt.Println("Package name is required")
			}
		},
	}

	return cmd
}
