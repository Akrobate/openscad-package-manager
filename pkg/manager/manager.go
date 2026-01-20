package manager

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"

	"github.com/Akrobate/openscad-package-manager/internal/utils"
	"gopkg.in/yaml.v3"
)

type Manager struct {
	tmpDir                  string
	localModulesForlderName string
	localModulesFolder      string
	PackageFile             string
}

type Package struct {
	Name         string            `json:"name" yaml:"name"`
	Version      string            `json:"version" yaml:"version"`
	Description  string            `json:"description" yaml:"description"`
	Author       string            `json:"author" yaml:"author"`
	Repository   string            `json:"repository" yaml:"repository"`
	Dependencies map[string]string `json:"dependencies,omitempty" yaml:"dependencies"`
	Commit       string            `json:"-"`
}

func NewManager() (*Manager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	var localModulesForlderName = "openscad_modules"
	var PackageFile = "scad.json"

	tmpDir := filepath.Join(homeDir, ".opm", "tmp")

	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}

	return &Manager{
		tmpDir:                  tmpDir,
		localModulesForlderName: localModulesForlderName,
		localModulesFolder:      filepath.Join(localModulesForlderName),
		PackageFile:             PackageFile,
	}, nil
}

/**
 * Install Curent
 */
func (m *Manager) InstallCurrent() error {

	dir, err := os.Getwd()
	pkg, err := m.loadPackageMetadata(dir)

	if err != nil {
		return fmt.Errorf(m.PackageFile + " not found")
	}

	for _, repository_url := range pkg.Dependencies {
		m.Install(repository_url, false)
	}

	return nil
}

/**
 * Install
 */
func (m *Manager) Install(packageSpec string, isSubDependecy bool) (string, error) {

	packageName, err := utils.GetNameFromDependencySpecString(packageSpec)
	packageRef, err := utils.GetRefFromDependencySpecString(packageSpec)
	packageURL, err := utils.GetURLFromDependencySpecString(packageSpec)

	if err != nil {
		return "", fmt.Errorf("Cannot parse url of dependency: " + packageName)
	}

	fmt.Println("Installing: " + packageName + " url: " + packageSpec)

	var finalFolderName = packageName

	os.RemoveAll(filepath.Join(m.tmpDir, packageName))
	m.downloadPackage(packageURL, packageRef, filepath.Join(m.tmpDir, packageName))
	pkg, err := m.loadPackageMetadata(filepath.Join(m.tmpDir, packageName))

	if isSubDependecy {
		finalFolderName = packageName + "#" + pkg.Commit
	}

	if utils.DirExists(filepath.Join(m.localModulesFolder, finalFolderName)) {
		fmt.Println(packageName + " Already installed")
		return finalFolderName, nil
	}

	if utils.DirExists(filepath.Join(m.localModulesFolder)) == false {
		if err = os.Mkdir(m.localModulesFolder, 0755); err != nil {
			return "", fmt.Errorf("Cannot create local modules folder")
		}
	}

	if err = os.Rename(filepath.Join(m.tmpDir, packageName), filepath.Join(m.localModulesFolder, finalFolderName)); err != nil {
		return "", fmt.Errorf("Cannot move file from: %s to: %s", filepath.Join(m.tmpDir, packageName), filepath.Join(m.localModulesFolder, finalFolderName))
	}

	err = os.RemoveAll(filepath.Join(m.tmpDir, packageName))
	if err != nil {
		return "", fmt.Errorf("RemoveAll fail %w", err)
	}

	for _, repository_url := range pkg.Dependencies {

		package_name, err := m.Install(repository_url, true)
		if err != nil {
			return "", fmt.Errorf("Install fail %s. Error: %w", repository_url, err)
		}
		dependecyName, err := utils.GetNameFromDependencySpecString(repository_url)
		if err != nil {
			return "", fmt.Errorf("GetNameFromDependencySpecString %s. Error: %w", repository_url, err)
		}
		utils.OpenscadReplaceDependienciesPathes(
			filepath.Join(m.localModulesFolder, finalFolderName),
			m.localModulesForlderName+"/"+dependecyName,
			"../"+package_name,
		)
	}

	return finalFolderName, nil
}

/**
 * Uninstall
 */
func (m *Manager) UninstallAll() error {
	os.RemoveAll(m.localModulesFolder)
	return nil
}

/**
 * Init
 */
func (m *Manager) Init(pkg *Package) error {
	return m.savePackageMetadata(pkg, filepath.Join(".", m.PackageFile))
}

/**
 * List
 */
func (m *Manager) List() ([]Package, error) {
	entries, err := os.ReadDir(m.localModulesFolder)
	if err != nil {
		if os.IsNotExist(err) {
			return []Package{}, nil
		}
		return nil, fmt.Errorf("failed to read install directory: %w", err)
	}

	var packages []Package
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		metadataFile := filepath.Join(m.localModulesFolder, entry.Name())
		pkg, err := m.loadPackageMetadata(metadataFile)
		if err != nil {
			fmt.Println(err)
			continue
		}
		packages = append(packages, *pkg)
	}

	return packages, nil
}

/**
 * downloadPackage
 */
func (m *Manager) downloadPackage(url string, git_ref string, destination_directory string) error {

	repository, err := git.PlainClone(
		destination_directory,
		false,
		&git.CloneOptions{
			URL:          url,
			SingleBranch: false,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	err = repository.Fetch(&git.FetchOptions{
		RefSpecs: []config.RefSpec{"refs/*:refs/*", "HEAD:refs/heads/HEAD"},
	})

	if err != nil {
		log.Println("Fetch", err)
		return err
	}

	h, _ := repository.ResolveRevision(plumbing.Revision(git_ref))
	work_tree, _ := repository.Worktree()
	work_tree.Checkout(&git.CheckoutOptions{
		Hash: *h,
	})

	return nil
}

func (m *Manager) savePackageMetadata(pkg *Package, filePath string) error {
	data, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal package metadata: %w", err)
	}

	if err = os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write metadata file: %w", err)
	}

	return nil
}

func (m *Manager) loadPackageMetadata(filePath string) (*Package, error) {

	data, err := os.ReadFile(filepath.Join(filePath, m.PackageFile))
	if err != nil {
		return nil, fmt.Errorf("failed to read metadata file: %w", err)
	}

	var pkg Package
	if err := yaml.Unmarshal(data, &pkg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal package metadata: %w", err)
	}

	commit, _ := utils.GetGitHeadShortCommit(filePath)
	pkg.Commit = commit

	return &pkg, nil
}

func (m *Manager) UpdateDependencyInPackageFile(newDependency string) (*Package, error) {

	data, err := os.ReadFile(filepath.Join(".", m.PackageFile))
	if err != nil {
		return nil, fmt.Errorf("failed to read metadata file: %w", err)
	}

	var pkg Package
	if err := yaml.Unmarshal(data, &pkg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal package metadata: %w", err)
	}

	packageName, _ := utils.GetNameFromDependencySpecString(newDependency)

	if pkg.Dependencies == nil {
		pkg.Dependencies = make(map[string]string)
	}

	pkg.Dependencies[packageName] = newDependency
	m.savePackageMetadata(&pkg, filepath.Join(".", m.PackageFile))

	return &pkg, nil
}

func (m *Manager) Info(url string) error {

	dir, cleanup, _ := utils.TempDir()

	defer cleanup()

	fmt.Println("Dossier temporaire :", dir)
	fmt.Println("URL for info url :", url)

	return nil
}
