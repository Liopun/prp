package prp

import (
	"fmt"
	"os"

	"github.com/liopun/prp/config"
	"github.com/spf13/cobra"
)

var version = "1.0.0"

const (
	BREW                     = "homebrew"
	PORT                     = "macports"
	NIX                      = "nix"
	CMD_SHORT_MSG            = "Create a restoring point for installed %s packages"
	COMMAND_NOT_DETECTED     = "%s not detected on this machine, install it first to get started"
	PACKAGE_RESTORE_POINT    = "%s package restore point has been successfully created! You will need to run '%s' in the future"
	GITHUB_REPO_FOUND        = "PRP backup git repository found: no need to create one..."
	GITHUB_REPO_DESC         = "This is an automatic created repo for backing up your package bundle dump files. PRP uses this repository to restore your packages. It's private by default, but you can change this if you wish to share your bundles files with others."
	GITHUB_REPO_COMMIT_MSG   = "New %s bundle file - %v"
	PACKAGE_RESTORE_PROGRESS = "restoring %s packages in progress..."
)

var rootCmd = &cobra.Command{
	Use:     "prp",
	Version: version,
	Short:   "prp - package restore point",
	Long:    "A simple backup and restore CLI tool for your favorite package managers.",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	config.Init()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was error running prp CLI: %s", err)
		os.Exit(1)
	}
}
