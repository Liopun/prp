# Package Restore Point CLI tool
A simple package backup manager CLI tool that restores your brew packages to a previously saved state point.

This is useful when setting up your new MacOS and you need to reinstall all of your homebrew packages.

This package use your github as data store for keeping your saved state for future use.

- prp gh TOKEN

## Supported OS
- MacOS

## Installation
`brew install liopun/prp`

## Supported Package Managers
- Homebrew
    - `backup` your current homebrew state -> `prp brew`
    - `restore` your previously saved state -> `prp restore brew`

## Roadmaps
- Support more package managers
    - [] macports
    - [] nix
    - [] pkgsrc
