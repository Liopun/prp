package prp

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/go-github/v50/github"
	"github.com/liopun/prp/pkg/api"
	"github.com/liopun/prp/pkg/prp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var portCmd = &cobra.Command{
	Use: "port",
	Short: "Create a restoring point for installed macports package",
	RunE: func(cmd *cobra.Command, args []string) error {
		if ok := prp.IsCommandAvailable("brew"); !ok {
			return errors.New("macports not detected on this machine, install Macports first to get started")
		}

		if _, err := api.VerifyToken(); err != nil {
			return err
		}

		err := prp.CreatePortDump()
		if err != nil {
			return err
		}

		fmt.Println("Macports package restore point has been successfully created! You will need to run 'prp restore port' in the future")

		ctx := context.Background()
		repo := prp.NewGhRepo(github.NewTokenClient(ctx, viper.GetString("token")))
		service := prp.NewGitService(repo)

		gitRepo := viper.GetString("REPO_NAME")

		if !api.CheckGitRepoExist(gitRepo) {
			res, err := service.AddGitPrivateRepo(ctx, prp.GitRepositoryInput{
				RepositoryName: gitRepo,
				Description: "This is an automatic created repo for backing up your package bundle dump files. PRP uses this repository to restore your packages. It's private by default, but you can change this if you wish to share your bundles files with others.",
				Private: true,
			})
			if err != nil {
				return err
			}

			fmt.Println(res)
		} else {
			fmt.Println("PRP backup git repository found: no need to create one...")
		}

		resp, err := service.AddBackupToRepo(ctx, prp.GitBackupInput{
			OwnerID: viper.GetString("user"),
			RepositoryName: gitRepo,
			OwnerName: viper.GetString("name"),
			OwnerEmail: viper.GetString("email"),
			CommitFiles: []string{fmt.Sprintf("%s/Portfile:Portfile", viper.GetString("PORT_DIR"))},
			CommitMessage: fmt.Sprintf("New macport bundle file - %v", time.Now()),
		})
		if err != nil {
			return err
		}

		fmt.Println(resp)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(portCmd)
}
