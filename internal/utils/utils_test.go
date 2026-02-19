package utils

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func TestGetNameFromDependencySpecString(t *testing.T) {
	result, _ := GetNameFromDependencySpecString("https://gitlab.com/openscad-modules/breadboard.git")
	if result != "breadboard" {
		t.Errorf(
			"GetNameFromDependencySpecString(\"https://gitlab.com/openscad-modules/breadboard.git\") = %s, attendu breadboard",
			result,
		)
	}
}

func TestGetRefFromDependencySpecString(t *testing.T) {
	result, _ := GetRefFromDependencySpecString("https://gitlab.com/openscad-modules/breadboard.git#Test")
	if result != "Test" {
		t.Errorf(
			"GetRefFromDependencySpecString(\"https://gitlab.com/openscad-modules/breadboard.git#Test\") = %s, attendu Test",
			result,
		)
	}
}

func TestGetURLFromDependencySpecString(t *testing.T) {
	result, _ := GetURLFromDependencySpecString("https://gitlab.com/openscad-modules/breadboard.git#Test")
	if result != "https://gitlab.com/openscad-modules/breadboard.git" {
		t.Errorf(
			"GetRefFromDependencySpecString(\"https://gitlab.com/openscad-modules/breadboard.git#Test\") = %s, attendu https://gitlab.com/openscad-modules/breadboard",
			result,
		)
	}
}

func TestOpenscadReplaceDependenciesPathes_TempDir(t *testing.T) {
	dir := t.TempDir() // crée un dossier temporaire

	scadFile := filepath.Join(dir, "a.scad")
	os.WriteFile(scadFile, []byte("<from/path/file>"), 0644)

	txtFile := filepath.Join(dir, "b.txt")
	os.WriteFile(txtFile, []byte("ne pas toucher"), 0644)

	OpenscadReplaceDependienciesPathes(dir, "from/path", "to/path")

	data, _ := os.ReadFile(scadFile)
	if string(data) != "<to/path/file>" {
		t.Errorf("contenu incorrect: %s", data)
	}

	data, _ = os.ReadFile(txtFile)
	if string(data) != "ne pas toucher" {
		t.Errorf("fichier non .scad modifié")
	}
}

func TestGetGitHeadShortCommit(t *testing.T) {
	dir := t.TempDir()

	// 1. Creating empty git
	repo, err := git.PlainInit(dir, false)
	if err != nil {
		t.Fatal(err)
	}

	// 2. Create file for commit
	file := filepath.Join(dir, "test.txt")
	os.WriteFile(file, []byte("hello"), 0644)

	// 3. Add commit
	w, _ := repo.Worktree()
	w.Add("test.txt")
	commit, err := w.Commit("initial commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Test",
			Email: "test@example.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	hash, err := GetGitHeadShortCommit(dir)
	if err != nil {
		t.Fatal(err)
	}

	expected := commit.String()[:7]
	if hash != expected {
		t.Errorf("hash incorrect: got %s, want %s", hash, expected)
	}
}

func TestDirExists(t *testing.T) {
	tmpDir := t.TempDir()
	fakeDirFilePath := filepath.Join(tmpDir, "FakeDir")

	result := DirExists(fakeDirFilePath)
	if result == true {
		t.Errorf("DirExists fake dir should return false")
	}

	existingDirFilePath := filepath.Join(tmpDir, "RealDir")
	os.Mkdir(existingDirFilePath, 0755)

	result = DirExists(existingDirFilePath)
	if result == false {
		t.Errorf("DirExists real dir should return true")
	}
}

func TestFileExists(t *testing.T) {
	dir := t.TempDir() // crée un dossier temporaire

	txtFile := filepath.Join(dir, "b.txt")
	os.WriteFile(txtFile, []byte("ne pas toucher"), 0644)

	if FileExists(txtFile) == false {
		t.Errorf("File created FileExists should return true")
	}

	notExistingTxtFile := filepath.Join(dir, "c.txt")
	if FileExists(notExistingTxtFile) == true {
		t.Errorf("File not existing FileExists should return false")
	}
}

func TestURLToFilenameHash(t *testing.T) {
	hashedString := URLToFilenameHash("StringToHash")
	if hashedString != "a1af46c3980254cd987b671e0d3c79a656fedb1451ae95cf832561f6e58479bf" {
		t.Errorf("URLToFilenameHash hash is not correct")
	}
}

// Fonction utilitaire pour comparer des slices de strings
func equal(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestExtractTagValues(t *testing.T) {
	code := `
	/**
	 * UsbChargerFacadeHolderPiece
	 * @name UsbChargerFacadeHolderPiece
	 * @type piece
	 * @parent UsbChargerComponent
	 */

	/**
	 * AnotherBlock
	 * @type    component
	 */
	`

	// Test @name
	result_name := extractTagValues(code, "name")
	expected_name := []string{"UsbChargerFacadeHolderPiece"}

	if !equal(result_name, expected_name) {
		t.Errorf("extractTagValues(code, \"name\") = %v; want %v", result_name, expected_name)
	}

	// Test @type
	result_type := extractTagValues(code, "type")
	expected_type := []string{"piece", "component"}

	if !equal(result_type, expected_type) {
		t.Errorf("extractTagValues(code, \"type\") = %v; want %v", result_type, expected_type)
	}

	// Test @parent
	result_parent := extractTagValues(code, "parent")
	expected_parent := []string{"UsbChargerComponent"}

	if !equal(result_parent, expected_parent) {
		t.Errorf("extractTagValues(code, \"parent\") = %v; want %v", result_parent, expected_parent)
	}
}

func TestExtractAnnotations(t *testing.T) {
	code := `
	/**
	 * UsbChargerFacadeHolderPiece
	 * @name UsbChargerFacadeHolderPiece
	 * @type piece
	 * @parent UsbChargerComponent
	 * @stl
	 * @render
	 */

	/**
	 * AnotherBlock
	 * @type    component
	 */
	`

	expected := []map[string]string{
		{
			"name":   "UsbChargerFacadeHolderPiece",
			"type":   "piece",
			"parent": "UsbChargerComponent",
			"stl":    "",
			"render": "",
		},
		{
			"type": "component",
		},
	}

	result := extractAnnotations(code)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, result)
	}
}

func TestContainsKeyValue(t *testing.T) {
	seed := []map[string]string{
		{
			"name":   "UsbChargerFacadeHolderPiece",
			"type":   "piece",
			"parent": "UsbChargerComponent",
			"stl":    "",
			"render": "",
		},
		{
			"type": "component",
		},
	}

	if !containsKeyValue(seed, "type", "piece") {
		t.Errorf("Expected:containsKeyValue(seed, type, piece) to be true")
	}
}
