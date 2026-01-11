package cmd

import (
	"fmt"

	"github.com/Akrobate/openscad-package-manager/pkg/manager"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List project's installed packages",
	Long:  `Show all project's installed packages`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mgr, err := manager.NewManager()
		if err != nil {
			return fmt.Errorf("failed to initialize manager: %w", err)
		}

		packages, err := mgr.List()
		if err != nil {
			return fmt.Errorf("failed to list packages: %w", err)
		}

		if len(packages) == 0 {
			fmt.Println("No packages installed")
			return nil
		}

		fmt.Println("Installed packages:")
		for _, pkg := range packages {
			fmt.Printf("  - %s#%s - %s\n", pkg.Name, pkg.Version, pkg.Commit)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
