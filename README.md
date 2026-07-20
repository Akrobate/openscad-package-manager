# OpenSCAD Package Manager (opm)

[![CI](https://github.com/Akrobate/openscad-package-manager/actions/workflows/tests.yml/badge.svg)](https://github.com/Akrobate/openscad-package-manager/actions/workflows/tests.yml)
[![Build](https://github.com/Akrobate/openscad-package-manager/actions/workflows/release.yml/badge.svg)](https://github.com/Akrobate/openscad-package-manager/actions/workflows/release.yml)

A dependency manager for [OpenSCAD](https://openscad.org/) projects, written in Go. It handles package installation, dependency resolution, and provides rendering utilities for generating STL, PNG, and Markdown documentation from annotated `.scad` files.

## Online documentation

[OpenSCAD Package Manager (opm)](https://openscad-modules.gitlab.io/openscad-package-manager-documentation/)

## Download and install

Pre-built binaries are available on the [Releases](https://github.com/Akrobate/openscad-package-manager/releases) page.

### Linux

```shell
curl -L -O https://github.com/Akrobate/openscad-package-manager/releases/latest/download/opm_linux_amd64.tar.gz
tar xvf opm_linux_amd64.tar.gz
sudo mv opm /usr/local/bin/
```

### Windows

Download and extract the archive from the [latest release](https://github.com/Akrobate/openscad-package-manager/releases/latest), then add `opm.exe` to your `PATH` or you can put the `opm.exe` in `C:\Windows\System32`

## Usage

### Initialize a project

Creates a `scad.json` manifest file in the current directory. You will be prompted for package name, version, description, repository, and author.

```shell
opm init
```

### Install packages

```shell
# Install all dependencies listed in scad.json
opm install

# Install a package (latest version)
opm install https://gitlab.com/openscad-modules/housing.git

# Install a specific tag
opm install https://gitlab.com/openscad-modules/housing.git#0.0.2

# Install a specific branch
opm install https://gitlab.com/openscad-modules/housing.git#develop

# Install a specific commit
opm install https://gitlab.com/openscad-modules/housing.git#5ebc661
```

Packages are cloned into the `openscad_modules/` directory. Sub-dependencies are resolved recursively and their `use`/`include` paths are automatically rewritten.

### Uninstall packages

```shell
opm uninstall
```

Removes the entire `openscad_modules/` directory. Must be run from the root of a project containing a `scad.json` file.

### List installed packages

```shell
opm list
```

Displays all installed packages with their name, version, and commit hash.

### Get package info

```shell
opm info https://gitlab.com/openscad-modules/housing.git
```

Output example:

```
Name: housing
Descrition: housing libraries helper
Latest commit: 5d18bef
Versions:
0.0.1    https://gitlab.com/openscad-modules/housing.git#0.0.1
0.0.2    https://gitlab.com/openscad-modules/housing.git#0.0.2
0.0.3    https://gitlab.com/openscad-modules/housing.git#0.0.3
0.0.4    https://gitlab.com/openscad-modules/housing.git#0.0.4
5d18bef  https://gitlab.com/openscad-modules/housing.git#5d18bef
```

### Search packages

Search available packages from registered repository source lists.

```shell
# List all available packages
opm search

# Search by name
opm search housing
```

### Repository source lists

Manage remote package source lists. Source list configuration is stored in `~/.opm/repository/`.

```shell
# Add a source list
opm repository add https://raw.githubusercontent.com/Akrobate/openscad-package-manager/refs/heads/master/data/sources-list/akrobate.source-list.txt

# Update cached source lists
opm repository update

# Show registered source lists
opm repository sourcelist
```

Source lists can be provided as a plain text file (one repository URL per line) or as a JSON file:

```json
[
    { "name": "housing", "repository": "https://gitlab.com/openscad-modules/housing.git" },
    { "name": "breadboard", "repository": "https://gitlab.com/openscad-modules/breadboard.git" }
]
```

### Rendering

Requires [OpenSCAD](https://openscad.org/) to be installed and available in your `PATH` (for PNG and STL generation).

#### Annotating `.scad` files

Use annotation blocks in your `.scad` files to mark them for rendering:

```scad
/**
 * @png
 * @colorscheme BeforeDawn
 * @view axes
 */
module my_module() {
    cube([10, 10, 10]);
}
```

Available annotations:
- `@png` — marks the file for PNG rendering
- `@stl` — marks the file for STL rendering
- `@colorscheme <name>` — sets the OpenSCAD color scheme (e.g. `BeforeDawn`, `Cornfield`)
- `@view <type>` — sets the camera view (e.g. `axes`)

#### Render commands

```shell
# Generate PNG files (output to opm_png_files/)
opm render process png

# Generate STL files (output to opm_stl_files/)
opm render process stl

# Generate a Markdown preview document from rendered PNGs
opm render process md

# List files that would be rendered
opm render list png
opm render list stl
```

## Package manifest (`scad.json`)

```json
{
  "name": "my-project",
  "version": "1.0.0",
  "description": "An OpenSCAD project",
  "author": "Your Name",
  "repository": "https://gitlab.com/your-user/my-project.git",
  "dependencies": {
    "housing": "https://gitlab.com/openscad-modules/housing.git#0.0.4",
    "breadboard": "https://gitlab.com/openscad-modules/breadboard.git"
  }
}
```

## Development

### Requirements

- Go 1.23 or newer

### Build (linux)

```bash
go build -o opm
```

### Build and install locally (linux)

```bash
go build -o opm && sudo cp opm /usr/local/bin/
```

### Testing

```bash
go test ./...
```

### Build (windows)

```bash
GOOS=windows GOARCH=amd64 go build -o opm.exe
```

#### Coverage

```bash
go test -cover ./...
```

##### Generate HTML coverage report

```bash
go test -coverprofile=coverage.out -coverpkg=./... ./...
go tool cover -html=coverage.out -o coverage.html
```

## Architecture

```
├── cmd/                    CLI commands (cobra)
│   ├── root.go             Root command definition
│   ├── init.go             opm init
│   ├── install.go          opm install
│   ├── uninstall.go        opm uninstall
│   ├── list.go             opm list
│   ├── info.go             opm info
│   ├── search.go           opm search
│   ├── render.go           opm render (parent command)
│   ├── render_process.go   opm render process
│   ├── render_list.go      opm render list
│   ├── repository.go       opm repository (parent command)
│   ├── repository_add.go   opm repository add
│   ├── repository_update.go opm repository update
│   └── repository_sourcelist.go opm repository sourcelist
├── pkg/
│   ├── manager/            Package management (install, uninstall, dependency resolution)
│   ├── renderer/           STL, PNG, and Markdown rendering from annotated .scad files
│   └── repository/         Remote source list management and search
├── internal/utils/         Shared utilities (git operations, file helpers, annotation parsing)
├── data/sources-list/      Example source list files
├── main.go                 Entry point
└── .goreleaser.yml         Release configuration
```

## Licence

MIT
