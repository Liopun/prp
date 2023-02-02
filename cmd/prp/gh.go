package prp

import (
	"fmt"

	"github.com/liopun/prp/pkg/api"
	"github.com/spf13/cobra"
)

var ghCmd = &cobra.Command{
	Use: "gh",
	Short: "Github token based authentication needed",
	Long: "Github token based authentication needed, prp needs to authenticate with token in order to create a new private repository where your backup files are kept",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if api.IsTokenAvailable() && api.VerifyToken() == nil {
			return fmt.Errorf("your current API Token is still valid. If you'd like to use a different API key, use %s first", "`logout` or `invalidate it`")
		}

		if err := api.SetToken(args[0]); err != nil {
			return err
		}

		if err := api.VerifyToken(); err != nil {
			return err
		}
		
		fmt.Println("github token based authentication has been successful.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(ghCmd)
}