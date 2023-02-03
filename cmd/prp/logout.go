package prp

import (
	"fmt"

	"github.com/liopun/prp/pkg/api"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use: "logout",
	Short: "Logout of prp",
	Long: "Logout of prp, you can sign in again with `./prp gh TOKEN_VALUE`",
	RunE: func(cmd *cobra.Command, args []string) error {
		if api.IsTokenAvailable() && api.VerifyToken() == nil {
			return api.RemoveToken()
		}
		
		fmt.Println("Github auth token has been revoked on this computer. You can sign in again with `./prp gh TOKEN_VALUE`")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}