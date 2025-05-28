package transport

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"

	"db2sem/internal/domain"
	servicedto "db2sem/internal/service/dto"
)

type requestReader interface {
	ReadAndValidateFiberBody(fiberCtx *fiber.Ctx, request any) error
}

type service interface {
	CreateSport(ctx context.Context, name string) error
	CreateSportsman(ctx context.Context, req servicedto.CreateSportsmanRequest) error
	DeleteSportByID(ctx context.Context, sportID int64) error
	DeleteSportsmanByID(ctx context.Context, sportsmanID int64) error
	GetInactiveSportsmenForPeriod(ctx context.Context, startAt, endAt time.Time) ([]domain.Sportsman, error)
	GetOrganizerTournamentCountsForPeriod(
		ctx context.Context,
		startAt time.Time,
		endAt time.Time,
	) ([]domain.OrganizerTournamentsCount, error)
	GetClubActiveSportsmenCountsForPeriod(
		ctx context.Context, startAt, endAt time.Time) ([]domain.ClubSportsmenCount, error)
	GetClubs(ctx context.Context) ([]domain.Club, error)
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
	GetOrganizers(ctx context.Context) ([]domain.Organizer, error)
	GetPlaces(ctx context.Context) ([]domain.Place, error)
	GetTrainers(ctx context.Context) ([]domain.Trainer, error)
	GetTrainersBySportsmanID(ctx context.Context, sportsmanID int64) ([]domain.Trainer, error)
	GetTrainersBySportID(ctx context.Context, sportID int64) ([]domain.Trainer, error)
	UpdateSportByID(ctx context.Context, req servicedto.UpdateSportByIDRequest) error
	UpdateSportsmanByID(ctx context.Context, req servicedto.UpdateSportsmanByIDRequest) error
}
