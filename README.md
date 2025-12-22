# OpenSCAD Package Manager (opm)

Un gestionnaire de paquets pour OpenSCAD écrit en Go.

## Installation

```bash
go build -o opm
sudo mv opm /usr/local/bin/
```

## Utilisation

### Installer un package

```bash
opm install BOSL2
opm install package@1.0.0
opm install github.com/user/repo
```

### Désinstaller un package

```bash
opm uninstall BOSL2
```

### Lister les packages installés

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

