package cmd

import (
	"fmt"
	"os"

	"github.com/abdfnx/gomo/core/pipe/get"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func GetCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a go package and add it through all your modules.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				if err := tea.NewProgram(get.Get(args[0], false)).Start(); err != nil {
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
