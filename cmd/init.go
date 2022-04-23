package cmd

import (
	"github.com/abdfnx/gomo/core/initx"
	"github.com/abdfnx/gomo/core/options"
	"github.com/spf13/cobra"
)

var opts = options.InitOptions{
	Module: "",
	Path: "",
}

func InitCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Creates a new gomo.json file in the current folder.",
		Aliases: []string{"."},
		Run: func(cmd *cobra.Command, args []string) {
			if opts.Module != "" {
				initx.InitModule(opts.Module, opts.Path)
			} else {
				initx.Init()
			}
		},
	}

	cmd.Flags().StringVarP(&opts.Module, "mod", "m", "", "Initialize a new module.")
	cmd.Flags().StringVarP(&opts.Path, "path", "p", "", "The Path of the module.")

	return cmd
}
