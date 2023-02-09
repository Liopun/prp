# Package Restore Point CLI tool
A simple package backup manager CLI tool that restores your brew packages to a previously saved state point.

This is useful when setting up your new MacOS and you need to reinstall all of your homebrew packages.

This package use your github as data store for keeping your saved state for future use.

- `prp gh TOKEN_HERE`

## Supported OS
- MacOS

## Installation
### Homebrew
`brew install liopun/brew/prp`

### Binary
The latest release can be found [here](https://github.com/Liopun/prp/releases) and it needs to be in you path

### Source
1. Clone this repository
2. Build with `make build ver="v0.15"`
3. Run `./.dist/prp -h` or Copy `prp` file to your `$PATH` and use it from there.

## Supported Package Managers
- Homebrew
    - `backup` your current homebrew state -> `prp brew`
    - `restore` your previously saved state -> `prp restore brew`

## Roadmap
- Support more package managers
    - macports
    - nix
    - pkgsrc

## Adding To This Project
1. Clone this repository
2. Add your branch `git checkout -b BRANCH_NAME_HERE main`
3. Fetch dependencies `go mod download`
3. Implement your changes
4. Run `make build ver="v0.16"` to build the project
5. Happy hacking!

## Publishing
1. `make release ver="v0.16"`