---
title: OpenSCAD Package Definition (scad.json) – Dependency & Module Management with OPM
description: Learn how to define OpenSCAD packages using scad.json with OpenSCAD Package Manager (OPM). Manage dependencies, install modules, and maintain version control efficiently.
keywords: OpenSCAD, package manager, OPM, scad.json, OpenSCAD modules, dependency management, OpenSCAD libraries, install OpenSCAD packages
image: assets/logo-openscad-package-manager-opm-min.png
---

# Package Definition (scad.json)

## Overview

- 📦 **Package** distribution & dependency management
- 🧩 **Module** OpenSCAD code you include/use

In opm, a package is an OpenSCAD files distributed as a Git repository.

Every package must:

- Be a Git project (GitHub, GitLab, etc.)
- Contain a scad.json file at the root of the repository

The scad.json file describes the package metadata and its dependencies.

### Example `scad.json`
```json
{
  "name": "housing",
  "version": "1.0.0",
  "description": "housing libraries helper",
  "author": "",
  "repository": "",
  "dependencies": {
    "commons": "https://gitlab.com/openscad-modules/commons.git#0.0.1"
  }
}
```

## Creating a scad.json

The scad.json file can be created in two ways:

### 1. Manually

You can write the file yourself using any text editor.

### 2. Using opm init

Running:
```bash
opm init
```

will prompt the user to fill in the required fields and automatically generate the `scad.json` file.

## Dependencies

### Declaring Dependencies

Dependencies are declared in the **dependencies** object of `scad.json`.

Each dependency is defined as:

- A HTTPS Git URL ending with .git
- Optionally followed by a version reference using #

```json
"dependencies": {
  "commons": "https://gitlab.com/openscad-modules/commons.git#<TAG or GIT SHORT COMMIT>"
}
```

### Version Reference

The part after `#` can be:

- A Git tag
- A Git branch
- A short Git commit hash

Examples:
```
https://gitlab.com/openscad-modules/commons.git
https://gitlab.com/openscad-modules/commons.git#0.0.1
https://gitlab.com/openscad-modules/commons.git#master
https://gitlab.com/openscad-modules/commons.git#a1b2c3d
```

If no version is specified, the default branch is used.

### Installing Dependencies

#### Install all dependencies

From a project containing a scad.json, simply run:

```bash
opm install
```

This command:

- Resolves all dependencies
- Installs them into a local directory named openscad_modules

### Dependency Installation Layout

#### Direct Dependencies

Modules directly listed in the dependencies section are installed using their package name:

```
openscad_modules/
└── housing/
```

This guarantees a stable include path, regardless of the version used.

#### Transitive Dependencies (Dependencies of Dependencies)

Dependencies of dependencies are installed with their **Git short commit appended** to the folder name, using `#` as a separator:

```
openscad_modules/
├── housing/
└── commons#a1b2c3d/
```

This design allows:

- Multiple versions of the same dependency to coexist
- Correct version resolution when different packages depend on different versions of the same module

#### Reference Rewriting

`opm` automatically rewrites include/use paths inside dependencies so that:

- Transitive dependencies correctly reference their versioned directory
- OpenSCAD includes continue to work without manual changes

### Using Installed Modules in OpenSCAD

Once dependencies are installed, modules can be used with standard OpenSCAD directives.

**Using `use`**
```
use <openscad_modules/housing/roundedPane.scad>
```

**Using `include`**
```
include <openscad_modules/housing/roundedPane.scad>
```

Choose `use` or `include` depending on your needs.


### Installing a Package with a Forced Version

You can install a specific version of a package directly:

```bash
opm install https://gitlab.com/openscad-modules/housing.git#0.0.1
```

**Important Behavior**

The package will still be installed under:

```
openscad_modules/housing/
```

- The version is NOT appended to the folder name for direct dependencies
- This ensures that your OpenSCAD source code does not need to be updated when switching versions

### Version Pinning Rules Summary

| Dependency Type | Folder Naming |
|---|---|
| **Direct dependency**	| `housing/` |
| **Dependency of a dependency** | `commons#<git-short-commit>` |

This guarantees:

- Stable include paths for your main project
- Strict version pinning for transitive dependencies
