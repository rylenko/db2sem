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

func (s *Service) GetPlacesWithTournamentDatesForPeriod(
	ctx context.Context,
	startAt time.Time,
	endAt time.Time,
) ([]domain.PlaceWithTournamentDates, error) {
	return s.repo.GetPlacesWithTournamentDatesForPeriod(ctx, startAt, endAt)
}

func (s *Service) GetGyms(ctx context.Context, req dto.GetGymsRequest) ([]domain.Gym, error) {
	return s.repo.GetGyms(ctx, repodto.GetGymsRequest{
		TrainersCount:  req.TrainersCount,
		DumbbellsCount: req.DumbbellsCount,
		HasBathhouse:   req.HasBathhouse,
	})
}

func (s *Service) GetCourts(ctx context.Context, req dto.GetCourtsRequest) ([]domain.Court, error) {
	return s.repo.GetCourts(ctx, repodto.GetCourtsRequest{
		WidthCm:   req.WidthCm,
		LengthCm:  req.LengthCm,
		IsOutdoor: req.IsOutdoor,
	})
}

func (s *Service) GetStadiums(ctx context.Context, req dto.GetStadiumsRequest) ([]domain.Stadium, error) {
	return s.repo.GetStadiums(ctx, repodto.GetStadiumsRequest{
		WidthCm:       req.WidthCm,
		LengthCm:      req.LengthCm,
		MaxSpectators: req.MaxSpectators,
		IsOutdoor:     req.IsOutdoor,
		Coating:       req.Coating,
	})
}

func (s *Service) GetArenas(ctx context.Context, req dto.GetArenasRequest) ([]domain.Arena, error) {
	return s.repo.GetArenas(ctx, repodto.GetArenasRequest{
		RefereesCount:     req.RefereesCount,
		TreadmillLengthCm: req.TreadmillLengthCm,
	})
}

func (s *Service) CreateOrganizer(ctx context.Context, name string, location *string) error {
	return s.repo.InsertOrganizer(ctx, name, location)
}

func (s *Service) CreateTrainer(ctx context.Context, name string) error {
	return s.repo.InsertTrainer(ctx, name)
}

func (s *Service) CreateClub(ctx context.Context, name string) error {
	return s.repo.InsertClub(ctx, name)
}

func (s *Service) CreateSport(ctx context.Context, name string) error {
	return s.repo.InsertSport(ctx, name)
}

func (s *Service) GetOrganizerTournamentCountsForPeriod(
	ctx context.Context,
	startAt time.Time,
	endAt time.Time,
) ([]domain.OrganizerTournamentsCount, error) {
	return s.repo.GetOrganizerTournamentCountsForPeriod(ctx, startAt, endAt)
}

func (s *Service) GetInactiveSportsmenForPeriod(
	ctx context.Context, startAt, endAt time.Time) ([]domain.Sportsman, error) {
	return s.repo.GetInactiveSportsmenForPeriod(ctx, startAt, endAt)
}

func (s *Service) CreateStadium(ctx context.Context, req dto.CreateStadiumRequest) error {
	return s.repo.InsertStadium(ctx, repodto.InsertStadiumRequest{
		Name:          req.Name,
		Location:      req.Location,
		WidthCm:       req.WidthCm,
		LengthCm:      req.LengthCm,
		MaxSpectators: req.MaxSpectators,
		IsOutdoor:     req.IsOutdoor,
		Coating:       req.Coating,
	})
}

func (s *Service) CreateArena(ctx context.Context, req dto.CreateArenaRequest) error {
	return s.repo.InsertArena(ctx, repodto.InsertArenaRequest{
		Name:              req.Name,
		Location:          req.Location,
		RefereesCount:     req.RefereesCount,
		TreadmillLengthCm: req.TreadmillLengthCm,
	})
}

