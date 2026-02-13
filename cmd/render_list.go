package cmd

import (
	"github.com/Akrobate/openscad-package-manager/pkg/renderer"
	"github.com/spf13/cobra"
)

var renderListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show repository list",
	RunE: func(cmd *cobra.Command, args []string) error {
		renderer, _ := renderer.NewRenderer()
		renderer.List()
		return nil
	},
}

func init() {
	renderCmd.AddCommand(renderListCmd)
}
