package cmd

import (
	"fmt"

	"github.com/Akrobate/openscad-package-manager/pkg/renderer"
	"github.com/spf13/cobra"
)

var renderProcessCmd = &cobra.Command{
	Use:   "process",
	Short: "Process png / stl / md renders",
	Long: `Processes png, stl or md renders.

Examples:
  opm render process stl
  opm render process png
  opm render process md`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("you must provide exactly one argument: \"stl\", \"png\" or \"md\"")
		}

		if args[0] != "stl" && args[0] != "png" && args[0] != "md" {
			return fmt.Errorf("invalid argument %q: must be \"stl\", \"png\" or \"md\"", args[0])
		}

		return nil
	},
	ValidArgs: []string{"stl", "png", "md"},
	RunE: func(cmd *cobra.Command, args []string) error {
		outputType := args[0]
		if outputType != "stl" && outputType != "png" && outputType != "md" {
			return fmt.Errorf("invalid argument %q: must be \"stl\", \"png\" or \"md\"", outputType)
		}

		renderer, err := renderer.NewRenderer()
		if err != nil {
			return fmt.Errorf("failed to initialize renderer: %w", err)
		}
		renderer.Process(outputType)
		return nil
	},
}

func init() {
	renderCmd.AddCommand(renderProcessCmd)
}
