package gomo

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/abdfnx/gomo/cmd"
	"github.com/abdfnx/gomo/cmd/factory"
	"github.com/abdfnx/gomo/core/options"
	"github.com/spf13/cobra"
)

var opts = options.RootOptions{
	Version: false,
}

func Execute(f *factory.Factory, version string, buildDate string) *cobra.Command {
	const desc = `📐 Simple Golang multi modules tool.`

	var rootCmd = &cobra.Command{
		Use:   "gomo <subcommand> [flags]",
		Short:  desc,
		Long: desc,
		SilenceErrors: true,
		Example: `
			# Creates a new gomo.json file in the current folder.
			gomo init

			# Download go packages through all your modules.
			gomo download

			# Update all packages.
			gomo update
		`,
		Annotations: map[string]string{
			"help:tellus": heredoc.Doc(`
				Open an issue at https://github.com/abdfnx/gomo/issues
			`),
		},
		Run: func(cmd *cobra.Command, args []string) {
			if opts.Version {
				fmt.Println("gomo version " + version + " " + buildDate)
			} else {
				cmd.Help()
			}
		},
	}

	rootCmd.SetOut(f.IOStreams.Out)
	rootCmd.SetErr(f.IOStreams.ErrOut)

	cs := f.IOStreams.ColorScheme()

	helpHelper := func(command *cobra.Command, args []string) {
		rootHelpFunc(cs, command, args)
	}

	rootCmd.PersistentFlags().Bool("help", false, "Help for gomo")
	rootCmd.SetHelpFunc(helpHelper)
	rootCmd.SetUsageFunc(rootUsageFunc)
	rootCmd.Flags().BoolVarP(&opts.Version, "version", "v", false, "Print the version of your gomo binary.")

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of your gomo binary.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("gomo version " + version + " " + buildDate)
		},
	}

	rootCmd.AddCommand(
		cmd.AddCMD(),
		cmd.DownloadCMD(),
		cmd.InitCMD(),
		cmd.TidyCMD(),
		cmd.UpdateCMD(),
		versionCmd,
	)

	return rootCmd
}
