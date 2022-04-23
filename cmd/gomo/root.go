package gomo

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/abdfnx/gomo/cmd"
	"github.com/abdfnx/gomo/cmd/factory"
	"github.com/abdfnx/gomo/core/options"
	"github.com/abdfnx/gomo/core/pipe/download"
	tea "github.com/charmbracelet/bubbletea"
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
			gomo

			# Get a go package and add it through all modules.
			gomo get github.com/gorilla/mux

			# Delete a go package through all modules.
			gomo delete github.com/example/example1

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
				if err := tea.NewProgram(download.Download(false)).Start(); err != nil {
					fmt.Printf("could not start program: %s\n", err)
					os.Exit(3)
				}
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
		cmd.DeleteCMD(),
		cmd.GetCMD(),
		cmd.InitCMD(),
		cmd.TidyCMD(),
		cmd.UpdateCMD(),
		versionCmd,
	)

	return rootCmd
}
