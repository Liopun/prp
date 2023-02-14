package prp

import (
	"context"
	"fmt"
)

type GhService struct {
	repo GitRepo
}

func NewGitService(repo GitRepo) *GhService {
	return &GhService{repo}
}

func (s *GhService) AddGitPrivateRepo(ctx context.Context, inp GitRepositoryInput) (string, error) {
	return s.repo.AddGitPrivateRepo(ctx, inp)
}

func (s *GhService) AddBackupToRepo(ctx context.Context, inp GitBackupInput) (string, error) {
	inp.BranchName = "master"
	ref, err := s.repo.GetGitRef(ctx, inp)
	if err != nil {
		inp.BranchName = "main"
		ref, err = s.repo.GetGitRef(ctx, inp)
		if err != nil {
			return "", err
		}
	}

	if ref == nil {
		return "", fmt.Errorf("git commit reference could not be created")
	}

	tree, err := s.repo.GetGitTree(ctx, ref, inp)
	if err != nil {
		return "", err
	}

	err = s.repo.PushGitCommit(ctx, ref, tree, inp)
	if err != nil {
		return "", err;
	}

	return fmt.Sprintf("Backup Repository Updated: %s", inp.CommitMessage), nil
}
