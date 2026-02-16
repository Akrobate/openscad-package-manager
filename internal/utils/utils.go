package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"slices"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
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

func GetGitTags(repository_path string) ([]string, error) {
	repo, err := git.PlainOpen(repository_path)
	if err != nil {
		return nil, err
	}

	var result []string

	tags, err := repo.Tags()
	if err != nil {
		return nil, err
	}

	tags.ForEach(func(ref *plumbing.Reference) error {
		result = append(result, ref.Name().Short())
		return nil
	})

	return result, err
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

func DirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return info.IsDir()
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

func URLToFilenameHash(u string) string {
	sum := sha256.Sum256([]byte(u))
	return hex.EncodeToString(sum[:])
}

func TempDir() (string, func(), error) {
	dir, err := os.MkdirTemp("", "tmp_opm_*")
	if err != nil {
		return "", nil, err
	}
	cleanup := func() { os.RemoveAll(dir) }
	return dir, cleanup, nil
}

// FindFilesWithSpecificType find all files with a specific type
func FindFilesWithSpecificType(rootDir string, element_type string) ([]string, error) {

	var results []string

	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && d.Name() == "openscad_modules" {
			return filepath.SkipDir
		}

		if d.IsDir() || filepath.Ext(path) != ".scad" {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		matches := extractTagValues(string(data), "type")
		if slices.Contains(matches, element_type) {
			results = append(results, path)
		}
		return nil
	})

	return results, err
}

// extractTagValues extract all values of a specific tag
func extractTagValues(code string, tag string) []string {
	var results []string

	escapedTag := regexp.QuoteMeta(tag)

	pattern := fmt.Sprintf(`@%s[ \t]+([^\s*]+)[ \t]*`, escapedTag)
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(code, -1)

	for _, m := range matches {
		if len(m) > 1 {
			results = append(results, m[1])
		}
	}

	return results
}
