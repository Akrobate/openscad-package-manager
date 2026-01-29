#💡 Usage

After installation, you can start managing your OpenSCAD dependencies.

## 🛠 Initialize a project

```shell
opm init
```

This creates a configuration file (for example scad.json) in the current directory

## 📋 Available Commands

Command	Description

- `opm install`	Install all dependencies defined in the config file
- `opm install <URL>`	Install a package from a Git repository
- `opm install <URL>#<version>`	Install a specific version
- `opm uninstall`	Uninstall all dependencies
- `opm list`	List installed packages
- `opm info <URL>`	Display information about a package

## 🔍 Commands examples

Install a module

```shell
opm install https://gitlab.com/openscad-modules/housing.git
```

Install a specific version

```shell
opm install https://gitlab.com/openscad-modules/housing.git#0.0.2
```

List installed packages

```shell
opm list
```

Show package information

```shell
opm info https://gitlab.com/openscad-modules/housing.git
```



