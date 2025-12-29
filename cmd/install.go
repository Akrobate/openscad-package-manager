package cmd

import (
	"fmt"

	"github.com/Akrobate/openscad-package-manager/pkg/manager"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [package]",
	Short: "Install Openscad module",
	Long: `Installs Openscad module and its dependecies.

Exemples:
  opm install
  opm install
  opm install`,
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
			return mgr.InstallCurrent()
		}

		fmt.Printf("Installation de %s...\n", packageName)

		if _, err := mgr.Install(packageName, false); err != nil {
			return fmt.Errorf("failed to install package: %w", err)
		}

		fmt.Printf("✓ Package %s installé avec succès\n", packageName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
