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
		repositoryMgr, _ := repository.NewRepositoryManager()
		repositoryMgr.List()
		fmt.Println("======== List List================")
		return nil
	},
}

func init() {
	repositoryCmd.AddCommand(repositoryListCmd)
}
