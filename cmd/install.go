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

Install without params will install dependencies from the current scad.json file	
Install with a repository will install this package latest version


Exemples:
  opm install
  opm install https://gitlab.com/openscad-modules/housing.git
  opm install https://gitlab.com/openscad-modules/housing.git#0.0.2
  opm install https://gitlab.com/openscad-modules/housing.git#develop
  opm install https://gitlab.com/openscad-modules/housing.git#5ebc661`,
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

		fmt.Printf("Installing  %s...\n", packageName)

		if _, err := mgr.Install(packageName, false); err != nil {
			return fmt.Errorf("failed to install package: %w", err)
		}

		fmt.Printf("✓ Package install success\n")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
