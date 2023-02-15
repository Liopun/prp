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
	Use:   "restore",
	Short: "Restore your packages from a previously created restore point",
	Long:  "Restore your packages from a previously created restore point, prp backup files are kept in a private github repo",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
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
			if ok := prp.IsCommandAvailable("brew"); !ok {
				return fmt.Errorf(COMMAND_NOT_DETECTED, BREW)
			}

			fmt.Printf(PACKAGE_RESTORE_PROGRESS, BREW)
			if err := prp.RestoreBrewPackages(); err != nil {
				return err
			}
		} else if args[0] == "port" {
			if ok := prp.IsCommandAvailable("port"); !ok {
				return fmt.Errorf(COMMAND_NOT_DETECTED, PORT)
			}

			fmt.Printf(PACKAGE_RESTORE_PROGRESS, PORT)
			if err := prp.RestorePortsPackages(); err != nil {
				return err
			}
		} else if args[0] == NIX {
			if ok := prp.IsCommandAvailable("nix-env"); !ok {
				return fmt.Errorf(COMMAND_NOT_DETECTED, NIX)
			}

			fmt.Printf(PACKAGE_RESTORE_PROGRESS, NIX)
			if err := prp.RestoreNixPackages(); err != nil {
				return err
			}
		} else {
			return errors.New("wrong command argument, try `prp restore brew` or `prp restore port` or `prp restore nix`")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}
