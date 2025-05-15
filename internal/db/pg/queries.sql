-- Query #1.1
--
-- Получить перечень спортивных сооружений указанного типа в целом или
-- удовлетворяющих заданным характеристикам (например, стадионы, вмещающие не менее
-- указанного числа зрителей).
--
-- name: GetArenaPlaces :many
SELECT
	p.name,
	p.location
FROM places p
JOIN arena_attributes aa ON aa.place_id = p.id
WHERE
	(
		aa.referees_count >= sqlc.narg('referees_count')
		OR sqlc.narg('referees_count') IS NULL
	)
	AND (
		aa.treadmill_length_cm >= sqlc.narg('treadmill_length_cm')
		OR sqlc.narg('treadmill_length_cm') IS NULL
	);

-- Query #1.2
--
-- Получить перечень спортивных сооружений указанного типа в целом или
-- удовлетворяющих заданным характеристикам (например, стадионы, вмещающие не менее
-- указанного числа зрителей).
--
-- name: GetStadiumPlaces :many
SELECT
	p.name,
	p.location
FROM places p
JOIN stadium_attributes sa ON sa.place_id = p.id
WHERE
	(
		sa.width_cm >= sqlc.narg('width_cm')
		OR sqlc.narg('width_cm') IS NULL
	)
	AND (
		sa.length_cm >= sqlc.narg('length_cm')
		OR sqlc.narg('length_cm') IS NULL
	)
	AND (
		sa.max_spectators >= sqlc.narg('max_spectators')
		OR sqlc.narg('max_spectators') IS NULL
	)
	AND (
		sa.is_outdoor = sqlc.narg('is_outdoor')
		OR sqlc.narg('is_outdoor') IS NULL
	)
	AND (
		sa.coating = sqlc.narg('coating')
		OR sqlc.narg('coating') IS NULL
	);

-- Query #1.3
--
-- Получить перечень спортивных сооружений указанного типа в целом или
-- удовлетворяющих заданным характеристикам (например, стадионы, вмещающие не менее
-- указанного числа зрителей).
--
-- name: GetCourtPlaces :many
SELECT
	p.name,
	p.location
FROM places p
JOIN court_attributes ca ON ca.place_id = p.id
WHERE
	(
		ca.width_cm >= sqlc.narg('width_cm')
		OR sqlc.narg('width_cm') IS NULL
	)
	AND (
		ca.length_cm >= sqlc.narg('length_cm')
		OR sqlc.narg('length_cm') IS NULL
	)
	AND (
		ca.is_outdoor = sqlc.narg('is_outdoor')
		OR sqlc.narg('is_outdoor') IS NULL
	);

-- Query #1.4
--
-- Получить перечень спортивных сооружений указанного типа в целом или
-- удовлетворяющих заданным характеристикам (например, стадионы, вмещающие не менее
-- указанного числа зрителей).
--
-- name: GetGymPlaces :many
SELECT
	p.name,
	p.location
FROM places p
JOIN gym_attributes ga ON ga.place_id = p.id
WHERE
	(
		ga.trainers_count >= sqlc.narg('trainers_count')
		OR sqlc.narg('trainers_count') IS NULL
	)
	AND (
		ga.dumbbells_count >= sqlc.narg('dumbbells_count')
		OR sqlc.narg('dumbbells_count') IS NULL
	)
	AND (
		ga.has_bathhouse = sqlc.narg('has_bathhouse')
		OR sqlc.narg('has_bathhouse') IS NULL
	);

-- Query #2
--
-- Получить список спортсменов, занимающихся указанным видом спорта в целом либо не
-- ниже определенного разряда.
--
-- name: GetSportsmenBySportID :many
SELECT
	sm.name,
	sm.birth_date,
	sm.height_cm,
	sm.weight_kg
FROM sportsmen sm
JOIN sportsman_sports ss ON ss.sportsman_id = sm.id
WHERE
	ss.sport_id = @sport_id
	AND (
		ss.rank >= sqlc.narg('rank')
		OR sqlc.narg('rank') IS NULL
	);

