package service

import (
	"context"
	"time"

	"db2sem/internal/domain"
	repodto "db2sem/internal/repo/dto"
	"db2sem/internal/service/dto"
)

type Service struct {
	repo repo
}

func New(repo repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateSport(ctx context.Context, name string) error {
	return s.repo.InsertSport(ctx, name)
}

func (s *Service) CreateSportsman(ctx context.Context, req dto.CreateSportsmanRequest) error {
	return s.repo.InsertSportsman(ctx, repodto.InsertSportsmanRequest{
		Name:      req.Name,
		HeightCm:  req.HeightCm,
		BirthDate: req.BirthDate,
		WeightKg:  req.WeightKg,
		ClubID:    req.ClubID,
		SportIDs:  req.SportIDs,
	})
}

func (s *Service) DeleteSportByID(ctx context.Context, sportID int64) error {
	return s.repo.DeleteSportByID(ctx, sportID)
}

func (s *Service) DeleteSportsmanByID(ctx context.Context, sportsmanID int64) error {
	return s.repo.DeleteSportsmanByID(ctx, sportsmanID)
}

func (s *Service) GetClubActiveSportsmenCountsForPeriod(
	ctx context.Context,
	startAt time.Time,
	endAt time.Time,
) ([]domain.ClubSportsmenCount, error) {
	return s.repo.GetClubActiveSportsmenCountsForPeriod(ctx, startAt, endAt)
}

func (s *Service) GetClubs(ctx context.Context) ([]domain.Club, error) {
	return s.repo.GetClubs(ctx)
}

func (s *Service) GetPrizeWinnersByTournamentID(ctx context.Context, tournamentID int64) ([]domain.PrizeWinner, error) {
	return s.repo.GetPrizeWinnersByTournamentID(ctx, tournamentID)
}

func (s *Service) GetSportsmanByID(ctx context.Context, sportsmanID int64) (*domain.Sportsman, error) {
	return s.repo.GetSportsmanByID(ctx, sportsmanID)
}

func (s *Service) GetSportsmen(ctx context.Context) ([]domain.Sportsman, error) {
	return s.repo.GetSportsmen(ctx)
}

func (s *Service) GetSportsmenBySportID(
	ctx context.Context,
	sportID int64,
	minRank *int16,
) ([]domain.RankedSportsman, error) {
	return s.repo.GetSportsmenBySportID(ctx, sportID, minRank)
}

func (s *Service) GetSportsmenByTrainerID(
	ctx context.Context,
	trainerID int64,
	minRank *int16,
) ([]domain.RankedSportsman, error) {
	return s.repo.GetSportsmenByTrainerID(ctx, trainerID, minRank)
}

func (s *Service) GetSportsmenInvolvedInSeveralSports(ctx context.Context) ([]domain.Sportsman, error) {
	return s.repo.GetSportsmenInvolvedInSeveralSports(ctx)
}

func (s *Service) GetSportByID(ctx context.Context, sportID int64) (*domain.Sport, error) {
	return s.repo.GetSportByID(ctx, sportID)
}

func (s *Service) GetSports(ctx context.Context) ([]domain.Sport, error) {
	return s.repo.GetSports(ctx)
}

func (s *Service) GetTournaments(ctx context.Context) ([]domain.Tournament, error) {
	return s.repo.GetTournaments(ctx)
}

func (s *Service) GetTournamentsByPlaceID(
	ctx context.Context,
	placeID int64,
	sportID *int64,
) ([]domain.Tournament, error) {
	return s.repo.GetTournamentsByPlaceID(ctx, placeID, sportID)
}

func (s *Service) GetTournamentsForPeriod(
	ctx context.Context,
	startAt time.Time,
	endAt time.Time,
	organizerID *int64,
) ([]domain.Tournament, error) {
	return s.repo.GetTournamentsForPeriod(ctx, startAt, endAt, organizerID)
}

func (s *Service) GetOrganizers(ctx context.Context) ([]domain.Organizer, error) {
	return s.repo.GetOrganizers(ctx)
}

func (s *Service) GetPlaces(ctx context.Context) ([]domain.Place, error) {
	return s.repo.GetPlaces(ctx)
}

func (s *Service) GetTrainers(ctx context.Context) ([]domain.Trainer, error) {
	return s.repo.GetTrainers(ctx)
}

func (s *Service) GetTrainersBySportsmanID(ctx context.Context, sportsmanID int64) ([]domain.Trainer, error) {
	return s.repo.GetTrainersBySportsmanID(ctx, sportsmanID)
}

func (s *Service) GetTrainersBySportID(ctx context.Context, sportID int64) ([]domain.Trainer, error) {
	return s.repo.GetTrainersBySportID(ctx, sportID)
}

func (s *Service) UpdateSportByID(ctx context.Context, req dto.UpdateSportByIDRequest) error {
	return s.repo.UpdateSportByID(ctx, repodto.UpdateSportByIDRequest{
		ID:   req.ID,
		Name: req.Name,
	})
}

func (s *Service) UpdateSportsmanByID(ctx context.Context, req dto.UpdateSportsmanByIDRequest) error {
	return s.repo.UpdateSportsmanByID(ctx, repodto.UpdateSportsmanByIDRequest{
		ID:        req.ID,
		Name:      req.Name,
		HeightCm:  req.HeightCm,
		BirthDate: req.BirthDate,
		WeightKg:  req.WeightKg,
		ClubID:    req.ClubID,
		SportIDs:  req.SportIDs,
	})
}
