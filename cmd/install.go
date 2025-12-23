package cmd

import (
	"fmt"

	"github.com/openscad-package-manager/opm/pkg/manager"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [package]",
	Short: "Installer un package OpenSCAD",
	Long: `Installe un package OpenSCAD et ses dépendances.

Exemples:
  opm install BOSL2
  opm install github.com/user/repo
  opm install package@1.0.0`,
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
			mgr.InstallCurrent()
			// comportement par défaut (ex: installer depuis un fichier, tout mettre à jour, etc.)
			return nil
		}

		fmt.Printf("Installation de %s...\n", packageName)

		if err := mgr.Install(packageName); err != nil {
			return fmt.Errorf("failed to install package: %w", err)
		}

		fmt.Printf("✓ Package %s installé avec succès\n", packageName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
