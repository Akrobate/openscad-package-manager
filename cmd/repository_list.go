package cmd

import (
	"fmt"

	"github.com/Akrobate/openscad-package-manager/pkg/repository"
	"github.com/spf13/cobra"
)

var repositoryListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show repository list",
	RunE: func(cmd *cobra.Command, args []string) error {
		repositoryMgr, err := repository.NewRepositoryManager()
		if err != nil {
			return fmt.Errorf("failed to initialize repository manager: %w", err)
		}
		repositoryMgr.List()
		fmt.Println("======== List List================")
		return nil
	},
}

func init() {
	repositoryCmd.AddCommand(repositoryListCmd)
}
