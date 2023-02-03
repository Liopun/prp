package prp

import "context"

type GhService struct {
	repo GhRepo
}

func NewGhService(repo GhRepo) *GhService {
	return &GhService{repo}
}

func (s *GhService) AddGitPrivateRepo(ctx context.Context, inp GitRepositoryInput) (string, error) {
	return s.repo.AddGitPrivateRepo(ctx, inp)
}
