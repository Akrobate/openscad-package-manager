package cmd

import (
	"fmt"

	"github.com/Akrobate/openscad-package-manager/pkg/manager"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Rechercher des packages dans le registre",
	Long: `Recherche des packages OpenSCAD disponibles dans le registre.

Exemples:
  opm search BOSL
  opm search utility`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := args[0]

		mgr, err := manager.NewManager()
		if err != nil {
			return fmt.Errorf("failed to initialize manager: %w", err)
		}

		fmt.Printf("Recherche de '%s'...\n", query)

		results, err := mgr.Search(query)
		if err != nil {
			return fmt.Errorf("failed to search packages: %w", err)
		}

		if len(results) == 0 {
			fmt.Printf("Aucun package trouvé pour '%s'\n", query)
			return nil
		}

		fmt.Printf("\nRésultats (%d):\n", len(results))
		for _, pkg := range results {
			fmt.Printf("  - %s@%s: %s\n", pkg.Name, pkg.Version, pkg.Description)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
