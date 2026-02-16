package renderer

import (
	"fmt"
	"github.com/Akrobate/openscad-package-manager/internal/utils"
)

type Renderer struct{}

func NewRenderer() (*Renderer, error) {

	return &Renderer{}, nil
}

func (r *Renderer) List() error {

	files, err := utils.FindFilesWithSpecificType(".", "piece")
	if err != nil {
		return fmt.Errorf("failed to find files with specific type: %w", err)
	}
	fmt.Println(files)
	return nil

}
