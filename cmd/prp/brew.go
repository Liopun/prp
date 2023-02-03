package prp

import (
	"context"
	"fmt"

	"github.com/google/go-github/v50/github"
	"github.com/liopun/prp/pkg/prp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var brewCmd = &cobra.Command{
	Use: "brew",
	Short: "Create a restoring point for installed homebrew package",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Generating homebrew dump file...")
		_, err := prp.CreateBrewDump()
		if err != nil {
			return err
		}

		fmt.Println("Brew package restore point has been successfully created! You will need to run 'prp brew restore' in the future")


		ctx := context.Background()
		repo := prp.NewGhRepo(github.NewTokenClient(ctx, viper.GetString("token")))
		service := prp.NewGhService(repo)

		res, err := service.AddGitPrivateRepo(ctx, prp.GitRepositoryInput{
			RepositoryName: viper.GetString("REPO_NAME"),
			Description: "This is an automatic created repo for backing up your package bundle dump files. PRP uses this repository by default to restore your packages. It's private by default, but you can change this if you wish to share your bundles files with others.",
			Private: true,
		})
		if err != nil {
			return err
		}

		fmt.Println(res)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(brewCmd)
}
