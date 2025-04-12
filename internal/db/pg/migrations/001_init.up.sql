CREATE TABLE IF NOT EXISTS arena_attributes (
	place_id            BIGINT    PRIMARY KEY REFERENCES places(id),
	referees_count      SMALLINT  NOT NULL CHECK (referees_count >= 0),
	treadmill_length_cm BIGINT    NOT NULL CHECK (treadmill_length_cm > 0)
);

CREATE TABLE IF NOT EXISTS stadium_attributes (
	place_id       BIGINT       PRIMARY KEY REFERENCES places(id),
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

CREATE TABLE IF NOT EXISTS place_types (
	id                    BIGSERIAL PRIMARY KEY,
	name                  TEXT      NOT NULL UNIQUE,
	attributes_table_name TEXT      NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS places (
	id              BIGSERIAL             PRIMARY KEY,
	name            VARCHAR(255)          NOT NULL,
	location        TEXT                  NOT NULL,
	type_id         BIGINT                NOT NULL REFERENCES place_types(id),
	created_at      TIMESTAMPTZ           NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS sports (
	id         BIGSERIAL    PRIMARY KEY,
	name       VARCHAR(255) UNIQUE,
	created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tournament_sports (
	tournament_id BIGINT REFERENCES tournaments(id),
	sport_id      BIGINT REFERENCES sports(id),
	PRIMARY KEY (tournament_id, sport_id)
);

CREATE TABLE IF NOT EXISTS tournaments (
	id           BIGSERIAL    PRIMARY KEY,
	place_id     BIGINT       NOT NULL REFERENCES places(id),
	start_at     TIMESTAMPTZ  NOT NULL,
	organizer_id BIGINT       NOT NULL REFERENCES organizers(id),
	created_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS sportsmen (
	id         BIGSERIAL     PRIMARY KEY,
	name       VARCHAR(255)  NOT NULL,
	birth_date DATE          NOT NULL,
	height_cm  SMALLINT      NOT NULL CHECK (height_cm > 0),
	weight_kg  NUMERIC(5, 2) NOT NULL CHECK (weight_kg > 0),
	club_id    BIGINT        NOT NULL REFERENCES clubs(id),
	created_at TIMESTAMPTZ   NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS participations (
	tournament_sport_id BIGINT   REFERENCES tournament_sports(id),
	sportsman_id        BIGINT   REFERENCES sportsmen(id),
	rank                SMALLINT NOT NULL CHECK (rank >= 0),
	results             TEXT,
	PRIMARY KEY (tournament_sport_id, sportsman_id)
);

CREATE TABLE IF NOT EXISTS sportsman_sports (
	id           BIGSERIAL    PRIMARY KEY,
	sportsman_id BIGINT       NOT NULL REFERENCES sportsmen(id),
	sport_id     BIGINT       NOT NULL REFERENCES sports(id),
	rank         SMALLINT     CHECK (rank IS NULL OR rank >= 0),
	UNIQUE (sportsman_id, sport_id)
);

CREATE TABLE IF NOT EXISTS trainers (
	id         BIGSERIAL    PRIMARY KEY,
	name       VARCHAR(255) NOT NULL,
	created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS sportsman_sport_trainers (
	sportsman_sport_id BIGINT REFERENCES sportsman_sports(id),
	trainer_id         BIGINT REFERENCES trainers(id),
	PRIMARY KEY (sportsman_sport_id, trainer_id)
);
