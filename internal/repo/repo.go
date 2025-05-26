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

func (r *Repo) DeleteSportByID(ctx context.Context, sportID int64) error {
	return r.conn.Queries(ctx).DeleteSportByID(ctx, sportID)
}

func (r *Repo) DeleteSportsmanByID(ctx context.Context, sportsmanID int64) error {
	return r.conn.Queries(ctx).DeleteSportsmanByID(ctx, sportsmanID)
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
