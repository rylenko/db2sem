package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"db2sem/internal/db/pg"
	"db2sem/internal/domain"
	"db2sem/internal/repo/dto"
)

type Repo struct {
	conn pg.Conn
}

func New(conn pg.Conn) *Repo {
	return &Repo{conn: conn}
}

func (r *Repo) GetCourts(ctx context.Context, req dto.GetCourtsRequest) ([]domain.Court, error) {
	pgPlaces, err := r.conn.Queries(ctx).GetCourts(ctx, pg.GetCourtsParams{
		WidthCm:   convertToPgInt8(req.WidthCm),
		LengthCm:  convertToPgInt8(req.LengthCm),
		IsOutdoor: convertToPgBool(req.IsOutdoor),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	places := make([]domain.Court, 0, len(pgPlaces))

	for _, pgPlace := range pgPlaces {
		place := domain.Court{
			ID:        pgPlace.ID,
			Name:      pgPlace.Name,
			Location:  pgPlace.Location,
			WidthCm:   pgPlace.WidthCm,
			LengthCm:  pgPlace.LengthCm,
			IsOutdoor: pgPlace.IsOutdoor,
		}

		places = append(places, place)
	}

	return places, nil
}

func (r *Repo) GetGyms(ctx context.Context, req dto.GetGymsRequest) ([]domain.Gym, error) {
	pgPlaces, err := r.conn.Queries(ctx).GetGyms(ctx, pg.GetGymsParams{
		TrainersCount:  convertToPgInt2(req.TrainersCount),
		DumbbellsCount: convertToPgInt2(req.DumbbellsCount),
		HasBathhouse:   convertToPgBool(req.HasBathhouse),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	places := make([]domain.Gym, 0, len(pgPlaces))

	for _, pgPlace := range pgPlaces {
		place := domain.Gym{
			ID:             pgPlace.ID,
			Name:           pgPlace.Name,
			Location:       pgPlace.Location,
			TrainersCount:  pgPlace.TrainersCount,
			DumbbellsCount: pgPlace.DumbbellsCount,
			HasBathhouse:   pgPlace.HasBathhouse,
		}

		places = append(places, place)
	}

	return places, nil
}

func (r *Repo) GetStadiums(ctx context.Context, req dto.GetStadiumsRequest) ([]domain.Stadium, error) {
	var coating *string

	if req.Coating != nil {
		coatingValue := "%" + *req.Coating + "%"
		coating = &coatingValue
	}

	pgPlaces, err := r.conn.Queries(ctx).GetStadiums(ctx, pg.GetStadiumsParams{
		MaxSpectators: convertToPgInt2(req.MaxSpectators),
		WidthCm:       convertToPgInt8(req.WidthCm),
		LengthCm:      convertToPgInt8(req.LengthCm),
		IsOutdoor:     convertToPgBool(req.IsOutdoor),
		Coating:       convertToPgText(coating),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	places := make([]domain.Stadium, 0, len(pgPlaces))

	for _, pgPlace := range pgPlaces {
		place := domain.Stadium{
			ID:            pgPlace.ID,
			Name:          pgPlace.Name,
			Location:      pgPlace.Location,
			WidthCm:       pgPlace.WidthCm,
			LengthCm:      pgPlace.LengthCm,
			MaxSpectators: pgPlace.MaxSpectators,
			IsOutdoor:     pgPlace.IsOutdoor,
			Coating:       pgPlace.Coating,
		}

		places = append(places, place)
	}

	return places, nil
}

func (r *Repo) GetArenaByID(ctx context.Context, id int64) (*domain.Arena, error) {
	pgPlace, err := r.conn.Queries(ctx).GetArenaByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil //nolint:nilnil // because
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	return &domain.Arena{
		ID:                pgPlace.ID,
		Name:              pgPlace.Name,
		Location:          pgPlace.Location,
		RefereesCount:     pgPlace.RefereesCount,
		TreadmillLengthCm: pgPlace.TreadmillLengthCm,
	}, nil
}

func (r *Repo) GetArenas(ctx context.Context, req dto.GetArenasRequest) ([]domain.Arena, error) {
	pgPlaces, err := r.conn.Queries(ctx).GetArenas(ctx, pg.GetArenasParams{
		RefereesCount:     convertToPgInt2(req.RefereesCount),
		TreadmillLengthCm: convertToPgInt8(req.TreadmillLengthCm),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	places := make([]domain.Arena, 0, len(pgPlaces))

	for _, pgPlace := range pgPlaces {
		place := domain.Arena{
			ID:                pgPlace.ID,
			Name:              pgPlace.Name,
			Location:          pgPlace.Location,
			RefereesCount:     pgPlace.RefereesCount,
			TreadmillLengthCm: pgPlace.TreadmillLengthCm,
		}

		places = append(places, place)
	}

	return places, nil
}

func (r *Repo) DeleteSportByID(ctx context.Context, sportID int64) error {
	return r.conn.Queries(ctx).DeleteSportByID(ctx, sportID)
}

func (r *Repo) DeleteSportsmanByID(ctx context.Context, sportsmanID int64) error {
	return r.conn.Queries(ctx).DeleteSportsmanByID(ctx, sportsmanID)
}

func (r *Repo) GetClubActiveSportsmenCountsForPeriod(
	ctx context.Context,
	startAt time.Time,
	endAt time.Time,
) ([]domain.ClubSportsmenCount, error) {
	pgClubs, err := r.conn.Queries(ctx).GetClubActiveSportsmenCountsForPeriod(
		ctx,
		pg.GetClubActiveSportsmenCountsForPeriodParams{
			StartAt: convertToPgTimestamptz(startAt),
			EndAt:   convertToPgTimestamptz(endAt),
		},
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	clubs := make([]domain.ClubSportsmenCount, 0, len(pgClubs))

	for _, pgClub := range pgClubs {
		club := domain.Club{
			ID:   pgClub.ID,
			Name: pgClub.Name,
		}

		clubs = append(clubs, domain.ClubSportsmenCount{
			Club:           club,
			SportsmenCount: uint64(pgClub.ActiveSportsmenCount),
		})
	}

	return clubs, nil
}

func (r *Repo) GetClubs(ctx context.Context) ([]domain.Club, error) {
	pgClubs, err := r.conn.Queries(ctx).GetClubs(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	clubs := make([]domain.Club, 0, len(pgClubs))

	for _, pgClub := range pgClubs {
		club := domain.Club{
			ID:   pgClub.ID,
			Name: pgClub.Name,
		}

		clubs = append(clubs, club)
	}

	return clubs, nil
}

func (r *Repo) GetPrizeWinnersByTournamentID(ctx context.Context, tournamentID int64) ([]domain.PrizeWinner, error) {
	pgWinners, err := r.conn.Queries(ctx).GetPrizeWinnersByTournamentID(ctx, tournamentID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("get sportsman: %w", err)
	}

	winners := make([]domain.PrizeWinner, 0, len(pgWinners))

	for _, pgWinner := range pgWinners {
		birthDate, err := convertFromPgDate(pgWinner.BirthDate)
		if err != nil {
			return nil, fmt.Errorf("birth date: %w", err)
		}

		weightKg, err := convertFromPgNumeric(pgWinner.WeightKg)
		if err != nil {
			return nil, fmt.Errorf("weight kg: %w", err)
		}

		pgSports, err := r.conn.Queries(ctx).GetSportsBySportsmanID(ctx, pgWinner.ID)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("get sports: %w", err)
		}

		sports := make([]domain.Sport, 0, len(pgSports))

		for _, pgSport := range pgSports {
			sport := domain.Sport{
				ID:   pgSport.ID,
				Name: pgSport.Name,
			}

			sports = append(sports, sport)
		}

		sportsman := domain.Sportsman{
			ID:        pgWinner.ID,
			Name:      pgWinner.Name,
			BirthDate: birthDate,
			HeightCm:  uint16(pgWinner.HeightCm),
			WeightKg:  weightKg,
			Club: domain.Club{
				ID:   pgWinner.ClubID,
				Name: pgWinner.ClubName,
			},
			Sports: sports,
		}

		winner := domain.PrizeWinner{
			Sportsman: sportsman,
			Rank:      pgWinner.Rank,
		}

		winners = append(winners, winner)
	}

	return winners, nil
}

func (r *Repo) GetSportsmanByID(ctx context.Context, sportsmanID int64) (*domain.Sportsman, error) {
	pgSportsman, err := r.conn.Queries(ctx).GetSportsmanByID(ctx, sportsmanID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil //nolint:nilnil // must be checked on the top level
		}

		return nil, fmt.Errorf("get sportsman: %w", err)
	}

	birthDate, err := convertFromPgDate(pgSportsman.BirthDate)
	if err != nil {
		return nil, fmt.Errorf("birth date: %w", err)
	}

	weightKg, err := convertFromPgNumeric(pgSportsman.WeightKg)
	if err != nil {
		return nil, fmt.Errorf("weight kg: %w", err)
	}

	pgSports, err := r.conn.Queries(ctx).GetSportsBySportsmanID(ctx, sportsmanID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("get sports: %w", err)
	}

	sports := make([]domain.Sport, 0, len(pgSports))

	for _, pgSport := range pgSports {
		sport := domain.Sport{
			ID:   pgSport.ID,
			Name: pgSport.Name,
		}

		sports = append(sports, sport)
	}

	sportsman := domain.Sportsman{
		ID:        pgSportsman.ID,
		Name:      pgSportsman.Name,
		BirthDate: birthDate,
		HeightCm:  uint16(pgSportsman.HeightCm),
		WeightKg:  weightKg,
		Club: domain.Club{
			ID:   pgSportsman.ClubID,
			Name: pgSportsman.ClubName,
		},
		Sports: sports,
	}

	return &sportsman, nil
}

func (r *Repo) GetTrainersBySportsmanID(ctx context.Context, sportsmanID int64) ([]domain.Trainer, error) {
	pgTrainers, err := r.conn.Queries(ctx).GetTrainersBySportsmanID(ctx, sportsmanID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	trainers := make([]domain.Trainer, 0, len(pgTrainers))

	for _, pgTrainer := range pgTrainers {
		trainer := domain.Trainer{
			ID:   pgTrainer.ID,
			Name: pgTrainer.Name,
		}

		trainers = append(trainers, trainer)
	}

	return trainers, nil
}

func (r *Repo) GetTrainersBySportID(ctx context.Context, sportID int64) ([]domain.Trainer, error) {
	pgTrainers, err := r.conn.Queries(ctx).GetTrainersBySportID(ctx, sportID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	trainers := make([]domain.Trainer, 0, len(pgTrainers))

	for _, pgTrainer := range pgTrainers {
		trainer := domain.Trainer{
			ID:   pgTrainer.ID,
			Name: pgTrainer.Name,
		}

		trainers = append(trainers, trainer)
	}

	return trainers, nil
}

func (r *Repo) GetInactiveSportsmenForPeriod(
	ctx context.Context, startAt, endAt time.Time) ([]domain.Sportsman, error) {
	pgSportsmen, err := r.conn.Queries(ctx).GetInactiveSportsmenForPeriod(ctx, pg.GetInactiveSportsmenForPeriodParams{
		StartAt: convertToPgTimestamptz(startAt),
		EndAt:   convertToPgTimestamptz(endAt),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	sportsmen := make([]domain.Sportsman, 0, len(pgSportsmen))

	for _, pgSportsman := range pgSportsmen {
		birthDate, err := convertFromPgDate(pgSportsman.BirthDate)
		if err != nil {
			return nil, fmt.Errorf("birth date: %w", err)
		}

		weightKg, err := convertFromPgNumeric(pgSportsman.WeightKg)
		if err != nil {
			return nil, fmt.Errorf("weight kg: %w", err)
		}

		pgSports, err := r.conn.Queries(ctx).GetSportsBySportsmanID(ctx, pgSportsman.ID)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("get sports: %w", err)
		}

		sports := make([]domain.Sport, 0, len(pgSports))

		for _, pgSport := range pgSports {
			sport := domain.Sport{
				ID:   pgSport.ID,
				Name: pgSport.Name,
			}

			sports = append(sports, sport)
		}

		sportsman := domain.Sportsman{
			ID:        pgSportsman.ID,
			Name:      pgSportsman.Name,
			BirthDate: birthDate,
			HeightCm:  uint16(pgSportsman.HeightCm),
			WeightKg:  weightKg,
			Club: domain.Club{
				ID:   pgSportsman.ClubID,
				Name: pgSportsman.ClubName,
			},
			Sports: sports,
		}

		sportsmen = append(sportsmen, sportsman)
	}

	return sportsmen, nil
}

func (r *Repo) GetSportsmen(ctx context.Context) ([]domain.Sportsman, error) {
	pgSportsmen, err := r.conn.Queries(ctx).GetSportsmen(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	sportsmen := make([]domain.Sportsman, 0, len(pgSportsmen))

	for _, pgSportsman := range pgSportsmen {
		birthDate, err := convertFromPgDate(pgSportsman.BirthDate)
		if err != nil {
			return nil, fmt.Errorf("birth date: %w", err)
		}

		weightKg, err := convertFromPgNumeric(pgSportsman.WeightKg)
		if err != nil {
			return nil, fmt.Errorf("weight kg: %w", err)
		}

		pgSports, err := r.conn.Queries(ctx).GetSportsBySportsmanID(ctx, pgSportsman.ID)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("get sports: %w", err)
		}

		sports := make([]domain.Sport, 0, len(pgSports))

		for _, pgSport := range pgSports {
			sport := domain.Sport{
				ID:   pgSport.ID,
				Name: pgSport.Name,
			}

			sports = append(sports, sport)
		}

		sportsman := domain.Sportsman{
			ID:        pgSportsman.ID,
			Name:      pgSportsman.Name,
			BirthDate: birthDate,
			HeightCm:  uint16(pgSportsman.HeightCm),
			WeightKg:  weightKg,
			Club: domain.Club{
				ID:   pgSportsman.ClubID,
				Name: pgSportsman.ClubName,
			},
			Sports: sports,
		}

		sportsmen = append(sportsmen, sportsman)
	}

	return sportsmen, nil
}

func (r *Repo) GetSportsmenBySportID(
	ctx context.Context,
	sportID int64,
	minRank *int16,
) ([]domain.RankedSportsman, error) {
	pgSportsmen, err := r.conn.Queries(ctx).GetSportsmenBySportID(ctx, pg.GetSportsmenBySportIDParams{
		SportID: sportID,
		MinRank: convertToPgInt2(minRank),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("get sportsman: %w", err)
	}

	sportsmen := make([]domain.RankedSportsman, 0, len(pgSportsmen))

	for _, pgSportsman := range pgSportsmen {
		birthDate, err := convertFromPgDate(pgSportsman.BirthDate)
		if err != nil {
			return nil, fmt.Errorf("birth date: %w", err)
		}

		weightKg, err := convertFromPgNumeric(pgSportsman.WeightKg)
		if err != nil {
			return nil, fmt.Errorf("weight kg: %w", err)
		}

		pgSports, err := r.conn.Queries(ctx).GetSportsBySportsmanID(ctx, pgSportsman.ID)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("get sports: %w", err)
		}

		sports := make([]domain.Sport, 0, len(pgSports))

		for _, pgSport := range pgSports {
			sport := domain.Sport{
				ID:   pgSport.ID,
				Name: pgSport.Name,
			}

			sports = append(sports, sport)
		}

		sportsman := domain.Sportsman{
			ID:        pgSportsman.ID,
			Name:      pgSportsman.Name,
			BirthDate: birthDate,
			HeightCm:  uint16(pgSportsman.HeightCm),
			WeightKg:  weightKg,
			Club: domain.Club{
				ID:   pgSportsman.ClubID,
				Name: pgSportsman.ClubName,
			},
			Sports: sports,
		}

		rankedSportsman := domain.RankedSportsman{
			Sportsman: sportsman,
			Rank:      convertFromPgInt2(pgSportsman.Rank),
		}

		sportsmen = append(sportsmen, rankedSportsman)
	}

	return sportsmen, nil
}

func (r *Repo) GetSportsmenByTrainerID(
	ctx context.Context,
	trainerID int64,
	minRank *int16,
) ([]domain.RankedSportsman, error) {
	pgSportsmen, err := r.conn.Queries(ctx).GetSportsmenByTrainerID(ctx, pg.GetSportsmenByTrainerIDParams{
		TrainerID: trainerID,
		MinRank:   convertToPgInt2(minRank),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("get sportsman: %w", err)
	}

	sportsmen := make([]domain.RankedSportsman, 0, len(pgSportsmen))

	for _, pgSportsman := range pgSportsmen {
		birthDate, err := convertFromPgDate(pgSportsman.BirthDate)
		if err != nil {
			return nil, fmt.Errorf("birth date: %w", err)
		}

		weightKg, err := convertFromPgNumeric(pgSportsman.WeightKg)
		if err != nil {
			return nil, fmt.Errorf("weight kg: %w", err)
		}

		pgSports, err := r.conn.Queries(ctx).GetSportsBySportsmanID(ctx, pgSportsman.ID)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("get sports: %w", err)
		}

		sports := make([]domain.Sport, 0, len(pgSports))

		for _, pgSport := range pgSports {
			sport := domain.Sport{
				ID:   pgSport.ID,
				Name: pgSport.Name,
			}

			sports = append(sports, sport)
		}

		sportsman := domain.Sportsman{
			ID:        pgSportsman.ID,
			Name:      pgSportsman.Name,
			BirthDate: birthDate,
			HeightCm:  uint16(pgSportsman.HeightCm),
			WeightKg:  weightKg,
			Club: domain.Club{
				ID:   pgSportsman.ClubID,
				Name: pgSportsman.ClubName,
			},
			Sports: sports,
		}

		rankedSportsman := domain.RankedSportsman{
			Sportsman: sportsman,
			Rank:      convertFromPgInt2(pgSportsman.Rank),
		}

		sportsmen = append(sportsmen, rankedSportsman)
	}

	return sportsmen, nil
}

func (r *Repo) GetSportsmenInvolvedInSeveralSports(ctx context.Context) ([]domain.Sportsman, error) {
	pgSportsmen, err := r.conn.Queries(ctx).GetSportsmenInvolvedInSeveralSports(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	sportsmen := make([]domain.Sportsman, 0, len(pgSportsmen))

	for _, pgSportsman := range pgSportsmen {
		birthDate, err := convertFromPgDate(pgSportsman.BirthDate)
		if err != nil {
			return nil, fmt.Errorf("birth date: %w", err)
		}

		weightKg, err := convertFromPgNumeric(pgSportsman.WeightKg)
		if err != nil {
			return nil, fmt.Errorf("weight kg: %w", err)
		}

		pgSports, err := r.conn.Queries(ctx).GetSportsBySportsmanID(ctx, pgSportsman.ID)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("get sports: %w", err)
		}

		sports := make([]domain.Sport, 0, len(pgSports))

		for _, pgSport := range pgSports {
			sport := domain.Sport{
				ID:   pgSport.ID,
				Name: pgSport.Name,
			}

			sports = append(sports, sport)
		}

		sportsman := domain.Sportsman{
			ID:        pgSportsman.ID,
			Name:      pgSportsman.Name,
			BirthDate: birthDate,
			HeightCm:  uint16(pgSportsman.HeightCm),
			WeightKg:  weightKg,
			Club: domain.Club{
				ID:   pgSportsman.ClubID,
				Name: pgSportsman.ClubName,
			},
			Sports: sports,
		}

		sportsmen = append(sportsmen, sportsman)
	}

	return sportsmen, nil
}

func (r *Repo) GetSportByID(ctx context.Context, sportID int64) (*domain.Sport, error) {
	pgSport, err := r.conn.Queries(ctx).GetSportByID(ctx, sportID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil //nolint:nilnil // must be checked on the top level
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	sport := domain.Sport{
		ID:   pgSport.ID,
		Name: pgSport.Name,
	}

	return &sport, nil
}

func (r *Repo) GetSports(ctx context.Context) ([]domain.Sport, error) {
	pgSports, err := r.conn.Queries(ctx).GetSports(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	sports := make([]domain.Sport, 0, len(pgSports))

	for _, pgSport := range pgSports {
		sport := domain.Sport{
			ID:   pgSport.ID,
			Name: pgSport.Name,
		}

		sports = append(sports, sport)
	}

	return sports, nil
}

func (r *Repo) GetTournaments(ctx context.Context) ([]domain.Tournament, error) {
	pgTournaments, err := r.conn.Queries(ctx).GetTournaments(ctx)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	tournaments := make([]domain.Tournament, 0, len(pgTournaments))

	for _, pgTournament := range pgTournaments {
		startAt, err := convertFromPgTimestamptz(pgTournament.StartAt)
		if err != nil {
			return nil, fmt.Errorf("convert start at: %w", err)
		}

		tournament := domain.Tournament{
			ID:            pgTournament.ID,
			OrganizerName: pgTournament.OrganizerName,
			PlaceName:     pgTournament.PlaceName,
			StartAt:       startAt,
		}

		tournaments = append(tournaments, tournament)
	}

	return tournaments, nil
}

func (r *Repo) GetTournamentsByPlaceID(
	ctx context.Context,
	placeID int64,
	sportID *int64,
) ([]domain.Tournament, error) {
	pgTournaments, err := r.conn.Queries(ctx).GetTournamentsByPlaceID(ctx, pg.GetTournamentsByPlaceIDParams{
		PlaceID: placeID,
		SportID: convertToPgInt8(sportID),
	})
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	tournaments := make([]domain.Tournament, 0, len(pgTournaments))

	for _, pgTournament := range pgTournaments {
		startAt, err := convertFromPgTimestamptz(pgTournament.StartAt)
		if err != nil {
			return nil, fmt.Errorf("convert start at: %w", err)
		}

		tournament := domain.Tournament{
			ID:            pgTournament.ID,
			OrganizerName: pgTournament.OrganizerName,
			PlaceName:     pgTournament.PlaceName,
			SportNames:    pgTournament.SportNames,
			StartAt:       startAt,
		}

		tournaments = append(tournaments, tournament)
	}

	return tournaments, nil
}

func (r *Repo) GetTournamentsForPeriod(
	ctx context.Context,
	startAt time.Time,
	endAt time.Time,
	organizerID *int64,
) ([]domain.Tournament, error) {
	pgTournaments, err := r.conn.Queries(ctx).GetTournamentsForPeriod(ctx, pg.GetTournamentsForPeriodParams{
		StartAt:     convertToPgTimestamptz(startAt),
		EndAt:       convertToPgTimestamptz(endAt),
		OrganizerID: convertToPgInt8(organizerID),
	})
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	tournaments := make([]domain.Tournament, 0, len(pgTournaments))

	for _, pgTournament := range pgTournaments {
		startAt, err := convertFromPgTimestamptz(pgTournament.StartAt)
		if err != nil {
			return nil, fmt.Errorf("convert start at: %w", err)
		}

		tournament := domain.Tournament{
			ID:            pgTournament.ID,
			OrganizerName: pgTournament.OrganizerName,
			PlaceName:     pgTournament.PlaceName,
			SportNames:    pgTournament.SportNames,
			StartAt:       startAt,
		}

		tournaments = append(tournaments, tournament)
	}

	return tournaments, nil
}

func (r *Repo) InsertSport(ctx context.Context, name string) error {
	return r.conn.Queries(ctx).InsertSport(ctx, name)
}

func (r *Repo) InsertSportsman(ctx context.Context, req dto.InsertSportsmanRequest) error {
	weightKg, err := convertToPgNumeric(req.WeightKg)
	if err != nil {
		return fmt.Errorf("convert weight kg: %w", err)
	}

	err = r.conn.Queries(ctx).InsertSportsman(ctx, pg.InsertSportsmanParams{
		Name:      req.Name,
		BirthDate: convertToPgDate(req.BirthDate),
		HeightCm:  int16(req.HeightCm),
		WeightKg:  weightKg,
		ClubID:    req.ClubID,
		SportIds:  req.SportIDs,
	})
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	return nil
}

func (r *Repo) UpdateArenaByID(ctx context.Context, req dto.UpdateArenaByIDRequest) error {
	return r.conn.Queries(ctx).UpdateArenaByID(ctx, pg.UpdateArenaByIDParams{
		ID:                req.ID,
		Name:              req.Name,
		Location:          req.Location,
		TreadmillLengthCm: req.TreadmillLengthCm,
		RefereesCount:     req.RefereesCount,
	})
}

func (r *Repo) UpdateSportByID(ctx context.Context, req dto.UpdateSportByIDRequest) error {
	return r.conn.Queries(ctx).UpdateSportByID(ctx, pg.UpdateSportByIDParams{
		ID:   req.ID,
		Name: req.Name,
	})
}

func (r *Repo) UpdateSportsmanByID(ctx context.Context, req dto.UpdateSportsmanByIDRequest) error {
	weightKg, err := convertToPgNumeric(req.WeightKg)
	if err != nil {
		return fmt.Errorf("convert weight kg: %w", err)
	}

	err = r.conn.Queries(ctx).UpdateSportsmanByID(ctx, pg.UpdateSportsmanByIDParams{
		ID:        req.ID,
		Name:      req.Name,
		BirthDate: convertToPgDate(req.BirthDate),
		HeightCm:  int16(req.HeightCm),
		WeightKg:  weightKg,
		ClubID:    req.ClubID,
		SportIds:  req.SportIDs,
	})
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	return nil
}

func (r *Repo) GetTrainers(ctx context.Context) ([]domain.Trainer, error) {
	pgTrainers, err := r.conn.Queries(ctx).GetTrainers(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	trainers := make([]domain.Trainer, 0, len(pgTrainers))

	for _, pgTrainer := range pgTrainers {
		trainer := domain.Trainer{
			ID:   pgTrainer.ID,
			Name: pgTrainer.Name,
		}

		trainers = append(trainers, trainer)
	}

	return trainers, nil
}

func (r *Repo) GetOrganizerTournamentCountsForPeriod(
	ctx context.Context,
	startAt time.Time,
	endAt time.Time,
) ([]domain.OrganizerTournamentsCount, error) {
	pgOrganizers, err := r.conn.Queries(ctx).GetOrganizerTournamentCountsForPeriod(
		ctx,
		pg.GetOrganizerTournamentCountsForPeriodParams{
			StartAt: convertToPgTimestamptz(startAt),
			EndAt:   convertToPgTimestamptz(endAt),
		},
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	organizers := make([]domain.OrganizerTournamentsCount, 0, len(pgOrganizers))

	for _, pgOrganizer := range pgOrganizers {
		organizer := domain.Organizer{
			ID:       pgOrganizer.ID,
			Name:     pgOrganizer.Name,
			Location: convertFromPgText(pgOrganizer.Location),
		}

		organizers = append(organizers, domain.OrganizerTournamentsCount{
			Organizer:        organizer,
			TournamentsCount: uint64(pgOrganizer.TournamentsCount),
		})
	}

	return organizers, nil
}

func (r *Repo) GetOrganizers(ctx context.Context) ([]domain.Organizer, error) {
	pgOrganizers, err := r.conn.Queries(ctx).GetOrganizers(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	organizers := make([]domain.Organizer, 0, len(pgOrganizers))

	for _, pgOrganizer := range pgOrganizers {
		organizer := domain.Organizer{
			ID:       pgOrganizer.ID,
			Name:     pgOrganizer.Name,
			Location: convertFromPgText(pgOrganizer.Location),
		}

		organizers = append(organizers, organizer)
	}

	return organizers, nil
}

func (r *Repo) GetPlaces(ctx context.Context) ([]domain.Place, error) {
	pgPlaces, err := r.conn.Queries(ctx).GetPlaces(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	places := make([]domain.Place, 0, len(pgPlaces))

	for _, pgPlace := range pgPlaces {
		place := domain.Place{
			ID:       pgPlace.ID,
			Name:     pgPlace.Name,
			Location: pgPlace.Location,
			TypeName: pgPlace.TypeName,
		}

		places = append(places, place)
	}

	return places, nil
}

func (r *Repo) GetPlacesWithTournamentDatesForPeriod(
	ctx context.Context, startAt, endAt time.Time) ([]domain.PlaceWithTournamentDates, error) {
	pgPlaces, err := r.conn.Queries(ctx).GetPlacesWithTournamentDatesForPeriod(
		ctx,
		pg.GetPlacesWithTournamentDatesForPeriodParams{
			StartAt: convertToPgTimestamptz(startAt),
			EndAt:   convertToPgTimestamptz(endAt),
		},
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	places := make([]domain.PlaceWithTournamentDates, 0, len(pgPlaces))

	for _, pgPlace := range pgPlaces {
		place := domain.Place{
			ID:       pgPlace.ID,
			Name:     pgPlace.Name,
			Location: pgPlace.Location,
			TypeName: pgPlace.TypeName,
		}

		dates := make([]time.Time, 0, len(pgPlace.TournamentDates))

		for _, pgDate := range pgPlace.TournamentDates {
			date, err := convertFromPgTimestamptz(pgDate)
			if err != nil {
				return nil, err
			}

			dates = append(dates, date)
		}

		places = append(places, domain.PlaceWithTournamentDates{
			Place:           place,
			TournamentDates: dates,
		})
	}

	return places, nil
}

func (r *Repo) DeletePlaceByID(ctx context.Context, placeID int64) error {
	return r.conn.Queries(ctx).DeletePlaceByID(ctx, placeID)
}
