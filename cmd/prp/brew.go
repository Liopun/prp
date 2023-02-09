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

var brewCmd = &cobra.Command{
	Use: "brew",
	Short: "Create a restoring point for installed homebrew package",
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := api.VerifyToken(); err != nil {
			return err
		}

		fmt.Println("Generating homebrew dump file...")
		_, err := prp.CreateBrewDump()
		if err != nil {
			return err
		}

		fmt.Println("Brew package restore point has been successfully created! You will need to run 'prp brew restore' in the future")

		ctx := context.Background()
		repo := prp.NewGhRepo(github.NewTokenClient(ctx, viper.GetString("token")))
		service := prp.NewGhService(repo)

		gitRepo := viper.GetString("REPO_NAME")

		if !api.CheckGitRepoExist(gitRepo) {
			res, err := service.AddGitPrivateRepo(ctx, prp.GitRepositoryInput{
				RepositoryName: gitRepo,
				Description: "This is an automatic created repo for backing up your package bundle dump files. PRP uses this repository by default to restore your brew packages. It's private by default, but you can change this if you wish to share your bundles files with others.",
				Private: true,
			})
			if err != nil {
				return err
			}

			fmt.Println(res)
		} else {
			fmt.Println("Brew backup git repository found: no need to create one...")
		}

		resp, err := service.AddBackupToRepo(ctx, prp.GitBackupInput{
			OwnerID: viper.GetString("user"),
			RepositoryName: gitRepo,
			OwnerName: viper.GetString("name"),
			OwnerEmail: viper.GetString("email"),
			CommitFiles: []string{fmt.Sprintf("%s/Brewfile:Brewfile", viper.GetString("brew_dir"))},
			CommitMessage: fmt.Sprintf("New brew bundle file - %v", time.Now()),
		})
		if err != nil {
			return err
		}

		fmt.Println(resp)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(brewCmd)
}
