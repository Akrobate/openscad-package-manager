package cmd

import (
	"github.com/spf13/cobra"
)

var repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Repositories management",
}

func init() {
	rootCmd.AddCommand(repositoryCmd)
}
