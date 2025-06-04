package domain

import "time"

type Tournament struct {
	ID            int64
	OrganizerName string
	PlaceName     string
	SportNames    []string
	StartAt       time.Time
}

type TournamentSport struct {
	ID           int64
	TournamentID int64
	SportName    string
}