-- Query #3
--
-- Получить список спортсменов, тренирующихся у некого тренера в целом либо не ниже
-- определенного разряда.
--
-- name: GetSportsmenByTrainerID :many
SELECT
	sm.name,
	sm.birth_date,
	sm.height_cm,
	sm.weight_kg
FROM sportsmen sm
JOIN sportsman_sports ss ON ss.sportsman_id = sm.id
JOIN sportsman_sport_trainers sst ON sst.sportsman_sport_id = ss.id
WHERE
	sst.trainer_id = @trainer_id
	AND (
		ss.rank >= sqlc.narg('rank')
		OR sqlc.narg('rank') IS NULL
	);

-- Query #4
--
-- Получить список спортсменов, занимающихся более чем одним видом спорта с указанием
-- этих видов спорта.
--
-- name: GetSportsmenInvolvedInSeveralSports :many
SELECT
	sm.id,
	sm.name,
	sm.birth_date,
	sm.height_cm,
	sm.weight_kg,
	c.id AS club_id,
	c.name AS club_name
FROM sportsmen sm
JOIN clubs c ON c.id = sm.club_id
JOIN sportsman_sports sms ON sms.sportsman_id = sm.id
JOIN sports s ON s.id = sms.sport_id
GROUP BY
	sm.id,
	sm.name,
	sm.birth_date,
	sm.height_cm,
	sm.weight_kg,
	c.id,
	c.name
HAVING COUNT(sms.id) > 1
ORDER BY sm.id DESC;

-- Query #5
--
-- Получить список тренеров указанного спортсмена.
--
-- name: GetTrainersBySportsmanID :many
SELECT
	t.id,
	t.name
FROM trainers t
JOIN sportsman_sport_trainers sst ON sst.trainer_id = t.id
JOIN sportsman_sports ss ON ss.id = sst.sportsman_sport_id
WHERE ss.sportsman_id = $1;

-- Query #6
--
-- Получить перечень соревнований, проведенных в течение заданного периода времени в
-- целом либо указанным организатором.
--
-- name: GetTournamentsForPeriod :many
SELECT
	p.name,
	o.name,
	t.start_at
FROM tournaments t
JOIN places p ON p.id = t.place_id
JOIN organizers o ON o.id = t.organizer_id
WHERE
	t.start_at BETWEEN @start_at AND @end_at
	AND (
		t.organizer_id = sqlc.narg('organizer_id')
		OR sqlc.narg('organizer_id') IS NULL
	);

-- Query #7
--
-- Получить список призеров указанного соревнования.
--
-- name: GetPrizeWinnersByTournamentID :many
SELECT
	sm.id,
	sm.name,
	sm.birth_date,
	sm.height_cm,
	sm.weight_kg,
	c.id AS club_id,
	c.name AS club_name,
	p.rank
FROM sportsmen sm
JOIN clubs c ON c.id = sm.club_id
JOIN participations p ON p.sportsman_id = sm.id
JOIN tournament_sports ts ON ts.id = p.tournament_sport_id
WHERE
	ts.tournament_id = $1
	AND p.rank <= 3
ORDER BY p.rank;

-- Query #8
--
-- Получить перечень соревнований, проведенных в указанном спортивном сооружении в
-- целом либо по определенному виду спорта.
--
-- name: GetTournamentsByPlaceID :many
SELECT
	p.name,
	o.name,
	t.start_at
FROM tournaments t
JOIN places p ON p.id = t.place_id
JOIN organizers o ON o.id = t.organizer_id
JOIN tournament_sports ts ON ts.tournament_id = t.id
WHERE
	t.place_id = $1
	AND (
		ts.sport_id = sqlc.narg('sport_id')
		OR sqlc.narg('sport_id') IS NULL
	);

-- Query #9
--
-- Получить перечень спортивных клубов и число спортсменов этих клубов, участвовавших в
-- спортивных соревнованиях в течение заданного интервала времени.
--
-- name: GetClubActiveSportsmenCountsForPeriod :many
SELECT
	c.name,
	COUNT(s.id)
FROM clubs c
LEFT JOIN sportsmen s ON s.club_id = c.id
LEFT JOIN participations p ON p.sportsman_id = s.id
LEFT JOIN tournament_sports ts ON ts.id = p.tournament_sport_id
LEFT JOIN tournaments t ON t.id = ts.tournament_id
WHERE t.start_at BETWEEN @start_at AND @end_at
GROUP BY
	c.id,
	c.name;

