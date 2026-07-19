package main

import (
	"os"

	"github.com/Akrobate/openscad-package-manager/cmd"
)

func main() {
	// Cette fonction magique s'adaptera toute seule selon l'OS
	fixWindowsConsole()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
