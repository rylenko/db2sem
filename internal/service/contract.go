package service

import (
	"context"
	"time"

	"db2sem/internal/domain"
	repodto "db2sem/internal/repo/dto"
)

type repo interface {
	InsertSport(ctx context.Context, name string) error
	InsertSportsman(ctx context.Context, req repodto.InsertSportsmanRequest) error
	DeleteSportByID(ctx context.Context, sportID int64) error
	DeleteSportsmanByID(ctx context.Context, sportsmanID int64) error
	GetClubActiveSportsmenCountsForPeriod(
		ctx context.Context, startAt, endAt time.Time) ([]domain.ClubSportsmenCount, error)
	GetInactiveSportsmenForPeriod(ctx context.Context, startAt, endAt time.Time) ([]domain.Sportsman, error)
	GetClubs(ctx context.Context) ([]domain.Club, error)
	GetOrganizers(ctx context.Context) ([]domain.Organizer, error)
	GetPlaces(ctx context.Context) ([]domain.Place, error)
	GetPrizeWinnersByTournamentID(ctx context.Context, tournamentID int64) ([]domain.PrizeWinner, error)
	GetSportsmanByID(ctx context.Context, sportsmanID int64) (*domain.Sportsman, error)
	GetSportsmen(ctx context.Context) ([]domain.Sportsman, error)
	GetSportsmenBySportID(ctx context.Context, sportID int64, minRank *int16) ([]domain.RankedSportsman, error)
	GetSportsmenByTrainerID(ctx context.Context, trainerID int64, minRank *int16) ([]domain.RankedSportsman, error)
	GetSportsmenInvolvedInSeveralSports(ctx context.Context) ([]domain.Sportsman, error)
	GetSportByID(ctx context.Context, sportID int64) (*domain.Sport, error)
	GetSports(ctx context.Context) ([]domain.Sport, error)
	GetTournaments(ctx context.Context) ([]domain.Tournament, error)
	GetTournamentsByPlaceID(ctx context.Context, placeID int64, sportID *int64) ([]domain.Tournament, error)
	GetTournamentsForPeriod(ctx context.Context, startAt, endAt time.Time, organizerID *int64) ([]domain.Tournament, error)
	GetTrainers(ctx context.Context) ([]domain.Trainer, error)
	GetTrainersBySportsmanID(ctx context.Context, sportsmanID int64) ([]domain.Trainer, error)
	GetTrainersBySportID(ctx context.Context, sportID int64) ([]domain.Trainer, error)
	UpdateSportByID(ctx context.Context, req repodto.UpdateSportByIDRequest) error
	UpdateSportsmanByID(ctx context.Context, req repodto.UpdateSportsmanByIDRequest) error
}
