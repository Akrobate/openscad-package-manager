package cmd

import (
	"fmt"

	"github.com/Akrobate/openscad-package-manager/pkg/repository"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [package]",
	Short: "Search for OpenSCAD packages",
	Long:  `Searches for OpenSCAD packages.`,
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		searchString := ""
		if len(args) > 0 {
			searchString = args[0]
		}

		repositoryMgr, err := repository.NewRepositoryManager()
		if err != nil {
			return fmt.Errorf("failed to initialize manager: %w", err)
		}

		foundPackageList, err := repositoryMgr.Search(searchString)
		if err != nil {
			return fmt.Errorf("Failed to search package: %w", err)
		}

		for _, foundPackage := range foundPackageList {
			fmt.Printf("%s\t%s\n", foundPackage.Name, foundPackage.Repository)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
