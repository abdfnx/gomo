package cmd

import (
	"github.com/abdfnx/gomo/core/pipe/add"
	"github.com/spf13/cobra"
)

func AddCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a module to gomo.json .",
		Run: func(cmd *cobra.Command, args []string) {
			add.Add(args[0])
		},
	}

	return cmd
}
