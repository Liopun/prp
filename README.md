# Package Restore Point CLI tool
A convenient solution for backing up and restoring your installed packages. This tool offers the portability of your current installed packages and utilizes your GitHub account as a data store for maintaining a record of your packages for future use.

## Use Cases
- Moving to a brand new device (MacOs, Linux)
- Duplicating your Homebrew/Macports/Nix setup across different devices

## Supported On
- MacOs
- Linux

## Supported Package Managers
- [Homebrew](https://brew.sh/)
- [Macports](https://www.macports.org/)
- [Nix](https://nixos.wiki/wiki/Main_Page)

## Installation
### Homebrew
`brew install liopun/brew/prp`

### Binary
The latest release of PRP CLI tool can be found [here](https://github.com/Liopun/prp/releases), and it must be in your PATH for effective usage.

### Source
1. Clone this repository
2. Build with `make build ver="v0.15"`
3. Run `./.dist/prp -h` or Copy `prp` file to your `$PATH` and use it from there.

## Get Started
- Run the following command to authenticate with Github `prp gh TOKEN_HERE`
- You can find out more information about Github personal tokens [here](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)

## Commands
- Homebrew
    - Backup your current Homebrew packages: `prp brew`
    - Restore/Install your previously saved Homebrew packages to another system: `prp restore brew`
- Macports
    - Backup your current Macports packages: `prp port`
    - Restore/Install your previously saved Macports packages to another system: `prp restore port`
- NixOS
    - Backup your current NixOs packages: `prp nix`
    - Restore/Install your previously saved NixOs packages to another system: `prp restore nix`

## Roadmap
- Support more package managers
    - macports [x]
    - nix [x]
    - pkgsrc

## Adding To This Project
1. Clone this repository
2. Add your branch `git checkout -b BRANCH_NAME_HERE main`
3. Fetch dependencies `go mod download`
4. Implement your changes
5. Run `make build ver="v0.16"` to build the project
6. Happy hacking!

## Publishing
`make release ver="v0.16"`
