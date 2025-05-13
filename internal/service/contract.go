package service

import (
	"context"

	"db2sem/internal/domain"
)

type repo interface {
	DeleteSportsmanByID(ctx context.Context, sportsmanID int64) error
	GetSportsmanByID(ctx context.Context, sportsmanID int64) (*domain.Sportsman, error)
	GetSportsmen(ctx context.Context) ([]domain.Sportsman, error)
	GetSportsmenInvolvedInSeveralSports(ctx context.Context) ([]domain.Sportsman, error)
	GetSportNames(ctx context.Context) ([]string, error)
}
