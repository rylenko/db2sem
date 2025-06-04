-- Query #1.1
--
-- Получить перечень спортивных сооружений указанного типа в целом или
-- удовлетворяющих заданным характеристикам (например, стадионы, вмещающие не менее
-- указанного числа зрителей).
--
-- name: GetArenas :many
SELECT
	p.id,
	p.name,
	p.location,
	aa.referees_count,
	aa.treadmill_length_cm
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
-- name: GetStadiums :many
SELECT
	p.id,
	p.name,
	p.location,
	sa.width_cm,
	sa.length_cm,
	sa.max_spectators,
	sa.is_outdoor,
	sa.coating
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
		sa.coating ILIKE sqlc.narg('coating')
		OR sqlc.narg('coating') IS NULL
	);

-- Query #1.3
--
-- Получить перечень спортивных сооружений указанного типа в целом или
-- удовлетворяющих заданным характеристикам (например, стадионы, вмещающие не менее
-- указанного числа зрителей).
--
-- name: GetCourts :many
SELECT
	p.id,
	p.name,
	p.location,
	ca.width_cm,
	ca.length_cm,
	ca.is_outdoor
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
-- name: GetGyms :many
SELECT
	p.id,
	p.name,
	p.location,
	ga.trainers_count,
	ga.dumbbells_count,
	ga.has_bathhouse
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
	sm.id,
	sm.name,
	sm.birth_date,
	sm.height_cm,
	sm.weight_kg,
	c.id AS club_id,
	c.name AS club_name,
	ss.rank AS rank
FROM sportsmen sm
JOIN clubs c ON c.id = sm.club_id
JOIN sportsman_sports ss ON ss.sportsman_id = sm.id
WHERE
	ss.sport_id = @sport_id
	AND (
		ss.rank >= sqlc.narg('min_rank')
		OR sqlc.narg('min_rank') IS NULL
	)
ORDER BY ss.rank DESC;

-- Query #3
--
-- Получить список спортсменов, тренирующихся у некого тренера в целом либо не ниже
-- определенного разряда.
--
-- name: GetSportsmenByTrainerID :many
SELECT
	sm.id,
	sm.name,
	sm.birth_date,
	sm.height_cm,
	sm.weight_kg,
	c.id AS club_id,
	c.name AS club_name,
	ss.rank AS rank
FROM sportsmen sm
JOIN clubs c ON c.id = sm.club_id
JOIN sportsman_sports ss ON ss.sportsman_id = sm.id
JOIN sportsman_sport_trainers sst ON sst.sportsman_sport_id = ss.id
WHERE
	sst.trainer_id = @trainer_id
	AND (
		ss.rank >= sqlc.narg('min_rank')
		OR sqlc.narg('min_rank') IS NULL
	)
ORDER BY ss.rank DESC;

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
ORDER BY sm.name;

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
	t.id,
	t.start_at,
	p.name AS place_name,
	o.name AS organizer_name,
	ARRAY_AGG(s.name)::TEXT[] as sport_names
FROM tournaments t
JOIN tournament_sports ts ON ts.tournament_id = t.id
JOIN sports s ON s.id = ts.sport_id
JOIN places p ON p.id = t.place_id
JOIN organizers o ON o.id = t.organizer_id
WHERE
	t.start_at BETWEEN @start_at AND @end_at
	AND (
		t.organizer_id = sqlc.narg('organizer_id')
		OR sqlc.narg('organizer_id') IS NULL
	)
GROUP BY
	t.id,
	t.start_at,
	p.name,
	o.name;

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
	t.id,
	p.name AS place_name,
	o.name AS organizer_name,
	t.start_at,
	ARRAY_AGG(s.name)::TEXT[] as sport_names
FROM tournaments t
JOIN tournament_sports ts ON ts.tournament_id = t.id
JOIN sports s ON s.id = ts.sport_id
JOIN places p ON p.id = t.place_id
JOIN organizers o ON o.id = t.organizer_id
WHERE
	t.place_id = $1
GROUP BY
	t.id,
	t.start_at,
	p.name,
	o.name
HAVING
	sqlc.narg('sport_id') = ANY(ARRAY_AGG(s.id))
	OR sqlc.narg('sport_id') IS NULL
