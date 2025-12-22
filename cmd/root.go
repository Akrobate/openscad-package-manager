package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "opm",
	Short: "OpenSCAD Package Manager",
	Long: `OpenSCAD Package Manager (opm) est un gestionnaire de paquets
pour les bibliothèques et modules OpenSCAD.

Il permet d'installer, gérer et utiliser des packages OpenSCAD
depuis un registre centralisé.`,
	Version: "0.1.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.opm/config.yaml)")
	rootCmd.PersistentFlags().String("registry", "https://registry.openscad-packages.org", "URL du registre de packages")
	
	viper.BindPFlag("registry", rootCmd.PersistentFlags().Lookup("registry"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home + "/.opm")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

