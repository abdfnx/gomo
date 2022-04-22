package cmd

import (
	"github.com/abdfnx/gomo/core/initx"
	"github.com/spf13/cobra"
)

func InitCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Creates a new gomo.json file in the current folder.",
		Aliases: []string{"."},
		Run: func(cmd *cobra.Command, args []string) {
			initx.Init()
		},
	}

	return cmd
}
