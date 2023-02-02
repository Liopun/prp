package prp

import (
	"fmt"
	"os"

	"github.com/liopun/prp/pkg/config"
	"github.com/spf13/cobra"
)

var version = "1.0.0"
var rootCmd = &cobra.Command{
	Use: "prp",
	Version: version,
	Short: "prp - package restore point",
	Long: "A simple common package backup manager CLI tool that restores your package to a previously saved state point.",
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	config.Init()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was error running prp CLI: %s", err)
		os.Exit(1)
	}
}