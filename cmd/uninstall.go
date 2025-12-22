package cmd

import (
	"fmt"

	"github.com/openscad-package-manager/opm/pkg/manager"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall [package]",
	Short: "Désinstaller un package OpenSCAD",
	Long: `Désinstalle un package OpenSCAD installé.

Exemples:
  opm uninstall BOSL2`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		packageName := args[0]
		
		mgr, err := manager.NewManager()
		if err != nil {
			return fmt.Errorf("failed to initialize manager: %w", err)
		}

		fmt.Printf("Désinstallation de %s...\n", packageName)
		
		if err := mgr.Uninstall(packageName); err != nil {
			return fmt.Errorf("failed to uninstall package: %w", err)
		}

		fmt.Printf("✓ Package %s désinstallé avec succès\n", packageName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}

