package cmd

import (
	"fmt"

	"github.com/Akrobate/openscad-package-manager/pkg/repository"
	"github.com/spf13/cobra"
)

var repositoryAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a repository list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		sourceList := args[0]
		repositoryMgr, err := repository.NewRepositoryManager()
		if err != nil {
			return fmt.Errorf("failed to initialize repository manager: %w", err)
		}
		err = repositoryMgr.Add(sourceList)
		if err != nil {
			return fmt.Errorf("Failed to add repository list: %w", err)
		}

		return nil
	},
}

func init() {
	repositoryCmd.AddCommand(repositoryAddCmd)
}