ORDER BY t.start_at DESC;

-- Query #9
--
-- Получить перечень спортивных клубов и число спортсменов этих клубов, участвовавших в
-- спортивных соревнованиях в течение заданного интервала времени.
--
-- name: GetClubActiveSportsmenCountsForPeriod :many
SELECT
	c.id,
	c.name,
	COUNT(DISTINCT s.id)
		FILTER (WHERE t.start_at BETWEEN @start_at AND @end_at)
		AS active_sportsmen_count
FROM clubs c
LEFT JOIN sportsmen s ON s.club_id = c.id
LEFT JOIN participations p ON p.sportsman_id = s.id
LEFT JOIN tournament_sports ts ON ts.id = p.tournament_sport_id
LEFT JOIN tournaments t ON t.id = ts.tournament_id
GROUP BY
	c.id,
	c.name
ORDER BY c.name;

-- Query #10
--
-- Получить список тренеров по определенному виду спорта.
--
-- name: GetTrainersBySportID :many
SELECT DISTINCT
	t.id,
	t.name
FROM trainers t
JOIN sportsman_sport_trainers sst ON sst.trainer_id = t.id
JOIN sportsman_sports ss ON ss.id = sst.sportsman_sport_id
WHERE ss.sport_id = $1
ORDER BY t.name;

-- Query: #11
--
-- Получить список спортсменов, не участвовавших ни в каких соревнованиях в течение
-- определенного периода времени.
--
-- name: GetInactiveSportsmenForPeriod :many
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
WHERE NOT EXISTS (
	SELECT 1
	FROM participations p
	JOIN tournament_sports ts ON ts.id = p.tournament_sport_id
	JOIN tournaments t ON t.id = ts.tournament_id
	WHERE
		t.start_at BETWEEN @start_at AND @end_at
		AND p.sportsman_id = sm.id
)
ORDER BY sm.name;

-- Query: #12
--
-- Получить список организаторов соревнований и число проведенных ими соревнований в
-- течение определенного периода времени.
--
-- name: GetOrganizerTournamentCountsForPeriod :many
SELECT
	o.id,
	o.name,
	o.location,
	COUNT(t.id) AS tournaments_count
FROM organizers o
LEFT JOIN tournaments t ON
	t.organizer_id = o.id
	AND t.start_at BETWEEN @start_at AND @end_at
GROUP BY
	o.id,
	o.name,
	o.location
ORDER BY COUNT(t.id) DESC;

-- Query: #13
--
-- Получить перечень спортивных сооружений и даты проведения на них соревнований в
-- течение определенного периода времени.
--
-- name: GetPlacesWithTournamentDatesForPeriod :many
SELECT
	p.id,
	p.name,
	p.location,
	pt.name AS type_name,
	(
		ARRAY_AGG(t.start_at ORDER BY t.start_at DESC)
		FILTER (WHERE t.start_at IS NOT NULL)
	)::TIMESTAMPTZ[] AS tournament_dates
FROM places p
JOIN place_types pt ON pt.id = p.type_id
LEFT JOIN tournaments t
	ON t.place_id = p.id
	AND t.start_at BETWEEN @start_at AND @end_at
GROUP BY
	p.id,
	p.name,
	p.location,
	pt.name
ORDER BY p.name;

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
	VALUES (@name, @location, (SELECT id FROM place_type))
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
	VALUES (@name, @location, (SELECT id FROM place_type))
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
	VALUES (@name, @location, (SELECT id FROM place_type))
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
FROM sports
ORDER BY name;

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
ORDER BY sm.name;

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
	o.name AS organizer_name,
	ARRAY_AGG(s.name)::TEXT[] as sport_names
FROM tournaments t
JOIN tournament_sports ts ON ts.tournament_id = t.id
JOIN sports s ON s.id = ts.sport_id
JOIN organizers o ON o.id = t.organizer_id
JOIN places p ON p.id = t.place_id
GROUP BY
	t.id,
	t.start_at,
	p.name,
	o.name
ORDER BY t.id DESC;

-- Query: #27 (custom)
--
-- Получает спорт по идентификатору.
--
-- name: GetSportByID :one
SELECT
	id,
	name
