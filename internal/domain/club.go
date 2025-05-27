package domain

type Club struct {
	ID   int64
	Name string
}

type ClubSportsmenCount struct {
	Club
	SportsmenCount uint64
}
