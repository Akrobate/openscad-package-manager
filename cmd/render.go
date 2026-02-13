package cmd

import (
	"github.com/spf13/cobra"
)

var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Render management",
}

func init() {
	rootCmd.AddCommand(renderCmd)
}