FROM sports
WHERE id = $1;

-- Query: #28 (custom)
--
-- Удаляет спорт по идентификатору.
--
-- name: DeleteSportByID :exec
DELETE FROM sports
WHERE id = $1;

-- Query: #29 (custom)
--
-- Создаёт вид спорта.
--
-- name: InsertSport :exec
INSERT INTO sports (name)
VALUES ($1);

-- Query: #30 (custom)
--
-- Обновляет вид спорта.
--
-- name: UpdateSportByID :exec
UPDATE sports
SET name = $1
WHERE id = $2;

-- Query: #31 (custom)
--
-- Получает всех тренеров.
--
-- name: GetTrainers :many
SELECT
	id,
	name
FROM trainers;

-- Query: #32 (custom)
--
-- Получает всех организаторов.
--
-- name: GetOrganizers :many
SELECT
	id,
	name,
	location
FROM organizers
ORDER BY name;

-- Query: #33 (custom)
--
-- Получает сооружения.
--
-- name: GetPlaces :many
SELECT
	p.id,
	p.name,
	p.location,
	pt.name AS type_name
FROM places p
JOIN place_types pt ON pt.id = p.type_id;

-- Query: #34 (custom)
--
-- Удаляет сооружение.
--
-- name: DeletePlaceByID :exec
DELETE FROM places
WHERE id = $1;

-- Query: #35 (custom)
--
-- Получить арену.
--
-- name: GetArenaByID :one
SELECT
	p.id,
	p.name,
	p.location,
	aa.referees_count,
	aa.treadmill_length_cm
FROM places p
JOIN arena_attributes aa ON aa.place_id = p.id
WHERE p.id = $1;

-- Query: #36 (custom)
--
-- Обновить арену.
--
-- name: UpdateArenaByID :exec
WITH updated_attributes AS (
	UPDATE arena_attributes
	SET
		referees_count = $1,
		treadmill_length_cm = $2
	WHERE place_id = $3
	RETURNING place_id
)
UPDATE places
SET
	name = $4,
	location = $5
WHERE id = $3;

-- Query: #37 (custom)
--
-- Получить стадион.
--
-- name: GetStadiumByID :one
SELECT
	p.id,
	p.name,
	p.location,
	sa.width_cm,
	sa.length_cm,
	sa.max_spectators,
	sa.is_outdoor,
	sa.coating
FROM places p
JOIN stadium_attributes sa ON sa.place_id = p.id
WHERE p.id = $1;

-- Query: #38 (custom)
--
-- Обновить стадион.
--
-- name: UpdateStadiumByID :exec
WITH updated_attributes AS (
	UPDATE stadium_attributes
	SET
		width_cm = $1,
		length_cm = $2,
		max_spectators = $3,
		is_outdoor = $4,
		coating = $5
	WHERE place_id = $6
	RETURNING place_id
)
UPDATE places
SET
	name = $7,
	location = $8
WHERE id = $6;

-- Query: #39 (custom)
--
-- Получить корт.
--
-- name: GetCourtByID :one
SELECT
	p.id,
	p.name,
	p.location,
	ca.width_cm,
	ca.length_cm,
	ca.is_outdoor
FROM places p
JOIN court_attributes ca ON ca.place_id = p.id
WHERE p.id = $1;

-- Query: #40 (custom)
--
-- Обновить корт.
--
-- name: UpdateCourtByID :exec
WITH updated_attributes AS (
	UPDATE court_attributes
	SET
		width_cm = $1,
		length_cm = $2,
		is_outdoor = $3
	WHERE place_id = $4
	RETURNING place_id
)
UPDATE places
SET
	name = $5,
	location = $6
WHERE id = $4;

-- Query: #41 (custom)
--
-- Получить зал.
--
-- name: GetGymByID :one
SELECT
	p.id,
	p.name,
	p.location,
	ga.trainers_count,
	ga.dumbbells_count,
	ga.has_bathhouse
FROM places p
JOIN gym_attributes ga ON ga.place_id = p.id
WHERE p.id = $1;

