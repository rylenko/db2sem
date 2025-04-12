// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: queries.sql

package pg

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getArenaPlaces = `-- name: GetArenaPlaces :many
SELECT p.id, p.name, p.location, p.type_id, p.created_at
FROM places p
JOIN place_types pt ON pt.id = p.type_id
JOIN arena_attributes aa ON aa.place_id = p.id
WHERE
	pt.attributes_table_name = 'arena_attributes'
	AND (aa.referees_count >= $1 OR $1 IS NULL)
	AND (aa.treadmill_length_cm >= $2 OR $2 IS NULL)
`

type GetArenaPlacesParams struct {
	RefereesCount     pgtype.Int2
	TreadmillLengthCm pgtype.Int8
}

// Query #1.1
func (q *Queries) GetArenaPlaces(ctx context.Context, arg GetArenaPlacesParams) ([]Place, error) {
	rows, err := q.db.Query(ctx, getArenaPlaces, arg.RefereesCount, arg.TreadmillLengthCm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Place
	for rows.Next() {
		var i Place
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Location,
			&i.TypeID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getClubCompetitorCountsForPeriod = `-- name: GetClubCompetitorCountsForPeriod :many
SELECT
	c.id, c.name, c.created_at,
	COUNT(s.id)
FROM clubs c
JOIN sportsmen s ON s.club_id = c.id
JOIN participations p ON p.sportsman_id = s.id
JOIN tournament_sports ts ON ts.id = p.tournament_sport_id
JOIN tournaments t ON t.id = ts.tournament_id
WHERE t.start_at BETWEEN $1 AND $2
GROUP BY c.id
`

type GetClubCompetitorCountsForPeriodParams struct {
	StartAt pgtype.Timestamptz
	EndAt   pgtype.Timestamptz
}

type GetClubCompetitorCountsForPeriodRow struct {
	ID        int64
	Name      string
	CreatedAt pgtype.Timestamptz
	Count     int64
}

// Query #9
func (q *Queries) GetClubCompetitorCountsForPeriod(ctx context.Context, arg GetClubCompetitorCountsForPeriodParams) ([]GetClubCompetitorCountsForPeriodRow, error) {
	rows, err := q.db.Query(ctx, getClubCompetitorCountsForPeriod, arg.StartAt, arg.EndAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetClubCompetitorCountsForPeriodRow
	for rows.Next() {
		var i GetClubCompetitorCountsForPeriodRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.Count,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSportsmanIDsInvolvedInSeveralSports = `-- name: GetSportsmanIDsInvolvedInSeveralSports :many
SELECT
	sportsman_id,
	ARRAY_AGG(sport_id)::BIGINT[] AS sport_ids
FROM sportsman_sports
GROUP BY sportsman_id
HAVING COUNT(*) > 1
`

type GetSportsmanIDsInvolvedInSeveralSportsRow struct {
	SportsmanID int64
	SportIds    []int64
}

// Query #4
func (q *Queries) GetSportsmanIDsInvolvedInSeveralSports(ctx context.Context) ([]GetSportsmanIDsInvolvedInSeveralSportsRow, error) {
	rows, err := q.db.Query(ctx, getSportsmanIDsInvolvedInSeveralSports)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetSportsmanIDsInvolvedInSeveralSportsRow
	for rows.Next() {
		var i GetSportsmanIDsInvolvedInSeveralSportsRow
		if err := rows.Scan(&i.SportsmanID, &i.SportIds); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSportsmenBySportID = `-- name: GetSportsmenBySportID :many
SELECT sm.id, sm.name, sm.birth_date, sm.height_cm, sm.weight_kg, sm.club_id, sm.created_at
FROM sportsmen sm
JOIN sportsman_sports ss ON ss.sportsman_id = sm.id
WHERE
	ss.sport_id = $1
	AND (ss.rank = $2 OR $2 IS NULL)
`

type GetSportsmenBySportIDParams struct {
	SportID int64
	Rank    pgtype.Int2
}

// Query #2
func (q *Queries) GetSportsmenBySportID(ctx context.Context, arg GetSportsmenBySportIDParams) ([]Sportsman, error) {
	rows, err := q.db.Query(ctx, getSportsmenBySportID, arg.SportID, arg.Rank)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Sportsman
	for rows.Next() {
		var i Sportsman
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.BirthDate,
			&i.HeightCm,
			&i.WeightKg,
			&i.ClubID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSportsmenByTournamentID = `-- name: GetSportsmenByTournamentID :many
SELECT sm.id, sm.name, sm.birth_date, sm.height_cm, sm.weight_kg, sm.club_id, sm.created_at
FROM sportsmen sm
JOIN participations p ON p.sportsman_id = sm.id
JOIN tournament_sports ts ON ts.id = p.tournament_sport_id
WHERE ts.tournament_id = $1
`

// Query #7
func (q *Queries) GetSportsmenByTournamentID(ctx context.Context, tournamentID pgtype.Int8) ([]Sportsman, error) {
	rows, err := q.db.Query(ctx, getSportsmenByTournamentID, tournamentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Sportsman
	for rows.Next() {
		var i Sportsman
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.BirthDate,
			&i.HeightCm,
			&i.WeightKg,
			&i.ClubID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSportsmenByTrainerID = `-- name: GetSportsmenByTrainerID :many
SELECT sm.id, sm.name, sm.birth_date, sm.height_cm, sm.weight_kg, sm.club_id, sm.created_at
FROM sportsmen sm
JOIN sportsman_sports ss ON ss.sportsman_id = sm.id
JOIN sportsman_sport_trainers sst ON sst.sportsman_sport_id = ss.id
WHERE
	sst.trainer_id = $1
	AND (ss.rank = $2 OR $2 IS NULL)
`

type GetSportsmenByTrainerIDParams struct {
	TrainerID pgtype.Int8
	Rank      pgtype.Int2
}

// Query #3
func (q *Queries) GetSportsmenByTrainerID(ctx context.Context, arg GetSportsmenByTrainerIDParams) ([]Sportsman, error) {
	rows, err := q.db.Query(ctx, getSportsmenByTrainerID, arg.TrainerID, arg.Rank)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Sportsman
	for rows.Next() {
		var i Sportsman
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.BirthDate,
			&i.HeightCm,
			&i.WeightKg,
			&i.ClubID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getStadiumPlaces = `-- name: GetStadiumPlaces :many
SELECT p.id, p.name, p.location, p.type_id, p.created_at
FROM places p
JOIN place_types pt ON pt.id = p.type_id
JOIN stadium_attributes sa ON sa.place_id = p.id
WHERE
	pt.attributes_table_name = 'stadium_attributes'
	AND (sa.width_cm >= $1 OR $1 IS NULL)
	AND (sa.length_cm >= $2 OR $2 IS NULL)
	AND (sa.max_spectators >= $3 OR $3 IS NULL)
	AND (sa.is_outdoor = $4 OR $4 IS NULL)
	AND (sa.coating = $5 OR $5 IS NULL)
`

type GetStadiumPlacesParams struct {
	WidthCm       pgtype.Int8
	LengthCm      pgtype.Int8
	MaxSpectators pgtype.Int2
	IsOutdoor     pgtype.Bool
	Coating       pgtype.Text
}

// Query #1.2
func (q *Queries) GetStadiumPlaces(ctx context.Context, arg GetStadiumPlacesParams) ([]Place, error) {
	rows, err := q.db.Query(ctx, getStadiumPlaces,
		arg.WidthCm,
		arg.LengthCm,
		arg.MaxSpectators,
		arg.IsOutdoor,
		arg.Coating,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Place
	for rows.Next() {
		var i Place
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Location,
			&i.TypeID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTournamentsByPlaceID = `-- name: GetTournamentsByPlaceID :many
SELECT t.id, t.place_id, t.start_at, t.organizer_id, t.created_at
FROM tournaments t
JOIN tournament_sports ts ON ts.tournament_id = t.id
WHERE
	t.place_id = $1
	AND (ts.sport_id = $2 OR $2 IS NULL)
`

type GetTournamentsByPlaceIDParams struct {
	PlaceID int64
	SportID pgtype.Int8
}

// Query #8
func (q *Queries) GetTournamentsByPlaceID(ctx context.Context, arg GetTournamentsByPlaceIDParams) ([]Tournament, error) {
	rows, err := q.db.Query(ctx, getTournamentsByPlaceID, arg.PlaceID, arg.SportID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tournament
	for rows.Next() {
		var i Tournament
		if err := rows.Scan(
			&i.ID,
			&i.PlaceID,
			&i.StartAt,
			&i.OrganizerID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTournamentsForPeriod = `-- name: GetTournamentsForPeriod :many
SELECT id, place_id, start_at, organizer_id, created_at
FROM tournaments
WHERE
	start_at BETWEEN $1 AND $2
	AND (organizer_id = $3 OR $3 IS NULL)
`

type GetTournamentsForPeriodParams struct {
	StartAt     pgtype.Timestamptz
	EndAt       pgtype.Timestamptz
	OrganizerID pgtype.Int8
}

// Query #6
func (q *Queries) GetTournamentsForPeriod(ctx context.Context, arg GetTournamentsForPeriodParams) ([]Tournament, error) {
	rows, err := q.db.Query(ctx, getTournamentsForPeriod, arg.StartAt, arg.EndAt, arg.OrganizerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tournament
	for rows.Next() {
		var i Tournament
		if err := rows.Scan(
			&i.ID,
			&i.PlaceID,
			&i.StartAt,
			&i.OrganizerID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTrainersBySportID = `-- name: GetTrainersBySportID :many
SELECT t.id, t.name, t.created_at
FROM trainers t
JOIN sportsman_sport_trainers sst ON sst.trainer_id = t.id
JOIN sportsman_sports ss ON ss.id = sst.sportsman_sport_id
WHERE ss.sport_id = $1
`

// Query #10
func (q *Queries) GetTrainersBySportID(ctx context.Context, sportID int64) ([]Trainer, error) {
	rows, err := q.db.Query(ctx, getTrainersBySportID, sportID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Trainer
	for rows.Next() {
		var i Trainer
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTrainersBySportsmanID = `-- name: GetTrainersBySportsmanID :many
SELECT t.id, t.name, t.created_at
FROM trainers t
JOIN sportsman_sport_trainers sst ON sst.trainer_id = t.id
JOIN sportsman_sports ss ON ss.id = sst.sportsman_sport_id
WHERE ss.sportsman_id = $1
`

// Query #5
func (q *Queries) GetTrainersBySportsmanID(ctx context.Context, sportsmanID int64) ([]Trainer, error) {
	rows, err := q.db.Query(ctx, getTrainersBySportsmanID, sportsmanID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Trainer
	for rows.Next() {
		var i Trainer
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
