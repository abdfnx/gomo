package cmd

import (
	"fmt"

	"github.com/abdfnx/gomo/core/pipe/add"
	"github.com/spf13/cobra"
)

func AddCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a module to gomo.json .",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				add.Add(args[0])
			} else {
				fmt.Println("Module name is required")
			}
		},
	}

	return cmd
}
