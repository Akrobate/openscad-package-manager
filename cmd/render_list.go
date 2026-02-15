package cmd

import (
	"github.com/Akrobate/openscad-package-manager/pkg/renderer"
	"github.com/spf13/cobra"
	"fmt"
)

var renderListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show repository list",
	RunE: func(cmd *cobra.Command, args []string) error {
		renderer, err := renderer.NewRenderer()
		if err != nil {
			return fmt.Errorf("failed to initialize renderer: %w", err)
		}
		renderer.List()
		return nil
	},
}

func init() {
	renderCmd.AddCommand(renderListCmd)
}
