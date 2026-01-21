# OpenSCAD Package Manager (opm)

[![CI](https://github.com/Akrobate/openscad-package-manager/actions/workflows/tests.yml/badge.svg)](https://github.com/Akrobate/openscad-package-manager/actions/workflows/tests.yml)
[![Build](https://github.com/Akrobate/openscad-package-manager/actions/workflows/release.yml/badge.svg)](https://github.com/Akrobate/openscad-package-manager/actions/workflows/release.yml)


Openscad dependencies manager wrotten in Go.

## Download and install

### Linux

```shell
curl -L -O https://github.com/Akrobate/openscad-package-manager/releases/download/0.0.2/opm_0.0.2_linux_amd64.tar.gz
tar xvf opm_0.0.2_linux_amd64.tar.gz
sudo mv opm /usr/local/bin/
```

### Windows

Download
`https://github.com/Akrobate/openscad-package-manager/releases/download/0.0.2/opm_0.0.2_windows_amd64.tar.gz`

Extract
`opm_0.0.2_windows_amd64.tar.gz`


## Usage

### Init package file

```shell
opm init
```

### Install a package

```shell
opm install
opm install https://gitlab.com/openscad-modules/housing.git
opm install https://gitlab.com/openscad-modules/housing.git#0.0.2
opm install https://gitlab.com/openscad-modules/housing.git#develop
opm install https://gitlab.com/openscad-modules/housing.git#5ebc661`,
```

### Uninstall all packages

```shell
opm uninstall
```

### List installed packages

```shell
opm list
```

### Package repository source list

```
https://raw.githubusercontent.com/Akrobate/openscad-package-manager/refs/heads/master/data/sources-list/akrobate.source-list.json
```


## Développement

### Requirements

- Go 1.21 or newer

### Build

```bash
go build -o opm
```

### Build and add to bin
```bash
go build -o opm
sudo mv opm /usr/local/bin/
```

### Developping

```bash
go build -o opm && sudo cp opm /usr/local/bin/
```

### Testing

```bash
go test ./...
```

#### Coverage

```bash
go test -cover ./...
```

##### Generate html coverage

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```


## Architecture

- `cmd/`: Commands CLI (install, uninstall, list, search)
- `pkg/manager/`: Business rules of package management
- `internal/utils`: Commons functions
- `main.go`: Entry point

## Roadmap @todo

- [ ] Normalize tmp folder when installing package
- [ ] Do not create tmp folder in .opm on init manager
- [ ] Info command normalize output (tag + full specification url)
- [ ] Info Read module's package scad.json file
- [ ] Documentation Jekyl
- [x] Build CI to build packages on tags

## Licence

MIT

