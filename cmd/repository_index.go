package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var repositoryIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "Show repository list index",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("========================")
		return nil
	},
}

func init() {
	repositoryCmd.AddCommand(repositoryIndexCmd)
}
