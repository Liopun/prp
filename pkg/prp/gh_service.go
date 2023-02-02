package prp

import "context"

type Service struct {
	repo Repo
}

func NewService(repo Repo) *Service {
	return &Service{repo}
}

func (s *Service) AddGithubRepo(c context.Context) (string, error) {
	return s.repo.AddGithubRepo(c)
}

func (s *Service) SignOut(c context.Context) (string, error) {
	return s.repo.SignOut(c)
}
