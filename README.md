# OpenSCAD Package Manager (opm)

Dependencies manager for OpenSCAD wrotten in Go.

## Install

```bash
go build -o opm
sudo mv opm /usr/local/bin/
```

## Developping

```bash
go build -o opm && sudo cp opm /usr/local/bin/
```

## Testing

```bash
go test ./...
```

### Coverage

```bash
go test -cover ./...
```

#### Generate html coverage

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Usage

### Install a package

```bash
opm install
opm install https://gitlab.com/openscad-modules/housing.git
opm install https://gitlab.com/openscad-modules/housing.git#0.0.2
opm install https://gitlab.com/openscad-modules/housing.git#develop
opm install https://gitlab.com/openscad-modules/housing.git#5ebc661`,
```

### Uninstall all packages

```bash
opm uninstall
```

### List installed packages

```bash
opm list
```

### Rechercher des packages

```bash
opm search BOSL
opm search utility
```

## Configuration

Le fichier de configuration se trouve dans `~/.opm/config.yaml`.

Exemple de configuration:

```yaml
registry: https://registry.openscad-packages.org
```

## Structure des packages

Les packages sont installés dans `~/.opm/packages/`.

Chaque package contient:
- `package.yaml`: Métadonnées du package
- `README.md`: Documentation du package
- Fichiers source OpenSCAD

## Développement

### Prérequis

- Go 1.21 ou supérieur

### Compilation

```bash
go build -o opm
```

### Tests

```bash
go test ./...
```

## Architecture

- `cmd/`: Commandes CLI (install, uninstall, list, search)
- `pkg/manager/`: Logique de gestion des packages
- `main.go`: Point d'entrée de l'application

## Roadmap

- [ ] Support Git pour télécharger les packages
- [ ] Gestion des versions et mise à jour
- [ ] Support des registres personnalisés
- [ ] Intégration avec OpenSCAD (use <opm:package>)
- [ ] Cache intelligent des packages
- [ ] Validation des packages
- [ ] Support des packages privés

## Licence

MIT

