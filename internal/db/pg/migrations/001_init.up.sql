CREATE TABLE IF NOT EXISTS place_types (
	id                    BIGSERIAL PRIMARY KEY,
	name                  TEXT      NOT NULL UNIQUE,
	attributes_table_name TEXT      NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS places (
	id              BIGSERIAL             PRIMARY KEY,
	name            VARCHAR(255)          NOT NULL,
	location        TEXT                  NOT NULL,
	type_id         BIGINT                NOT NULL REFERENCES place_types(id) ON DELETE CASCADE,
	created_at      TIMESTAMPTZ           NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS arena_attributes (
	place_id            BIGINT    PRIMARY KEY REFERENCES places(id) ON DELETE CASCADE,
	referees_count      SMALLINT  NOT NULL CHECK (referees_count >= 0),
	treadmill_length_cm BIGINT    NOT NULL CHECK (treadmill_length_cm > 0)
);

CREATE TABLE IF NOT EXISTS court_attributes (
	place_id   BIGINT  PRIMARY KEY REFERENCES places(id) ON DELETE CASCADE,
	width_cm   BIGINT  NOT NULL CHECK (width_cm > 0),
	length_cm  BIGINT  NOT NULL CHECK (length_cm > 0),
	is_outdoor BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS gym_attributes (
	place_id        BIGINT   PRIMARY KEY REFERENCES places(id) ON DELETE CASCADE,
	trainers_count  SMALLINT NOT NULL CHECK (trainers_count >= 0),
	dumbbells_count SMALLINT NOT NULL CHECK (dumbbells_count >= 0),
	has_bathhouse   BOOLEAN  NOT NULL
);

CREATE TABLE IF NOT EXISTS stadium_attributes (
	place_id       BIGINT       PRIMARY KEY REFERENCES places(id) ON DELETE CASCADE,
	width_cm       BIGINT       NOT NULL CHECK (width_cm > 0),
	length_cm      BIGINT       NOT NULL CHECK (length_cm > 0),
	max_spectators SMALLINT     NOT NULL CHECK (max_spectators >= 0),
	is_outdoor     BOOLEAN      NOT NULL,
	coating        VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS clubs (
	id         BIGSERIAL    PRIMARY KEY,
	name       VARCHAR(255) NOT NULL,
	created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS organizers (
	id         BIGSERIAL    PRIMARY KEY,
	name       VARCHAR(255) NOT NULL,
	location   TEXT,
	created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS sports (
	id         BIGSERIAL    PRIMARY KEY,
	name       VARCHAR(255) NOT NULL UNIQUE,
	created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tournaments (
	id           BIGSERIAL    PRIMARY KEY,
	place_id     BIGINT       NOT NULL REFERENCES places(id) ON DELETE CASCADE,
	start_at     TIMESTAMPTZ  NOT NULL,
	organizer_id BIGINT       NOT NULL REFERENCES organizers(id) ON DELETE CASCADE,
	created_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tournament_sports (
	id            BIGSERIAL PRIMARY KEY,
	tournament_id BIGINT    NOT NULL REFERENCES tournaments(id) ON DELETE CASCADE,
	sport_id      BIGINT    NOT NULL REFERENCES sports(id) ON DELETE CASCADE,
	UNIQUE (tournament_id, sport_id)
);

CREATE TABLE IF NOT EXISTS sportsmen (
	id         BIGSERIAL     PRIMARY KEY,
	name       VARCHAR(255)  NOT NULL,
	birth_date DATE          NOT NULL,
	height_cm  SMALLINT      NOT NULL CHECK (height_cm > 0),
	weight_kg  NUMERIC(5, 2) NOT NULL CHECK (weight_kg > 0),
	club_id    BIGINT        NOT NULL REFERENCES clubs(id) ON DELETE CASCADE,
	created_at TIMESTAMPTZ   NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS participations (
	id                  BIGSERIAL PRIMARY KEY,
	tournament_sport_id BIGINT    NOT NULL REFERENCES tournament_sports(id) ON DELETE CASCADE,
	sportsman_id        BIGINT    NOT NULL REFERENCES sportsmen(id) ON DELETE CASCADE,
	rank                SMALLINT  NOT NULL CHECK (rank >= 0),
	results             TEXT,
	UNIQUE (tournament_sport_id, sportsman_id)
);

CREATE TABLE IF NOT EXISTS sportsman_sports (
	id           BIGSERIAL    PRIMARY KEY,
	sportsman_id BIGINT       NOT NULL REFERENCES sportsmen(id) ON DELETE CASCADE,
	sport_id     BIGINT       NOT NULL REFERENCES sports(id) ON DELETE CASCADE,
	rank         SMALLINT     CHECK (rank IS NULL OR rank >= 0),
	UNIQUE (sportsman_id, sport_id)
);

CREATE TABLE IF NOT EXISTS trainers (
	id         BIGSERIAL    PRIMARY KEY,
	name       VARCHAR(255) NOT NULL,
	created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS sportsman_sport_trainers (
	id                 BIGSERIAL PRIMARY KEY,
	sportsman_sport_id BIGINT    NOT NULL REFERENCES sportsman_sports(id) ON DELETE CASCADE,
	trainer_id         BIGINT    NOT NULL REFERENCES trainers(id) ON DELETE CASCADE,
	UNIQUE (sportsman_sport_id, trainer_id)
);
