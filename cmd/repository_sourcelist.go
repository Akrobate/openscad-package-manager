package cmd

import (
	"fmt"

	"github.com/Akrobate/openscad-package-manager/pkg/repository"
	"github.com/spf13/cobra"
)

var repositorySourceListCmd = &cobra.Command{
	Use:   "sourcelist",
	Short: "Show the repository lists",
	Long:  `Shows the configured repository lists.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repositoryMgr, err := repository.NewRepositoryManager()
		if err != nil {
			return fmt.Errorf("failed to initialize repository manager: %w", err)
		}
		sourceList, err := repositoryMgr.ListSources()
		if err != nil {
			return fmt.Errorf("failed to repositoryMgr.List: %w", err)
		}
		for _, sourceListItem := range sourceList {
			fmt.Printf("%s\n", sourceListItem)
		}
		return nil
	},
}

func init() {
	repositoryCmd.AddCommand(repositorySourceListCmd)
}