func (s *Service) CreateTournament(ctx context.Context, req dto.CreateTournamentRequest) error {
	return s.repo.InsertTournament(ctx, repodto.InsertTournamentRequest{
		PlaceID:     req.PlaceID,
		OrganizerID: req.PlaceID,
		StartAt:     req.StartAt,
	})
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

func (s *Service) DeleteOrganizerByID(ctx context.Context, sportID int64) error {
	return s.repo.DeleteOrganizerByID(ctx, sportID)
}

func (s *Service) DeleteTrainerByID(ctx context.Context, sportID int64) error {
	return s.repo.DeleteTrainerByID(ctx, sportID)
}

func (s *Service) DeleteClubByID(ctx context.Context, sportID int64) error {
	return s.repo.DeleteClubByID(ctx, sportID)
}

func (s *Service) DeleteSportByID(ctx context.Context, sportID int64) error {
	return s.repo.DeleteSportByID(ctx, sportID)
}

func (s *Service) DeletePlaceByID(ctx context.Context, placeID int64) error {
	return s.repo.DeletePlaceByID(ctx, placeID)
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

func (s *Service) GetOrganizerByID(ctx context.Context, sportID int64) (*domain.Organizer, error) {
	return s.repo.GetOrganizerByID(ctx, sportID)
}

func (s *Service) GetTrainerByID(ctx context.Context, clubID int64) (*domain.Trainer, error) {
	return s.repo.GetTrainerByID(ctx, clubID)
}

func (s *Service) GetClubByID(ctx context.Context, clubID int64) (*domain.Club, error) {
	return s.repo.GetClubByID(ctx, clubID)
}

func (s *Service) GetSportByID(ctx context.Context, sportID int64) (*domain.Sport, error) {
	return s.repo.GetSportByID(ctx, sportID)
}

func (s *Service) GetArenaByID(ctx context.Context, id int64) (*domain.Arena, error) {
	return s.repo.GetArenaByID(ctx, id)
}

func (s *Service) GetStadiumByID(ctx context.Context, id int64) (*domain.Stadium, error) {
	return s.repo.GetStadiumByID(ctx, id)
}

func (s *Service) GetCourtByID(ctx context.Context, id int64) (*domain.Court, error) {
	return s.repo.GetCourtByID(ctx, id)
}

func (s *Service) GetGymByID(ctx context.Context, id int64) (*domain.Gym, error) {
	return s.repo.GetGymByID(ctx, id)
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

func (s *Service) UpdateOrganizerByID(ctx context.Context, req dto.UpdateOrganizerByIDRequest) error {
	return s.repo.UpdateOrganizerByID(ctx, repodto.UpdateOrganizerByIDRequest{
		ID:       req.ID,
		Name:     req.Name,
		Location: req.Location,
	})
}

func (s *Service) UpdateTrainerByID(ctx context.Context, req dto.UpdateTrainerByIDRequest) error {
	return s.repo.UpdateTrainerByID(ctx, repodto.UpdateTrainerByIDRequest{
		ID:   req.ID,
		Name: req.Name,
	})
}

func (s *Service) UpdateClubByID(ctx context.Context, req dto.UpdateClubByIDRequest) error {
	return s.repo.UpdateClubByID(ctx, repodto.UpdateClubByIDRequest{
		ID:   req.ID,
		Name: req.Name,
	})
}

func (s *Service) UpdateSportByID(ctx context.Context, req dto.UpdateSportByIDRequest) error {
	return s.repo.UpdateSportByID(ctx, repodto.UpdateSportByIDRequest{
		ID:   req.ID,
		Name: req.Name,
	})
}

func (s *Service) UpdateGymByID(ctx context.Context, req dto.UpdateGymByIDRequest) error {
	return s.repo.UpdateGymByID(ctx, repodto.UpdateGymByIDRequest{
		ID:             req.ID,
		Name:           req.Name,
		Location:       req.Location,
		TrainersCount:  req.TrainersCount,
		DumbbellsCount: req.DumbbellsCount,
		HasBathhouse:   req.HasBathhouse,
	})
}

func (s *Service) UpdateCourtByID(ctx context.Context, req dto.UpdateCourtByIDRequest) error {
	return s.repo.UpdateCourtByID(ctx, repodto.UpdateCourtByIDRequest{
		ID:        req.ID,
		Name:      req.Name,
		Location:  req.Location,
		WidthCm:   req.WidthCm,
		LengthCm:  req.LengthCm,
		IsOutdoor: req.IsOutdoor,
	})
}

func (s *Service) UpdateStadiumByID(ctx context.Context, req dto.UpdateStadiumByIDRequest) error {
	return s.repo.UpdateStadiumByID(ctx, repodto.UpdateStadiumByIDRequest{
		ID:            req.ID,
		Name:          req.Name,
		Location:      req.Location,
		WidthCm:       req.WidthCm,
		LengthCm:      req.LengthCm,
		MaxSpectators: req.MaxSpectators,
		IsOutdoor:     req.IsOutdoor,
		Coating:       req.Coating,
	})
}

func (s *Service) UpdateArenaByID(ctx context.Context, req dto.UpdateArenaByIDRequest) error {
	return s.repo.UpdateArenaByID(ctx, repodto.UpdateArenaByIDRequest{
		ID:                req.ID,
		Name:              req.Name,
		Location:          req.Location,
		RefereesCount:     req.RefereesCount,
		TreadmillLengthCm: req.TreadmillLengthCm,
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
