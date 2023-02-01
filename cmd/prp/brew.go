package prp

import (
	"fmt"

	"github.com/liopun/prp/pkg/prp"
	"github.com/spf13/cobra"
)

var brewCmd = &cobra.Command{
	Use: "brew",
	Short: "create a restoring point for installed homebrew package",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := prp.CreateBrewDump()
		fmt.Println(res)
		return err
	},
}

func init() {
	rootCmd.AddCommand(brewCmd)
}