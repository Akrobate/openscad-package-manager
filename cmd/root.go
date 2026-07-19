package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "opm",
	Short: "OpenSCAD Package Manager",
	Long: `OpenSCAD Package Manager (opm)

install, list, uninstall and init packages for OpenSCAD projects
`,
	Version: "0.4.0",
}

func Execute() error {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	return rootCmd.Execute()
}
