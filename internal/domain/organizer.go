package domain

type Organizer struct {
	ID       int64
	Name     string
	Location *string
}

type OrganizerTournamentsCount struct {
	Organizer
	TournamentsCount uint64
}
