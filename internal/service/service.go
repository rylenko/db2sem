package service

import (
	"context"

	"db2sem/internal/domain"
)

type Service struct {
	repo repo
}

func New(repo repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) DeleteSportsmanByID(ctx context.Context, sportsmanID int64) error {
	return s.repo.DeleteSportsmanByID(ctx, sportsmanID)
}

func (s *Service) GetSportsmanByID(ctx context.Context, sportsmanID int64) (*domain.Sportsman, error) {
	return s.repo.GetSportsmanByID(ctx, sportsmanID)
}

func (s *Service) GetSportsmen(ctx context.Context) ([]domain.Sportsman, error) {
	return s.repo.GetSportsmen(ctx)
}

func (s *Service) GetSportsmenInvolvedInSeveralSports(ctx context.Context) ([]domain.Sportsman, error) {
	return s.repo.GetSportsmenInvolvedInSeveralSports(ctx)
}

func (s *Service) GetSportNames(ctx context.Context) ([]string, error) {
	return s.repo.GetSportNames(ctx)
}
