package renderer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Akrobate/openscad-package-manager/internal/utils"
)

type Renderer struct{}

func NewRenderer() (*Renderer, error) {

	return &Renderer{}, nil
}

func (r *Renderer) List(renderType string) error {
	if err := checkRenderTypeParam(renderType); err != nil {
		return err
	}

	if !checkOpenscadInstalled() {
		return fmt.Errorf("Openscad bin not foud")
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
			if renderType == "png" {
				fmt.Println(file)
			} else {
				fmt.Println(file)
			}
		}
	}
	return nil
}

func (r *Renderer) Process(renderType string) error {
	if err := checkRenderTypeParam(renderType); err != nil {
		return err
	}

	if !checkOpenscadInstalled() {
		return fmt.Errorf("Openscad bin not foud")
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
			if renderType == "png" {
				fmt.Println(file)
				generatePngFile(file, anotationsList)
			} else {
				fmt.Println(file)
			}
		}
	}

	return nil
}

func generatePngFile(file string, anotationsList []map[string]string) error {

	dir := filepath.Dir(file)
	pngFileFolder := filepath.Join("opm_png_files", dir)
	err := os.MkdirAll(pngFileFolder, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Error %s", err)
	}

	paramsPngGeneration := generateOpenscadPngCommandLine(anotationsList)

	filename := filepath.Base(file)
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	args := append(paramsPngGeneration, "-o ", filepath.Join(pngFileFolder, name+".png", " .", file))

	cmd := exec.Command("openscad", args...)

	fmt.Println(strings.Join(args, ""))

	// exécuter
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Error openscad %s", err)
	}
	return nil
}

func generateOpenscadPngCommandLine(anotationsList []map[string]string) []string {
	commandLineArgs := []string{}

	if utils.AnnotationsContainsKey(anotationsList, "colorscheme") {
		commandLineArgs = append(commandLineArgs, fmt.Sprintf(" --colorscheme=\"%s\" ", utils.AnnotationsGetValue(anotationsList, "colorscheme")))
	}

	if utils.AnnotationsContainsKey(anotationsList, "view") {
		commandLineArgs = append(commandLineArgs, fmt.Sprintf(" --view=\"%s\" ", utils.AnnotationsGetValue(anotationsList, "view")))
	}

	return commandLineArgs
}

// checkOpenscadInstalled
func checkOpenscadInstalled() bool {
	_, err := exec.LookPath("openscad")
	return err == nil
}

// checkRenderTypeParam
func checkRenderTypeParam(renderType string) error {
	if renderType != "stl" && renderType != "png" {
		return fmt.Errorf("invalid render type: %s", renderType)
	}
	return nil
}
