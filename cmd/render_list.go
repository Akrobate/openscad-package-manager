package cmd

import (
	"github.com/Akrobate/openscad-package-manager/pkg/renderer"
	"github.com/spf13/cobra"
	"fmt"
)

var renderListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show render list files",
	Args:  cobra.ExactArgs(1),
	ValidArgs: []string{"stl", "png"},
	RunE: func(cmd *cobra.Command, args []string) error {
		outputType := args[0]
		if outputType != "stl" && outputType != "png" {
			return fmt.Errorf("invalid argument %q: must be \"stl\" or \"png\"", outputType)
		}

		renderer, err := renderer.NewRenderer()
		if err != nil {
			return fmt.Errorf("failed to initialize renderer: %w", err)
		}
		renderer.List(outputType)
		return nil
	},
}

func init() {
	renderCmd.AddCommand(renderListCmd)
}
