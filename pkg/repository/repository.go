package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/Akrobate/openscad-package-manager/internal/utils"
)

type PackageItem struct {
	Name       string `json:"repository" yaml:"repository"`
	Repository string `json:"name" yaml:"name"`
}

type RepositoryManager struct {
	repositorySourcesListFile string
	repositorySourcesCache    string
}

func NewRepositoryManager() (*RepositoryManager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	repositorySourcesCache := filepath.Join(homeDir, ".opm", "repository", "cache")
	if err := os.MkdirAll(repositorySourcesCache, 0755); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}

	repositorySourcesListFile := filepath.Join(homeDir, ".opm", "repository", "sources-list.txt")

	return &RepositoryManager{
		repositorySourcesListFile: repositorySourcesListFile,
		repositorySourcesCache:    repositorySourcesCache,
	}, nil
}

/**
 * List Curent
 */
func (m *RepositoryManager) List() error {

	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf(" not found")
	}
	fmt.Print(dir)

	return nil
}

/**
 * Add
 */
func (m *RepositoryManager) Add(repositorySourcesUrl string) error {

	content, err := m.getSourceList(repositorySourcesUrl)
	if err != nil {
		return fmt.Errorf("Failed to reach URL\n %w", err)
	}

	f, err := os.OpenFile(
		m.repositorySourcesListFile,
		os.O_APPEND|os.O_CREATE|os.O_RDWR,
		0644,
	)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Println("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")
	contentBytes, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	contentString := string(contentBytes)
	fmt.Println(contentString)
	repositorySourcesListFileLines := strings.Split(contentString, "\n")

	fmt.Println(repositorySourcesListFileLines)
	fmt.Println("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO")

	for _, v := range repositorySourcesListFileLines {
		if v == repositorySourcesUrl {
			return fmt.Errorf("Repository list already exists")
		}
	}

	_, err = f.WriteString(repositorySourcesUrl + "\n")

	contentLines := strings.Split(content, "\n")
	var modules []PackageItem
	for _, u := range contentLines {
		name := strings.TrimSuffix(path.Base(u), ".git")
		modules = append(modules, PackageItem{
			Name:       name,
			Repository: u,
		})
	}

	jsonBytes, err := json.MarshalIndent(modules, "", "    ")
	if err != nil {
		panic(err)
	}

	cacheFileName := utils.URLToFilenameHash(repositorySourcesUrl)
	if err := os.WriteFile(filepath.Join(m.repositorySourcesCache, cacheFileName), []byte(jsonBytes), 0755); err != nil {
		return err
	}

	return err
}

/**
 * getSourceList
 */
func (m *RepositoryManager) getSourceList(url string) (string, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Failed getSourceList %s", resp.Status)
	}

	return string(body), nil
}
