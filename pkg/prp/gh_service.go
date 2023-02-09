package prp

import (
	"context"
	"fmt"
)

type GhService struct {
	repo GhRepo
}

func NewGhService(repo GhRepo) *GhService {
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
	// fmt.Println("gggggg", inp.OwnerID, inp.RepositoryName)
	// fmt.Printf("******** %+v", ref)
	// fmt.Printf("$$$$$$$$ %+v", err)

	return "Brewfile has been updated in the backup repository", nil
}
