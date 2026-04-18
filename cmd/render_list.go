package cmd

import (
	"fmt"

	"github.com/Akrobate/openscad-package-manager/pkg/renderer"
	"github.com/spf13/cobra"
)

var renderListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show render png / stl files",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("you must provide exactly one argument: \"stl\" or \"png\"")
		}

		if args[0] != "stl" && args[0] != "png" {
			return fmt.Errorf("invalid argument %q: must be \"stl\" or \"png\"", args[0])
		}

		return nil
	},
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