-- Query #10
--
-- Получить список тренеров по определенному виду спорта.
--
-- name: GetTrainersBySportID :many
SELECT t.name
FROM trainers t
JOIN sportsman_sport_trainers sst ON sst.trainer_id = t.id
JOIN sportsman_sports ss ON ss.id = sst.sportsman_sport_id
WHERE ss.sport_id = $1;

-- Query: #11
--
-- Получить список спортсменов, не участвовавших ни в каких соревнованиях в течение
-- определенного периода времени.
--
-- name: GetInactiveSportsmenForPeriod :many
SELECT
	sm.name,
	sm.birth_date,
	sm.height_cm,
	sm.weight_kg
FROM sportsmen sm
WHERE NOT EXISTS (
	SELECT 1
	FROM participations p
	JOIN tournament_sports ts ON ts.id = p.tournament_sport_id
	JOIN tournaments t ON t.id = ts.tournament_id
	WHERE
		t.start_at BETWEEN @start_at AND @end_at
		AND p.sportsman_id = sm.id
);

-- Query: #12
--
-- Получить список организаторов соревнований и число проведенных ими соревнований в
-- течение определенного периода времени.
--
-- name: GetOrganizerTournamentCountsForPeriod :many
SELECT
	o.name,
	o.location,
	COUNT(t.id)
FROM organizers o
LEFT JOIN tournaments t ON t.organizer_id = o.id
GROUP BY
	o.id,
	o.name,
	o.location;

-- Query: #13
--
-- Получить перечень спортивных сооружений и даты проведения на них соревнований в
-- течение определенного периода времени.
--
-- name: GetPlaceTournamentDatesForPeriod :many
SELECT
	p.name,
	p.location,
	ARRAY_AGG(t.start_at)::TIMESTAMPTZ[] as dates
FROM places p
LEFT JOIN tournaments t ON t.place_id = p.id
WHERE t.start_at BETWEEN @start_at AND @end_at
GROUP BY
	p.id,
	p.name,
	p.location;

-- Query: #14 (custom)
--
-- Создаёт манеж и задаёт для него аттрибуты.
--
-- name: InsertArena :exec
WITH
place_type AS (
	SELECT id FROM place_types WHERE attributes_table_name = 'arena_attributes'
),
place AS (
	INSERT INTO places (name, location, type_id)
	VALUES (@name, @location, (SELECT id FROM place_type))
	RETURNING id
)
INSERT INTO arena_attributes (place_id, referees_count, treadmill_length_cm)
VALUES ((SELECT id FROM place), @referees_count, @treadmill_length_cm);

-- Query: #15 (custom)
--
-- Создаёт стадион и задаёт для него аттрибуты.
--
-- name: InsertStadium :exec
WITH
place_type AS (
	SELECT id FROM place_types WHERE attributes_table_name = 'stadium_attributes'
),
place AS (
	INSERT INTO places (name, location, type_id)
	VALUES (@name, @location, place_type.id)
	RETURNING id
)
INSERT INTO stadium_attributes (place_id, width_cm, length_cm, max_spectators, is_outdoor, coating)
VALUES ((SELECT id FROM place), @width_cm, @length_cm, @max_spectators, @is_outdoor, @coating);

-- Query: #16 (custom)
--
-- Создаёт корт и задаёт для него аттрибуты.
--
-- name: InsertCourt :exec
WITH
place_type AS (
	SELECT id FROM place_types WHERE attributes_table_name = 'court_attributes'
),
place AS (
	INSERT INTO places (name, location, type_id)
	VALUES (@name, @location, place_type.id)
	RETURNING id
)
INSERT INTO court_attributes (place_id, width_cm, length_cm, is_outdoor)
VALUES ((SELECT id FROM place), @width_cm, @length_cm, @is_outdoor);

