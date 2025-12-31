package utils

import (
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-git/go-git/v5"
)

func OpenscadReplaceDependienciesPathes(rootDir string, from string, to string) {

	re := regexp.MustCompile(`<([^>]+)>`)

	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || filepath.Ext(path) != ".scad" {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		modified := re.ReplaceAllStringFunc(string(data), func(s string) string {
			content := s[1 : len(s)-1]
			content = regexp.MustCompile(from).ReplaceAllString(content, to)
			return "<" + content + ">"
		})

		if err := os.WriteFile(path, []byte(modified), 0755); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println("Erreur :", err)
	}
}

func GetGitHeadShortCommit(repository_path string) (string, error) {
	repo, err := git.PlainOpen(repository_path)
	if err != nil {
		return "", err
	}
	ref, err := repo.Head()
	if err != nil {
		return "", err
	}
	hash := ref.Hash().String()
	return hash[:7], nil
}

func GetNameFromDependencySpecString(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	base := path.Base(u.Path)
	name := strings.TrimSuffix(base, ".git")
	return name, nil
}

func GetRefFromDependencySpecString(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	ref := strings.TrimSpace(u.Fragment)
	return ref, nil
}

func GetURLFromDependencySpecString(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	u.Fragment = ""
	urlWithoutFragment := u.String()
	return urlWithoutFragment, nil
}