-- Query: #42 (custom)
--
-- Обновить зал.
--
-- name: UpdateGymByID :exec
WITH updated_attributes AS (
	UPDATE gym_attributes
	SET
		trainers_count = $1,
		dumbbells_count = $2,
		has_bathhouse = $3
	WHERE place_id = $4
	RETURNING place_id
)
UPDATE places
SET
	name = $5,
	location = $6
WHERE id = $4;

-- Query: #43 (custom)
--
-- Создать организатора.
--
-- name: InsertOrganizer :exec
INSERT INTO organizers (name, location)
VALUES ($1, $2);

-- Query: #44 (custom)
--
-- Обновляет организатора.
--
-- name: UpdateOrganizerByID :exec
UPDATE organizers
SET
	name = $1,
	location = $2
WHERE id = $3;

-- Query: #45 (custom)
--
-- Удаляет организатора по ID.
--
-- name: DeleteOrganizerByID :exec
DELETE FROM organizers
WHERE id = $1;

-- Query: #46 (custom)
--
-- Получает организатора по идентификатору.
--
-- name: GetOrganizerByID :one
SELECT
	id,
	name,
	location
FROM organizers
WHERE id = $1;

-- Query: #47 (custom)
--
-- Создать клуб.
--
-- name: InsertClub :exec
INSERT INTO clubs (name)
VALUES ($1);

-- Query: #48 (custom)
--
-- Обновляет клуб.
--
-- name: UpdateClubByID :exec
UPDATE clubs
SET name = $1
WHERE id = $2;

-- Query: #49 (custom)
--
-- Удаляет клуб по ID.
--
-- name: DeleteClubByID :exec
DELETE FROM clubs
WHERE id = $1;

-- Query: #50 (custom)
--
-- Получает клуб по идентификатору.
--
-- name: GetClubByID :one
SELECT
	id,
	name
FROM clubs
WHERE id = $1;

-- Query: #51 (custom)
--
-- Создать тренера.
--
-- name: InsertTrainer :exec
INSERT INTO trainers (name)
VALUES ($1);

-- Query: #52 (custom)
--
-- Обновляет тренера.
--
-- name: UpdateTrainerByID :exec
UPDATE trainers
SET name = $1
WHERE id = $2;

-- Query: #53 (custom)
--
-- Удаляет клуб по ID.
--
-- name: DeleteTrainerByID :exec
DELETE FROM trainers
WHERE id = $1;

-- Query: #54 (custom)
--
-- Получает клуб по идентификатору.
--
-- name: GetTrainerByID :one
SELECT
	id,
	name
FROM trainers
WHERE id = $1;

-- Query: #55 (custom)
--
-- Создать соревнование.
--
-- name: InsertTournament :exec
WITH tournament AS (
	INSERT INTO tournaments (place_id, organizer_id, start_at)
	VALUES ($1, $2, $3)
	RETURNING id
)
INSERT INTO tournament_sports (tournament_id, sport_id)
SELECT
	id,
	sport_id
FROM
	tournament,
	UNNEST(@sport_ids::BIGINT[]) AS sport_id;

-- Query: #56 (custom)
--
-- Добавляет участие.
--
-- name: InsertParticipation :exec
INSERT INTO participations (tournament_sport_id, sportsman_id, rank, results)
VALUES ($1, $2, $3, $4);

-- Query: #57 (custom)
--
-- Получает все участия.
--
-- name: GetParticipations :many
SELECT
	ts.tournament_id,
	s.name AS sport_name,
	sm.name AS sportsman_name,
	p.rank,
	p.results
FROM participations p
JOIN tournament_sports ts ON ts.id = p.tournament_sport_id
JOIN tournaments t ON t.id = ts.tournament_id
JOIN sports s ON s.id = ts.sport_id
JOIN sportsmen sm ON sm.id = p.sportsman_id
ORDER BY t.id DESC;

-- Query: #58 (custom)
--
-- Получает все соревнования с их видами спорта.
--
-- name: GetTournamentSports :many
SELECT
	ts.id,
	ts.tournament_id,
	s.name AS sport_name
FROM tournament_sports ts
JOIN sports s ON s.id = ts.sport_id
JOIN tournaments t ON t.id = ts.tournament_id
ORDER BY t.id DESC;
