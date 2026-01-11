# OpenSCAD Package Manager (opm)

Dependencies manager for OpenSCAD wrotten in Go.

## Install

```bash
go build -o opm
sudo mv opm /usr/local/bin/
```

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

## Développement

### Requirements

- Go 1.21 or newer

### Compilation

```bash
go build -o opm
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

## Roadmap

- [ ] Documentation Jekyl
- [ ] Build CI to build packages on tags

## Licence

MIT

