package prp

import (
	"context"
	"fmt"

	"github.com/google/go-github/v50/github"
)

type GhRepo interface {
	AddGitPrivateRepo(ctx context.Context, inp GitRepositoryInput) (string, error)
}

type GhRepository struct {
	client *github.Client
}

// const (
// 	repoName string = "prp-backup-%s"
// 	description string = "This is an automatic created repo for backing up your package bundle dump files. PRP uses this repository by default to restore your packages. It's private by default, but you can change this if you wish to share your bundles files with others."
// )

type GitRepoInput struct {
	Name string
	Description string
}

func NewGhRepo(client *github.Client) *GhRepository {
	return &GhRepository{
		client: client,
	}
}

func (r *GhRepository) AddGitPrivateRepo(ctx context.Context, inp GitRepositoryInput) (string, error) {
	gitRepo := &github.Repository{
		Name: &inp.RepositoryName,
		Private: &inp.Private,
		Description: &inp.Description,
	}
	repo, _, err := r.client.Repositories.Create(ctx, "", gitRepo)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("New repo created: %s", repo.GetName()), nil
}
