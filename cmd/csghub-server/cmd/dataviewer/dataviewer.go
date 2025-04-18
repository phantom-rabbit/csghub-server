package dataviewer

import (
	"github.com/spf13/cobra"
)

func init() {
	// add subcommands here
	Cmd.AddCommand(launchCmd)
}

var Cmd = &cobra.Command{
	Use:   "dataviewer",
	Short: "entry point for data viewer",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
