package cmd

import (
	"fmt"

	"github.com/Akrobate/openscad-package-manager/pkg/repository"
	"github.com/spf13/cobra"
)

var repositoryUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update source lists",
	RunE: func(cmd *cobra.Command, args []string) error {
		repositoryMgr, err := repository.NewRepositoryManager()
		if err != nil {
			return fmt.Errorf("failed to initialize repository manager: %w", err)
		}
		err = repositoryMgr.Update()
		if err != nil {
			return fmt.Errorf("Update error: %w", err)
		}
		return nil
	},
}

func init() {
	repositoryCmd.AddCommand(repositoryUpdateCmd)
}
