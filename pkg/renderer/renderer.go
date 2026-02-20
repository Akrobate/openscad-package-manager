package renderer

import (
	"fmt"
	"os"
	"github.com/Akrobate/openscad-package-manager/internal/utils"
)

type Renderer struct{}

func NewRenderer() (*Renderer, error) {

	return &Renderer{}, nil
}

func (r *Renderer) List(renderType string) error {
	if renderType != "stl" && renderType != "png" {
		return fmt.Errorf("invalid render type: %s", renderType)
	}

	files, err := utils.ListAllProjectScadFiles(".")

	if err != nil {
		return fmt.Errorf("failed to find files with specific type: %w", err)
	}

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		anotationsList := utils.ExtractAnnotations(string(data))

		if utils.AnnotationsContainsKeyValue(anotationsList, renderType, "") {
			fmt.Println(file)
		}

	}

	return nil

}
