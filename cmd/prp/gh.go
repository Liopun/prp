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
		_, err := api.VerifyToken()
		if api.IsTokenAvailable() && api.IsTokenUserAvailable() && err == nil {
			return fmt.Errorf("your current API Token is still valid. If you'd like to use a different API token, use %s first", "`logout`(./prp logout)")
		}

		if err := api.SetToken(args[0]); err != nil {
			return err
		}

		res, err := api.VerifyToken()
		if err != nil {
			return err
		}

		if err := api.SetTokenUser(res.Login, res.Email, res.Name); err != nil {
			return err
		}
		
		fmt.Println("github token based authentication has been successful.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(ghCmd)
}