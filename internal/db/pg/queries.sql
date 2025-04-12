-- Query #1.1
-- name: GetArenaPlaces :many
SELECT p.*
FROM places p
JOIN place_types pt ON pt.id = p.type_id
JOIN arena_attributes aa ON aa.place_id = p.id
WHERE
	pt.attributes_table_name = 'arena_attributes'
	AND (aa.referees_count >= sqlc.narg('referees_count') OR sqlc.narg('referees_count') IS NULL)
	AND (aa.treadmill_length_cm >= sqlc.narg('treadmill_length_cm') OR sqlc.narg('treadmill_length_cm') IS NULL);

-- Query #1.2
-- name: GetStadiumPlaces :many
SELECT p.*
FROM places p
JOIN place_types pt ON pt.id = p.type_id
JOIN stadium_attributes sa ON sa.place_id = p.id
WHERE
	pt.attributes_table_name = 'stadium_attributes'
	AND (sa.width_cm >= sqlc.narg('width_cm') OR sqlc.narg('width_cm') IS NULL)
	AND (sa.length_cm >= sqlc.narg('length_cm') OR sqlc.narg('length_cm') IS NULL)
	AND (sa.max_spectators >= sqlc.narg('max_spectators') OR sqlc.narg('max_spectators') IS NULL)
	AND (sa.is_outdoor = sqlc.narg('is_outdoor') OR sqlc.narg('is_outdoor') IS NULL)
	AND (sa.coating = sqlc.narg('coating') OR sqlc.narg('coating') IS NULL);

-- Query #2
-- name: GetSportsmenBySportID :many
SELECT sm.*
FROM sportsmen sm
JOIN sportsman_sports ss ON ss.sportsman_id = sm.id
WHERE
	ss.sport_id = @sport_id
	AND (ss.rank = sqlc.narg('rank') OR sqlc.narg('rank') IS NULL);

-- Query #3
-- name: GetSportsmenByTrainerID :many
SELECT sm.*
FROM sportsmen sm
JOIN sportsman_sports ss ON ss.sportsman_id = sm.id
JOIN sportsman_sport_trainers sst ON sst.sportsman_sport_id = ss.id
WHERE
	sst.trainer_id = @trainer_id
	AND (ss.rank = sqlc.narg('rank') OR sqlc.narg('rank') IS NULL);

-- Query #4
-- name: GetSportsmanIDsInvolvedInSeveralSports :many
SELECT
	sportsman_id,
	ARRAY_AGG(sport_id)::BIGINT[] AS sport_ids
FROM sportsman_sports
GROUP BY sportsman_id
HAVING COUNT(*) > 1;

-- Query #5
-- name: GetTrainersBySportsmanID :many
SELECT t.*
FROM trainers t
JOIN sportsman_sport_trainers sst ON sst.trainer_id = t.id
JOIN sportsman_sports ss ON ss.id = sst.sportsman_sport_id
WHERE ss.sportsman_id = $1;

-- Query #6
-- name: GetTournamentsForPeriod :many
SELECT *
FROM tournaments
WHERE
	start_at BETWEEN @start_at AND @end_at
	AND (organizer_id = sqlc.narg('organizer_id') OR sqlc.narg('organizer_id') IS NULL);

-- Query #7
-- name: GetSportsmenByTournamentID :many
SELECT sm.*
FROM sportsmen sm
JOIN participations p ON p.sportsman_id = sm.id
JOIN tournament_sports ts ON ts.id = p.tournament_sport_id
WHERE ts.tournament_id = $1;

-- Query #8
-- name: GetTournamentsByPlaceID :many
SELECT t.*
FROM tournaments t
JOIN tournament_sports ts ON ts.tournament_id = t.id
WHERE
	t.place_id = $1
	AND (ts.sport_id = sqlc.narg('sport_id') OR sqlc.narg('sport_id') IS NULL);

-- Query #9
-- name: GetClubCompetitorCountsForPeriod :many
SELECT
	c.*,
	COUNT(s.id)
FROM clubs c
JOIN sportsmen s ON s.club_id = c.id
JOIN participations p ON p.sportsman_id = s.id
JOIN tournament_sports ts ON ts.id = p.tournament_sport_id
JOIN tournaments t ON t.id = ts.tournament_id
WHERE t.start_at BETWEEN @start_at AND @end_at
GROUP BY c.id;

-- Query #10
-- name: GetTrainersBySportID :many
SELECT t.*
FROM trainers t
JOIN sportsman_sport_trainers sst ON sst.trainer_id = t.id
JOIN sportsman_sports ss ON ss.id = sst.sportsman_sport_id
WHERE ss.sport_id = $1;
