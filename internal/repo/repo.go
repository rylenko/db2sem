package repo

import (
	"context"
	"errors"
	"fmt"

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

func (r *Repo) DeleteSportsmanByID(ctx context.Context, sportsmanID int64) error {
	return r.conn.Queries(ctx).DeleteSportsmanByID(ctx, sportsmanID)
}

func (r *Repo) GetSportsmanByID(ctx context.Context, sportsmanID int64) (*domain.Sportsman, error) {
	pgSportsman, err := r.conn.Queries(ctx).GetSportsmanByID(ctx, sportsmanID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil //nolint:nilnil // must be checked on the top level
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	birthDate, err := convertFromPgDate(pgSportsman.BirthDate)
	if err != nil {
		return nil, fmt.Errorf("birth date: %w", err)
	}

	weightKg, err := convertFromPgNumeric(pgSportsman.WeightKg)
	if err != nil {
		return nil, fmt.Errorf("weight kg: %w", err)
	}

	sportsman := domain.Sportsman{
		ID:         pgSportsman.ID,
		Name:       pgSportsman.Name,
		BirthDate:  birthDate,
		HeightCm:   uint16(pgSportsman.HeightCm),
		WeightKg:   weightKg,
		SportNames: pgSportsman.SportNames,
	}

	return &sportsman, nil
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

		sportsman := domain.Sportsman{
			ID:         pgSportsman.ID,
			Name:       pgSportsman.Name,
			BirthDate:  birthDate,
			HeightCm:   uint16(pgSportsman.HeightCm),
			WeightKg:   weightKg,
			SportNames: pgSportsman.SportNames,
		}

		sportsmen = append(sportsmen, sportsman)
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

		sportsman := domain.Sportsman{
			ID:         pgSportsman.ID,
			Name:       pgSportsman.Name,
			BirthDate:  birthDate,
			HeightCm:   uint16(pgSportsman.HeightCm),
			WeightKg:   weightKg,
			SportNames: pgSportsman.SportNames,
		}

		sportsmen = append(sportsmen, sportsman)
	}

	return sportsmen, nil
}

func (r *Repo) GetSportNames(ctx context.Context) ([]string, error) {
	sportNames, err := r.conn.Queries(ctx).GetSportNames(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	return sportNames, nil
}

func (r *Repo) UpdateSportsmanByID(ctx context.Context, req dto.UpdateSportsmanByIDRequest) error {
	weightKg, err := convertToPgNumeric(req.WeightKg)
	if err != nil {
		return fmt.Errorf("convert weight kg: %w", err)
	}

	err = r.conn.Queries(ctx).UpdateSportsmanByID(ctx, pg.UpdateSportsmanByIDParams{
		ID:         req.ID,
		Name:       req.Name,
		BirthDate:  convertToPgDate(req.BirthDate),
		HeightCm:   int16(req.HeightCm),
		WeightKg:   weightKg,
		SportNames: req.SportNames,
	})
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	return nil
}
