package manager

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/config"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var scad_modules_foldername = "openscad_modules"


type Package struct {
	Name         string            `json:"name" yaml:"name"`
	Version      string            `json:"version" yaml:"version"`
	Description  string            `json:"description" yaml:"description"`
	Repository   string            `json:"repository" yaml:"repository"`
	Dependencies map[string]string `json:"dependencies" yaml:"dependencies"`
	Author       string            `json:"author" yaml:"author"`
}

type Dependency struct {
	Name       string
	Repository string
	commit     string
}

type Manager struct {
	registryURL string
	installDir  string
	cacheDir    string
}

type GitRef struct {
	Name string
	Ref  string // commit / tag / branch
	URL string
}

func NewManager() (*Manager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	registryURL := viper.GetString("registry")
	if registryURL == "" {
		registryURL = "https://registry.openscad-packages.org"
	}

	installDir := filepath.Join(homeDir, ".opm", "packages")
	cacheDir := filepath.Join(homeDir, ".opm", "cache")

	// Créer les répertoires si nécessaire
	if err := os.MkdirAll(installDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create install directory: %w", err)
	}

	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}

	return &Manager{
		registryURL: registryURL,
		installDir:  installDir,
		cacheDir:    cacheDir,
	}, nil
}

/**
 *
 * Install Curent
 *
 */
func (m *Manager) InstallCurrent() error {
	fmt.Println("Reading current scad.jsons")

	pkg, err := m.loadPackageMetadata(filepath.Join("scad.json"))

	if err != nil {
		fmt.Println("scad.json not found")
		return nil
	}

	os.RemoveAll(scad_modules_foldername)

	err = os.Mkdir(scad_modules_foldername, 0755)
	if err != nil {
		fmt.Println("Cannot create temporary folder")
		return nil
	}

	// var dependencies []Dependency
	// fmt.Println(pkg.Dependencies)

	// Installer les dépendances d'abord
	for name, repository_url := range pkg.Dependencies {
		fmt.Println("Installing: " + name + " url: " + repository_url)
		
		ref, err := parseGitURL(repository_url)

		if err != nil {
			fmt.Println("Cannot parse url of dependency: " + name)
		}

		var repo *git.Repository
		repo, err = git.PlainClone(
			filepath.Join(scad_modules_foldername, ref.Name),
			false,
			&git.CloneOptions{
				URL: ref.URL,
				SingleBranch:  false,
			},
		)

		if err != nil {
			log.Fatal(err)
		}

		err = repo.Fetch(&git.FetchOptions{
			RefSpecs: []config.RefSpec{"refs/*:refs/*", "HEAD:refs/heads/HEAD"},
		})

		if err != nil {
			log.Println("Fetch", err)
			return nil
		}
				
		h, _ := repo.ResolveRevision(plumbing.Revision(ref.Ref))

		w, _ := repo.Worktree()
		w.Checkout(&git.CheckoutOptions{
			Hash: *h,
		})

	}

	// dependencies = append(dependencies, Dependency{
	// 	Name:       "Coucou",
	// 	Repository: "Coucou",
	// 	commit:     "Coucou",
	// })

	// fmt.Println(pkg)
	// fmt.Println(dependencies)

	return nil
}



/**
 * Install
 */
func (m *Manager) Install(packageSpec string) error {

	// Parser le nom du package et la version (format: package@version)
	name, version := parsePackageSpec(packageSpec)

	// Récupérer les informations du package depuis le registre
	pkg, err := m.fetchPackageInfo(name, version)
	if err != nil {
		return fmt.Errorf("failed to fetch package info: %w", err)
	}

	// Installer les dépendances d'abord
	// for _, dep := range pkg.Dependencies {
	// 	if err := m.Install(dep); err != nil {
	// 		return fmt.Errorf("failed to install dependency %s: %w", dep, err)
	// 	}
	// }

	// Télécharger et installer le package
	packageDir := filepath.Join(m.installDir, pkg.Name)
	if err := os.MkdirAll(packageDir, 0755); err != nil {
		return fmt.Errorf("failed to create package directory: %w", err)
	}

	// Pour l'instant, on simule l'installation
	// Dans une implémentation complète, on téléchargerait depuis le repository
	if err := m.downloadPackage(pkg, packageDir); err != nil {
		return fmt.Errorf("failed to download package: %w", err)
	}

	// Enregistrer les métadonnées du package
	metadataFile := filepath.Join(packageDir, "package.yaml")
	if err := m.savePackageMetadata(pkg, metadataFile); err != nil {
		return fmt.Errorf("failed to save package metadata: %w", err)
	}

	return nil
}

