// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package pg

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type ArenaAttribute struct {
	PlaceID           int64
	RefereesCount     int16
	TreadmillLengthCm int64
}

type Club struct {
	ID        int64
	Name      string
	CreatedAt pgtype.Timestamptz
}

type Organizer struct {
	ID        int64
	Name      string
	Location  pgtype.Text
	CreatedAt pgtype.Timestamptz
}

type Participation struct {
	ID                int64
	TournamentSportID pgtype.Int8
	SportsmanID       pgtype.Int8
	Rank              int16
	Results           pgtype.Text
}

type Place struct {
	ID        int64
	Name      string
	Location  string
	TypeID    int64
	CreatedAt pgtype.Timestamptz
}

type PlaceType struct {
	ID                  int64
	Name                string
	AttributesTableName string
}

type Sport struct {
	ID        int64
	Name      pgtype.Text
	CreatedAt pgtype.Timestamptz
}

type Sportsman struct {
	ID        int64
	Name      string
	BirthDate pgtype.Date
	HeightCm  int16
	WeightKg  pgtype.Numeric
	ClubID    int64
	CreatedAt pgtype.Timestamptz
}

type SportsmanSport struct {
	ID          int64
	SportsmanID int64
	SportID     int64
	Rank        pgtype.Int2
}

type SportsmanSportTrainer struct {
	ID               int64
	SportsmanSportID pgtype.Int8
	TrainerID        pgtype.Int8
}

type StadiumAttribute struct {
	PlaceID       int64
	WidthCm       int64
	LengthCm      int64
	MaxSpectators int16
	IsOutdoor     bool
	Coating       string
}

type Tournament struct {
	ID          int64
	PlaceID     int64
	StartAt     pgtype.Timestamptz
	OrganizerID int64
	CreatedAt   pgtype.Timestamptz
}

type TournamentSport struct {
	ID           int64
	TournamentID pgtype.Int8
	SportID      pgtype.Int8
}

type Trainer struct {
	ID        int64
	Name      string
	CreatedAt pgtype.Timestamptz
}
