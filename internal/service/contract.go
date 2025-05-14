package service

import (
	"context"

	"db2sem/internal/domain"
	repodto "db2sem/internal/repo/dto"
)

type repo interface {
	DeleteSportsmanByID(ctx context.Context, sportsmanID int64) error
	GetClubs(ctx context.Context) ([]domain.Club, error)
	GetSportsmanByID(ctx context.Context, sportsmanID int64) (*domain.Sportsman, error)
	GetSportsmen(ctx context.Context) ([]domain.Sportsman, error)
	GetSportsmenInvolvedInSeveralSports(ctx context.Context) ([]domain.Sportsman, error)
	GetSports(ctx context.Context) ([]domain.Sport, error)
	UpdateSportsmanByID(ctx context.Context, req repodto.UpdateSportsmanByIDRequest) error
}