-- Query: #17 (custom)
--
-- Создаёт зал и задаёт для него аттрибуты.
--
-- name: InsertGym :exec
WITH
place_type AS (
	SELECT id FROM place_types WHERE attributes_table_name = 'gym_attributes'
),
place AS (
	INSERT INTO places (name, location, type_id)
	VALUES (@name, @location, place_type.id)
	RETURNING id
)
INSERT INTO gym_attributes (place_id, trainers_count, dumbbells_count, has_bathhouse)
VALUES ((SELECT id FROM place), @trainers_count, @dumbbells_count, @has_bathhouse);

-- Query: #18 (custom)
--
-- Получает спортсмена по идентификатору.
--
-- name: GetSportsmanByID :one
SELECT
	sm.id,
	sm.name,
	sm.birth_date,
	sm.height_cm,
	sm.weight_kg,
	c.id AS club_id,
	c.name AS club_name
FROM sportsmen sm
JOIN clubs c ON c.id = sm.club_id
WHERE sm.id = $1
GROUP BY
	sm.id,
	sm.name,
	sm.birth_date,
	sm.height_cm,
	sm.weight_kg,
	c.id,
	c.name;

-- Query: #19 (custom)
--
-- Получение видов спорта, которыми занимается спортсмен.
--
-- name: GetSportsBySportsmanID :many
SELECT
	s.id,
	s.name
FROM sports s
JOIN sportsman_sports sms ON sms.sport_id = s.id
WHERE sms.sportsman_id = $1;

-- Query: #20 (custom)
--
-- Получает все виды спорта.
--
-- name: GetSports :many
SELECT
	id,
	name
FROM sports;

-- Query: #21 (custom)
--
-- Удаляет спортсмена по ID.
--
-- name: DeleteSportsmanByID :exec
DELETE FROM sportsmen
WHERE id = $1;

-- Query: #22 (custom)
--
-- Получает всех спортсменов.
--
-- name: GetSportsmen :many
SELECT
	sm.id,
	sm.name,
	sm.birth_date,
	sm.height_cm,
	sm.weight_kg,
	c.id AS club_id,
	c.name AS club_name
FROM sportsmen sm
JOIN clubs c ON c.id = sm.club_id
LEFT JOIN sportsman_sports sms ON sms.sportsman_id = sm.id
LEFT JOIN sports s ON s.id = sms.sport_id
GROUP BY
	sm.id,
	sm.name,
	sm.birth_date,
	sm.height_cm,
	sm.weight_kg,
	c.id,
	c.name
ORDER BY sm.id DESC;

-- Query: #23 (custom)
--
-- Обновляет спортсмена по идентификатору.
--
-- name: UpdateSportsmanByID :exec
WITH deleted_sportsman_sport_ids AS (
	DELETE FROM sportsman_sports
	WHERE
		sportsman_id = $1
		AND NOT (sport_id = ANY(@sport_ids::BIGINT[]))
	RETURNING id
),
inserted_sportsman_sport_ids AS (
	INSERT INTO sportsman_sports (sportsman_id, sport_id)
	SELECT
		$1,
		sport_id
	FROM UNNEST(@sport_ids::BIGINT[]) AS sport_id
	ON CONFLICT (sportsman_id, sport_id) DO NOTHING
	RETURNING id
)
UPDATE sportsmen AS sm
SET
	name = $2,
	birth_date = $3,
	height_cm = $4,
	weight_kg = $5,
	club_id = $6
WHERE sm.id = $1;

-- Query: #24 (custom)
--
-- Получает все клубы.
--
-- name: GetClubs :many
SELECT
	id,
	name
FROM clubs;

-- Query: #25 (custom)
--
-- Создаёт спортсмена.
--
-- name: InsertSportsman :exec
WITH sportsman AS (
	INSERT INTO sportsmen (name, birth_date, height_cm, weight_kg, club_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
)
INSERT INTO sportsman_sports (sportsman_id, sport_id)
SELECT
	id,
	sport_id
FROM
	sportsman,
	UNNEST(@sport_ids::BIGINT[]) AS sport_id;

-- Query: #26 (custom)
--
-- Получает все соревнования.
--
-- name: GetTournaments :many
SELECT
	t.id,
	t.start_at,
	p.name AS place_name,
	o.name AS organizer_name
FROM tournaments t
JOIN organizers o ON o.id = t.organizer_id
JOIN places p ON p.id = t.place_id
ORDER BY t.id DESC;
