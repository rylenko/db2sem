package domain

import "time"

type Place struct {
	ID       int64
	Name     string
	Location string
	TypeName string
}

type PlaceWithTournamentDates struct {
	Place
	TournamentDates []time.Time
}
