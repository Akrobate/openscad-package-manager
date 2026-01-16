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
		repository.List()
		fmt.Println("========================")
		return nil
	},
}

func init() {
	repositoryCmd.AddCommand(repositoryListCmd)
}
