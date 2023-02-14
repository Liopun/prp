package prp

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v50/github"
	"github.com/liopun/prp/pkg/api"
	"github.com/liopun/prp/pkg/prp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var nixCmd = &cobra.Command{
	Use: "nix",
	Short: fmt.Sprintf(CMD_SHORT_MSG, NIX),
	RunE: func(cmd *cobra.Command, args []string) error {
		if ok := prp.IsCommandAvailable("nix-env"); !ok {
			return fmt.Errorf(COMMAND_NOT_DETECTED, NIX)
		}

		if _, err := api.VerifyToken(); err != nil {
			return err
		}

		err := prp.CreateNixDump()
		if err != nil {
			return err
		}

		fmt.Printf(PACKAGE_RESTORE_POINT, "Nix", "prp restore nix")

		ctx := context.Background()
		repo := prp.NewGhRepo(github.NewTokenClient(ctx, viper.GetString("token")))
		service := prp.NewGitService(repo)

		gitRepo := viper.GetString("REPO_NAME")

		if !api.CheckGitRepoExist(gitRepo) {
			res, err := service.AddGitPrivateRepo(ctx, prp.GitRepositoryInput{
				RepositoryName: gitRepo,
				Description: GITHUB_REPO_DESC,
				Private: true,
			})
			if err != nil {
				return err
			}

			fmt.Println(res)
		} else {
			fmt.Println(GITHUB_REPO_FOUND)
		}

		resp, err := service.AddBackupToRepo(ctx, prp.GitBackupInput{
			OwnerID: viper.GetString("user"),
			RepositoryName: gitRepo,
			OwnerName: viper.GetString("name"),
			OwnerEmail: viper.GetString("email"),
			CommitFiles: []string{fmt.Sprintf("%s/Nixfile:Nixfile", viper.GetString("PORT_DIR"))},
			CommitMessage: fmt.Sprintf(GITHUB_REPO_COMMIT_MSG, NIX, time.Now()),
		})
		if err != nil {
			return err
		}

		fmt.Println(resp)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(nixCmd)
}
