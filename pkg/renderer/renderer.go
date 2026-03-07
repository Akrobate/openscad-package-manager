package renderer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Akrobate/openscad-package-manager/internal/utils"
)

type Renderer struct {
	pngFilesFolderName string
	stlFilesFolderName string
	openscadBinFile    string
}

func NewRenderer() (*Renderer, error) {

	return &Renderer{
		pngFilesFolderName: "opm_png_files",
		stlFilesFolderName: "opm_stl_files",
		openscadBinFile:    "openscad",
	}, nil
}

func (r *Renderer) List(renderType string) error {
	if err := checkRenderTypeParam(renderType); err != nil {
		return err
	}

	if !r.checkOpenscadInstalled() {
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

	if !r.checkOpenscadInstalled() {
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

			var generationError error

			if renderType == "png" {
				generationError = r.generatePngFile(file, anotationsList)
			} else {
				generationError = r.generateStlFile(file)
			}

			if generationError != nil {
				fmt.Printf("❌ %s\n", file)
				fmt.Println(generationError)
				continue
			} else {
				fmt.Printf("✅ %s\n", file)
			}
		}
	}

	return nil
}

func (r *Renderer) generateStlFile(file string) error {
	args, err := r.generateOpenscadStlCommandLineParams(file)
	if err != nil {
		return fmt.Errorf("generateOpenscadStlCommandLineParams %s", err)
	}
	return r.runOpenscadCommand(args)
}

func (r *Renderer) generatePngFile(file string, anotationsList []map[string]string) error {
	args, err := r.generateOpenscadPngCommandLineParams(file, anotationsList)
	if err != nil {
		return fmt.Errorf("generateOpenscadPngCommandLineParams %s", err)
	}
	return r.runOpenscadCommand(args)
}

func (r *Renderer) generateOpenscadPngCommandLineParams(file string, anotationsList []map[string]string) ([]string, error) {

	dir := filepath.Dir(file)
	pngFileFolder := filepath.Join(r.pngFilesFolderName, dir)
	err := os.MkdirAll(pngFileFolder, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("Error %s", err)
	}

	commandLineArgs := []string{}

	if utils.AnnotationsContainsKey(anotationsList, "colorscheme") {
		commandLineArgs = append(commandLineArgs, fmt.Sprintf("--colorscheme=%s", utils.AnnotationsGetValue(anotationsList, "colorscheme")))
	}

	if utils.AnnotationsContainsKey(anotationsList, "view") {
		commandLineArgs = append(commandLineArgs, fmt.Sprintf("--view=%s", utils.AnnotationsGetValue(anotationsList, "view")))
	}
	filename := filepath.Base(file)
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	commandLineArgs = append(commandLineArgs, "-o", filepath.Join(pngFileFolder, name+".png"), file)

	return commandLineArgs, nil
}

func (r *Renderer) generateOpenscadStlCommandLineParams(file string) ([]string, error) {

	stlFileFolder := filepath.Join(r.stlFilesFolderName)
	err := os.MkdirAll(stlFileFolder, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("Error %s", err)
	}

	paramsStlGeneration := []string{}
	filename := filepath.Base(file)
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	paramsStlGeneration = append(paramsStlGeneration, "-o", filepath.Join(stlFileFolder, name+".stl"), file)

	return paramsStlGeneration, nil
}

// runOpenscadCommand
func (r *Renderer) runOpenscadCommand(args []string) error {
	cmd := exec.Command(r.openscadBinFile, args...)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error runOpenscadCommand %s", err)
	}
	return nil
}

// checkOpenscadInstalled
func (r *Renderer) checkOpenscadInstalled() bool {
	_, err := exec.LookPath(r.openscadBinFile)
	return err == nil
}

// checkRenderTypeParam
func checkRenderTypeParam(renderType string) error {
	if renderType != "stl" && renderType != "png" {
		return fmt.Errorf("invalid render type: %s", renderType)
	}
	return nil
}
