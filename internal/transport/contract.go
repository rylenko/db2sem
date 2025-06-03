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
	ReadAndValidateFiberQuery(fiberCtx *fiber.Ctx, request any) error
}

type service interface {
	GetPlacesWithTournamentDatesForPeriod(
		ctx context.Context, startAt, endAt time.Time) ([]domain.PlaceWithTournamentDates, error)
	GetArenas(ctx context.Context, req servicedto.GetArenasRequest) ([]domain.Arena, error)
	GetStadiums(ctx context.Context, req servicedto.GetStadiumsRequest) ([]domain.Stadium, error)
	GetGyms(ctx context.Context, req servicedto.GetGymsRequest) ([]domain.Gym, error)
	GetCourts(ctx context.Context, req servicedto.GetCourtsRequest) ([]domain.Court, error)
	CreateSport(ctx context.Context, name string) error
	CreateSportsman(ctx context.Context, req servicedto.CreateSportsmanRequest) error
	CreateArena(ctx context.Context, req servicedto.CreateArenaRequest) error
	DeleteSportByID(ctx context.Context, sportID int64) error
	DeletePlaceByID(ctx context.Context, placeID int64) error
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
	GetArenaByID(ctx context.Context, id int64) (*domain.Arena, error)
	GetStadiumByID(ctx context.Context, id int64) (*domain.Stadium, error)
	GetCourtByID(ctx context.Context, id int64) (*domain.Court, error)
	GetGymByID(ctx context.Context, id int64) (*domain.Gym, error)
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
	UpdateArenaByID(ctx context.Context, req servicedto.UpdateArenaByIDRequest) error
	UpdateStadiumByID(ctx context.Context, req servicedto.UpdateStadiumByIDRequest) error
	UpdateCourtByID(ctx context.Context, req servicedto.UpdateCourtByIDRequest) error
	UpdateGymByID(ctx context.Context, req servicedto.UpdateGymByIDRequest) error
	UpdateSportsmanByID(ctx context.Context, req servicedto.UpdateSportsmanByIDRequest) error
}
