package prp

import (
	"context"

	"github.com/google/go-github/v50/github"
)

type Repo interface {
	AddGithubRepo(ctx context.Context) (string, error)
	SignOut(ctx context.Context) (string, error)
}

type Repository struct {
	client *github.Client
}

func NewRepository(client *github.Client) *Repository {
	return &Repository{
		client,
	}
}

func (r *Repository) AddGithubRepo(ctx context.Context) (string, error) {
	return "", nil
}

func (r *Repository) SignOut(ctx context.Context) (string, error) {
	return "", nil
}
