package cmd

import (
	"fmt"

	"github.com/Akrobate/openscad-package-manager/pkg/repository"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [package]",
	Short: "Search Openscad module",
	Long: `Search Openscad modules.
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		searchString := args[0]

		repositoryMgr, err := repository.NewRepositoryManager()
		if err != nil {
			return fmt.Errorf("failed to initialize manager: %w", err)
		}

		foundPackageList, err := repositoryMgr.Search(searchString)
		if err != nil {
			return fmt.Errorf("Failed to search package: %w", err)
		}

		fmt.Println(foundPackageList)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