func (m *Manager) Uninstall(packageName string) error {
	fmt.Println("uninstalling " + packageName)
	return nil
}

func (m *Manager) List() ([]Package, error) {
	entries, err := os.ReadDir(m.installDir)
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

		metadataFile := filepath.Join(m.installDir, entry.Name(), "package.yaml")
		pkg, err := m.loadPackageMetadata(metadataFile)
		if err != nil {
			// Ignorer les packages sans métadonnées valides
			continue
		}

		packages = append(packages, *pkg)
	}

	return packages, nil
}

func (m *Manager) Search(query string) ([]Package, error) {
	// Pour l'instant, on simule une recherche
	// Dans une implémentation complète, on interrogerait le registre
	url := fmt.Sprintf("%s/api/search?q=%s", m.registryURL, query)

	resp, err := http.Get(url)
	if err != nil {
		// Si le registre n'est pas disponible, retourner une liste vide
		return []Package{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []Package{}, nil
	}

	var results []Package
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return []Package{}, nil
	}

	return results, nil
}

func (m *Manager) fetchPackageInfo(name, version string) (*Package, error) {
	// Pour l'instant, on simule la récupération
	// Dans une implémentation complète, on interrogerait le registre
	url := fmt.Sprintf("%s/api/package/%s", m.registryURL, name)
	if version != "" {
		url += "?version=" + version
	}

	resp, err := http.Get(url)
	if err != nil {
		// Si le registre n'est pas disponible, créer un package par défaut
		return &Package{
			Name:        name,
			Version:     version,
			Description: fmt.Sprintf("Package %s", name),
			Repository:  fmt.Sprintf("https://github.com/%s", name),
		}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &Package{
			Name:        name,
			Version:     version,
			Description: fmt.Sprintf("Package %s", name),
			Repository:  fmt.Sprintf("https://github.com/%s", name),
		}, nil
	}

	var pkg Package
	if err := json.NewDecoder(resp.Body).Decode(&pkg); err != nil {
		return nil, fmt.Errorf("failed to decode package info: %w", err)
	}

	return &pkg, nil
}

func (m *Manager) downloadPackage(pkg *Package, destDir string) error {
	// Pour l'instant, on simule le téléchargement
	// Dans une implémentation complète, on clonerait le repository Git
	// ou téléchargerait depuis une URL

	// Créer un fichier README pour indiquer que le package est installé
	readmePath := filepath.Join(destDir, "README.md")
	readmeContent := fmt.Sprintf("# %s\n\n%s\n\nVersion: %s\nRepository: %s\n",
		pkg.Name, pkg.Description, pkg.Version, pkg.Repository)

	if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
		return fmt.Errorf("failed to create README: %w", err)
	}

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
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read metadata file: %w", err)
	}

	var pkg Package
	if err := yaml.Unmarshal(data, &pkg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal package metadata: %w", err)
	}

	return &pkg, nil
}

func parsePackageSpec(spec string) (name, version string) {
	parts := strings.Split(spec, "@")
	name = parts[0]
	if len(parts) > 1 {
		version = parts[1]
	}
	return name, version
}


func parseGitURL(raw string) (*GitRef, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	// Nom du repo
	base := path.Base(u.Path)
	repo := strings.TrimSuffix(base, ".git")

	// Récupérer le fragment avant de le supprimer
	ref := strings.TrimSpace(u.Fragment)

	// Supprimer le fragment pour reconstruire l'URL sans #
	u.Fragment = ""
	urlWithoutFragment := u.String()

	return &GitRef{
		Name: repo,
		Ref:  ref,
		URL:  urlWithoutFragment,
	}, nil
}
