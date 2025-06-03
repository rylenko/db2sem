package service

import (
	"context"
	"time"

	"db2sem/internal/domain"
	repodto "db2sem/internal/repo/dto"
)

type repo interface {
	GetPlacesWithTournamentDatesForPeriod(
		ctx context.Context, startAt, endAt time.Time) ([]domain.PlaceWithTournamentDates, error)
	GetArenas(ctx context.Context, req repodto.GetArenasRequest) ([]domain.Arena, error)
	GetStadiums(ctx context.Context, req repodto.GetStadiumsRequest) ([]domain.Stadium, error)
	GetGyms(ctx context.Context, req repodto.GetGymsRequest) ([]domain.Gym, error)
	GetCourts(ctx context.Context, req repodto.GetCourtsRequest) ([]domain.Court, error)
	InsertSport(ctx context.Context, name string) error
	InsertClub(ctx context.Context, name string) error
	InsertOrganizer(ctx context.Context, name string, location *string) error
	InsertArena(ctx context.Context, req repodto.InsertArenaRequest) error
	InsertStadium(ctx context.Context, req repodto.InsertStadiumRequest) error
	InsertSportsman(ctx context.Context, req repodto.InsertSportsmanRequest) error
	DeleteSportByID(ctx context.Context, sportID int64) error
	DeleteClubByID(ctx context.Context, sportID int64) error
	DeleteOrganizerByID(ctx context.Context, organizerID int64) error
	DeletePlaceByID(ctx context.Context, placeID int64) error
	DeleteSportsmanByID(ctx context.Context, sportsmanID int64) error
	GetClubActiveSportsmenCountsForPeriod(
		ctx context.Context, startAt, endAt time.Time) ([]domain.ClubSportsmenCount, error)
	GetOrganizerTournamentCountsForPeriod(
		ctx context.Context,
		startAt time.Time,
		endAt time.Time,
	) ([]domain.OrganizerTournamentsCount, error)
	GetInactiveSportsmenForPeriod(ctx context.Context, startAt, endAt time.Time) ([]domain.Sportsman, error)
	GetClubs(ctx context.Context) ([]domain.Club, error)
	GetOrganizers(ctx context.Context) ([]domain.Organizer, error)
	GetPlaces(ctx context.Context) ([]domain.Place, error)
	GetPrizeWinnersByTournamentID(ctx context.Context, tournamentID int64) ([]domain.PrizeWinner, error)
	GetSportsmanByID(ctx context.Context, sportsmanID int64) (*domain.Sportsman, error)
	GetOrganizerByID(ctx context.Context, sportsmanID int64) (*domain.Organizer, error)
	GetSportsmen(ctx context.Context) ([]domain.Sportsman, error)
	GetSportsmenBySportID(ctx context.Context, sportID int64, minRank *int16) ([]domain.RankedSportsman, error)
	GetSportsmenByTrainerID(ctx context.Context, trainerID int64, minRank *int16) ([]domain.RankedSportsman, error)
	GetSportsmenInvolvedInSeveralSports(ctx context.Context) ([]domain.Sportsman, error)
	GetSportByID(ctx context.Context, sportID int64) (*domain.Sport, error)
	GetClubByID(ctx context.Context, sportID int64) (*domain.Club, error)
	GetArenaByID(ctx context.Context, id int64) (*domain.Arena, error)
	GetStadiumByID(ctx context.Context, id int64) (*domain.Stadium, error)
	GetCourtByID(ctx context.Context, id int64) (*domain.Court, error)
	GetGymByID(ctx context.Context, id int64) (*domain.Gym, error)
	GetSports(ctx context.Context) ([]domain.Sport, error)
	GetTournaments(ctx context.Context) ([]domain.Tournament, error)
	GetTournamentsByPlaceID(ctx context.Context, placeID int64, sportID *int64) ([]domain.Tournament, error)
	GetTournamentsForPeriod(ctx context.Context, startAt, endAt time.Time, organizerID *int64) ([]domain.Tournament, error)
	GetTrainers(ctx context.Context) ([]domain.Trainer, error)
	GetTrainersBySportsmanID(ctx context.Context, sportsmanID int64) ([]domain.Trainer, error)
	GetTrainersBySportID(ctx context.Context, sportID int64) ([]domain.Trainer, error)
	UpdateSportByID(ctx context.Context, req repodto.UpdateSportByIDRequest) error
	UpdateClubByID(ctx context.Context, req repodto.UpdateClubByIDRequest) error
	UpdateOrganizerByID(ctx context.Context, req repodto.UpdateOrganizerByIDRequest) error
	UpdateArenaByID(ctx context.Context, req repodto.UpdateArenaByIDRequest) error
	UpdateStadiumByID(ctx context.Context, req repodto.UpdateStadiumByIDRequest) error
	UpdateCourtByID(ctx context.Context, req repodto.UpdateCourtByIDRequest) error
	UpdateGymByID(ctx context.Context, req repodto.UpdateGymByIDRequest) error
	UpdateSportsmanByID(ctx context.Context, req repodto.UpdateSportsmanByIDRequest) error
}
