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

### Info about a package

```shell
opm info https://gitlab.com/openscad-modules/housing.git
```

```
Name: housing
Descrition: housing libraries helper
Latest commit: 5d18bef
Versions:
0.0.1	 https://gitlab.com/openscad-modules/housing.git#0.0.1
0.0.2	 https://gitlab.com/openscad-modules/housing.git#0.0.2
0.0.3	 https://gitlab.com/openscad-modules/housing.git#0.0.3
0.0.4	 https://gitlab.com/openscad-modules/housing.git#0.0.4
5d18bef	 https://gitlab.com/openscad-modules/housing.git#5d18bef
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
go test -coverprofile=coverage.out -coverpkg=./... ./...
go tool cover -html=coverage.out -o coverage.html

```


## Architecture

- `cmd/`: Commands CLI (install, uninstall, list, search)
- `pkg/manager/`: Business rules of package management
- `internal/utils`: Commons functions
- `main.go`: Entry point


## Generating local documentation

### Install venv

```shell
python3 -m virtualenv venv
source venv/bin/activate
```

### Install requirements

```shell
pip install mkdocs
pip install mkdocs-material
```

### Serve documentation local

```shell
mkdocs serve
```


## Roadmap @todo

- [x] Info Read module's package scad.json file
- [ ] Documentation Jekyl
- [x] Normalize tmp folder when installing package
- [x] Do not create tmp folder in .opm on init manager
- [x] Info command normalize output (tag + full specification url)
- [x] Build CI to build packages on tags

## Licence

MIT

