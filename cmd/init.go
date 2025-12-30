package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Akrobate/openscad-package-manager/pkg/manager"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes scad.json file",
	Long: `Initializes scad.json file.

Exemples:
  opm init`,
	Args: cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {

		mgr, err := manager.NewManager()
		if err != nil {
			return fmt.Errorf("failed to initialize manager: %w", err)
		}

		dir, err := os.Getwd()
		base := filepath.Base(dir)

		var pkg manager.Package
		pkg.Name = askUser("package name", base)
		pkg.Version = askUser("version", "1.0.0")
		pkg.Description = askUser("description", "")
		pkg.Repository = askUser("repository", "")
		pkg.Author = askUser("author", "")

		if err := mgr.Init(&pkg); err != nil {
			return fmt.Errorf("failed to install package: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func askUser(promt string, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(promt + ": ")
	if defaultValue != "" {
		fmt.Printf("(" + defaultValue + ") ")
	}
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		input = defaultValue
	}
	return input
}
