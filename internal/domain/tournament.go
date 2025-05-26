package domain

import "time"

type Tournament struct {
	ID            int64
	OrganizerName string
	PlaceName     string
	SportNames    []string
	StartAt       time.Time
}
