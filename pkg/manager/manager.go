package manager

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"

	"github.com/Akrobate/openscad-package-manager/internal/utils"
	"gopkg.in/yaml.v3"
)

type Manager struct {
	tmpDir             string
	localModulesFolder string
	packageFile        string
}

type Package struct {
	Name         string            `json:"name" yaml:"name"`
	Version      string            `json:"version" yaml:"version"`
	Description  string            `json:"description" yaml:"description"`
	Repository   string            `json:"repository" yaml:"repository"`
	Dependencies map[string]string `json:"dependencies" yaml:"dependencies"`
	Author       string            `json:"author" yaml:"author"`
	Commit       string
}

type GitRef struct {
	Name string
	Ref  string // commit / tag / branch
	URL  string
}

func NewManager() (*Manager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	var scad_modules_foldername = "openscad_modules"
	var packageFile = "scad.json"

	tmpDir := filepath.Join(homeDir, ".opm", "tmp")

	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}

	return &Manager{
		tmpDir:             tmpDir,
		localModulesFolder: filepath.Join(scad_modules_foldername),
		packageFile:        packageFile,
	}, nil
}

/**
 * Install Curent
 */
func (m *Manager) InstallCurrent() error {
	fmt.Println("Reading current " + m.packageFile)

	dir, err := os.Getwd()
	pkg, err := m.loadPackageMetadata(dir)

	if err != nil {
		fmt.Println(m.packageFile + " not found")
		return nil
	}

	os.RemoveAll(m.localModulesFolder)
	err = os.Mkdir(m.localModulesFolder, 0755)
	if err != nil {
		fmt.Println("Cannot create temporary folder")
		return nil
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

	ref, err := parseGitURL(packageSpec)
	if err != nil {
		fmt.Println("Cannot parse url of dependency: " + ref.Name)
	}
	fmt.Println("Installing: " + ref.Name + " url: " + packageSpec)

	var finalFolderName = ref.Name

	os.RemoveAll(filepath.Join(m.tmpDir, ref.Name))
	m.downloadPackage(ref.URL, ref.Ref, filepath.Join(m.tmpDir, ref.Name))
	pkg, err := m.loadPackageMetadata(filepath.Join(m.tmpDir, ref.Name))

	if isSubDependecy {
		finalFolderName = ref.Name + "#" + pkg.Commit
	}

	_, err = os.Stat(filepath.Join(m.localModulesFolder, finalFolderName))
	if err == nil {
		fmt.Println(ref.Name + " Already installed")
		return finalFolderName, nil
	}

	err = os.Rename(filepath.Join(m.tmpDir, ref.Name), filepath.Join(m.localModulesFolder, finalFolderName))
	if err != nil {
		fmt.Println("Cannot move file from: " + filepath.Join(m.tmpDir, ref.Name+" to: "+filepath.Join(m.localModulesFolder, finalFolderName)))
	}

	err = os.RemoveAll(filepath.Join(m.tmpDir, ref.Name))
	if err != nil {
		fmt.Println("Erreur :", err)
	}

	for _, repository_url := range pkg.Dependencies {

		package_name, err := m.Install(repository_url, true)

		if err != nil {
			fmt.Println("Install fail " + repository_url)
		}

		dependecyRef, err := parseGitURL(repository_url)

		if err != nil {
			fmt.Println("parseGitURL " + repository_url)
		}

		utils.OpenscadReplaceDependienciesPathes(
			filepath.Join(m.localModulesFolder, finalFolderName),
			"openscad_modules/"+dependecyRef.Name,
			"../"+package_name,
		)
	}

	return finalFolderName, nil
}

/**
 * Uninstall
 */
func (m *Manager) Uninstall(packageName string) error {
	return fmt.Errorf("Unitary uninstal, not implemented, use \"opm uninstall\" instead")
}

/**
 * Uninstall
 */
func (m *Manager) UninstallAll() error {
	os.RemoveAll(m.localModulesFolder)
	os.Mkdir(m.localModulesFolder, 0755)
	return nil
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
 * Search
 */
func (m *Manager) Search(query string) ([]Package, error) {
	var results []Package
	return results, nil
}

/**
 * fetchPackageInfo
 */
func (m *Manager) fetchPackageInfo(name, version string) (*Package, error) {
	fmt.Println(name, version)
	return nil, nil
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
	data, err := yaml.Marshal(pkg)
	if err != nil {
		return fmt.Errorf("failed to marshal package metadata: %w", err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write metadata file: %w", err)
	}

	return nil
}

func (m *Manager) loadPackageMetadata(filePath string) (*Package, error) {

	data, err := os.ReadFile(filepath.Join(filePath, m.packageFile))
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

func parseGitURL(raw string) (*GitRef, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	base := path.Base(u.Path)
	repo := strings.TrimSuffix(base, ".git")
	ref := strings.TrimSpace(u.Fragment)

	u.Fragment = ""
	urlWithoutFragment := u.String()

	return &GitRef{
		Name: repo,
		Ref:  ref,
		URL:  urlWithoutFragment,
	}, nil
}
