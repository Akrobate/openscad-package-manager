package cmd

import (
	"fmt"

	"github.com/Akrobate/openscad-package-manager/pkg/manager"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info [package]",
	Short: "Info about an Openscad module",
	Long: `Info about an Openscad module.
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		url := args[0]

		mgr, err := manager.NewManager()
		if err != nil {
			return fmt.Errorf("failed to initialize manager: %w", err)
		}

		if err = mgr.Info(url); err != nil {
			return fmt.Errorf("Failed to search package: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
