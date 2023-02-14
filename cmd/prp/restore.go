package prp

import (
	"errors"
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
		if ok := prp.IsCommandAvailable("brew"); !ok {
			return errors.New("homebrew not detected on this machine, install Homebrew first to get started")
		}

		_, err := api.VerifyToken()
		if err != nil {
			return err
		}
		if !(api.IsTokenAvailable() && api.IsTokenUserAvailable()) {
			return errors.New("your token does not exist. try `prp gh GIT_TOKEN` to authenticate")
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
		} else if args[0] == "port" {
			fmt.Println("restoring macports packages in progress...")
			if err := prp.RestorePortsPackages(); err != nil {
				return err
			}
		} else {
			return errors.New("wrong command argument, try `prp restore brew` or `prp restore port`")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}