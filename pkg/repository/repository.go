package repository

import (
	"fmt"
	"os"
)

/**
 * Install Curent
 */
func List() error {

	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf(" not found")
	}
	fmt.Print(dir)

	return nil
}
