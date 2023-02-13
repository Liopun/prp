package prp

import (
	"fmt"

	"github.com/liopun/prp/pkg/api"
	"github.com/liopun/prp/pkg/prp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


var restoreCmd = &cobra.Command{
	Use: "restore",
	Short: "Restore your packages from a previously created restore point",
	Long: "Restore your packages from a previously created restore point, prp backup files are kept in a private github repo",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := api.VerifyToken()
		if err != nil {
			return err
		}
		if !(api.IsTokenAvailable() && api.IsTokenUserAvailable()) {
			return fmt.Errorf("your token does not exist. try `prp gh GIT_TOKEN` to authenticate")
		}

		fmt.Println("cloning backup repo to a local path...")
		ok, err := api.CloneGitRepoLocal(viper.GetString("user"), viper.GetString("REPO_NAME"))
		if err != nil {
			return err
		}

		if args[0] == "brew" && ok {
			fmt.Println("restoring homebrew packages in progress...")
			if err := prp.RestoreBrewPackages(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("wrong command argument, try prp restore brew")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}