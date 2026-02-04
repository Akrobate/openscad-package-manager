---
title: Quick Start using OpenSCAD Package Definition (opm)
description: Get started with OpenSCAD Package Manager (opm) in minutes. Learn how to install opm, define dependencies, and use OpenSCAD packages in your projects.
keywords: Quick Start OPM, package manager, OPM, scad.json, OpenSCAD modules, dependency management, OpenSCAD libraries, install OpenSCAD packages
image: assets/logo-openscad-package-manager-opm-min.png
---

# 🚀 Quick Start

This quick start guide will help you get up and running with OpenSCAD Package Manager (opm) in just a few steps.

## 1. Install opm

Download the latest release for your platform from the GitHub releases page and add it to your system PATH. For example on Linux:

```shell
curl -L -O https://github.com/Akrobate/openscad-package-manager/releases/latest/download/opm_<VERSION>_linux_amd64.tar.gz
tar xvf opm_<VERSION>_linux_amd64.tar.gz
sudo mv opm /usr/local/bin/
```

Replace `<VERSION>` with the appropriate version string.

Once installed, verify the installation:

```shell
opm --version
```

Please refer the [Install OpenScad Package Manager section](download-and-install-opm.md) for more details like other platforms

## 2. Create a new project

In your OpenSCAD project folder, initialize a new package file:

```shell
opm init
```

This command will interactively prompt you for:

- name
- version
- description
- author
- repository URL

It generates a `scad.json` file describing your project and its future dependencies.

Please refer to [OpenScad Packages](openscad-packages.md) for more details about packages

## 3. Add dependencies

### 3.1. Add new dependency

This command download the package and its dependecies and update `scad.json` file with the new dependency

```shell
opm install https://gitlab.com/openscad-modules/housing.git#0.0.1
```

### 3.2. Install all dependencies from existing `scad.json`

From your project root, run:

```shell
opm install
```

## 4. Use packages in your OpenSCAD code

Once installed, include modules from packages like this:

```
use <openscad_modules/commons/some_module.scad>;
```

or

```
include <openscad_modules/commons/another_module.scad>;
```

Use use when you want to import a module for use in your design logic and include when you want to literally include its contents in your file tree.
