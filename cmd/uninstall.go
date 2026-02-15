package cmd

import (
	"fmt"

	"github.com/Akrobate/openscad-package-manager/pkg/manager"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall [package]",
	Short: "Uninstall all installed packages",
	Long: `Uninstall all installed packages

Exemples:
  opm uninstall`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var packageName string
		if len(args) > 0 {
			packageName = args[0]
		}

		mgr, err := manager.NewManager()
		if err != nil {
			return fmt.Errorf("failed to initialize manager: %w", err)
		}

		if packageName == "" {
			err = mgr.UninstallAll()
			if err != nil {
				return err
			}
			fmt.Println("✓ Packages uninstall succes")
			return nil
		} else {
			fmt.Println("Uninstall a specific package is not implemented")
			return nil
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
