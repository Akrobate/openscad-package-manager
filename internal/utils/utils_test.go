package utils

import (
	"os"
	"path/filepath"
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
	os.Mkdir(existingDirFilePath, 755)

	result = DirExists(existingDirFilePath)
	if result == false {
		t.Errorf("DirExists real dir should return true")
	}
}
