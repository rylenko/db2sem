package service

import (
	"context"

	"db2sem/internal/domain"
	repodto "db2sem/internal/repo/dto"
)

type repo interface {
	InsertSport(ctx context.Context, name string) error
	InsertSportsman(ctx context.Context, req repodto.InsertSportsmanRequest) error
	DeleteSportByID(ctx context.Context, sportID int64) error
	DeleteSportsmanByID(ctx context.Context, sportsmanID int64) error
	GetClubs(ctx context.Context) ([]domain.Club, error)
	GetPrizeWinnersByTournamentID(ctx context.Context, tournamentID int64) ([]domain.PrizeWinner, error)
	GetSportsmanByID(ctx context.Context, sportsmanID int64) (*domain.Sportsman, error)
	GetSportsmen(ctx context.Context) ([]domain.Sportsman, error)
	GetSportsmenInvolvedInSeveralSports(ctx context.Context) ([]domain.Sportsman, error)
	GetSportByID(ctx context.Context, sportID int64) (*domain.Sport, error)
	GetSports(ctx context.Context) ([]domain.Sport, error)
	GetTournaments(ctx context.Context) ([]domain.Tournament, error)
	GetTrainersBySportsmanID(ctx context.Context, sportsmanID int64) ([]domain.Trainer, error)
	GetTrainersBySportID(ctx context.Context, sportID int64) ([]domain.Trainer, error)
	UpdateSportByID(ctx context.Context, req repodto.UpdateSportByIDRequest) error
	UpdateSportsmanByID(ctx context.Context, req repodto.UpdateSportsmanByIDRequest) error
}
