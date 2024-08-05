package main

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/polyrepopro/cli/workspace"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "poly",
	Short: "Polyrepo CLI tool",
	Long:  "Polyrepo CLI tool",
}

func main() {
	root.AddCommand(workspace.WorkspaceCommand)
	root.PersistentFlags().StringP("config", "c", "", "The path to the polyrepo config file.")
	root.PersistentFlags().BoolP("verbose", "v", false, "Output detailed logs.")

	logLevel := multilog.INFO

	if verbose, _ := root.PersistentFlags().GetBool("verbose"); verbose {
		logLevel = multilog.DEBUG
	}

	multilog.RegisterLogger(multilog.LogMethod("console"), multilog.NewConsoleLogger(&multilog.NewConsoleLoggerArgs{
		Level:  logLevel,
		Format: multilog.FormatText,
	}))

	root.Execute()

}
