package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "opm",
	Short: "OpenSCAD Package Manager",
	Long: `OpenSCAD Package Manager (opm) est un gestionnaire de paquets
pour les bibliothèques et modules OpenSCAD.

Il permet d'installer, gérer et utiliser des packages OpenSCAD
`,
	Version: "0.1.0",
}

func Execute() error {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	return rootCmd.Execute()
}
