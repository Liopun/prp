package prp

import (
	"fmt"

	"github.com/liopun/prp/api"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout of prp",
	Long:  "Logout of prp, you can sign in again with `prp gh TOKEN_VALUE`",
	RunE: func(_ *cobra.Command, _ []string) error {
		if api.IsTokenAvailable() && api.IsTokenUserAvailable() {
			return api.RemoveToken()
		}

		fmt.Println("you have previously signed out successfully. You can sign in again with `prp gh TOKEN_VALUE`")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
