package cmd

import (
	"fmt"

	"github.com/Akrobate/openscad-package-manager/pkg/manager"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lister les packages installés",
	Long:  `Affiche la liste de tous les packages OpenSCAD installés localement.`,
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
			fmt.Println("Aucun package installé.")
			return nil
		}

		fmt.Println("Packages installés:")
		for _, pkg := range packages {
			fmt.Printf("  - %s#%s - %s\n", pkg.Name, pkg.Version, pkg.Commit)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
