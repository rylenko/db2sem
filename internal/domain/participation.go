package domain

type Participation struct {
	TournamentID  int64
	SportName     string
	SportsmanName string
	Rank          int16
	Results       *string
}
