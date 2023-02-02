package prp

import (
	"fmt"

	"github.com/liopun/prp/pkg/prp"
	"github.com/spf13/cobra"
)

var brewCmd = &cobra.Command{
	Use: "brew",
	Short: "Create a restoring point for installed homebrew package",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Generating homebrew dump file...")
		_, err := prp.CreateBrewDump()
		fmt.Println("Brew package restore point has been successfully created! You will need to run 'prp brew restore' in the future")
		return err
	},
}

func init() {
	rootCmd.AddCommand(brewCmd)
}
