package transport

type getPlaceTournamentsForm struct {
	PlaceID int64 `validate:"required" form:"place_id"`
	SportID int64 `form:"sport_id"`
}

type getSportsmanTrainersForm struct {
	SportsmanID int64 `validate:"required" form:"sportsman_id"`
}

type getSportSportsmenForm struct {
	SportID int64 `validate:"required" form:"sport_id"`
	MinRank int16 `form:"min_rank"`
}

type getSportTrainersForm struct {
	SportID int64 `validate:"required" form:"sport_id"`
}

type getTournamentPrizeWinnersForm struct {
	TournamentID int64 `validate:"required" form:"tournament_id"`
}

type getInactiveSportsmenForm struct {
	StartAt string `validate:"required" form:"start_at"`
	EndAt   string `validate:"required" form:"end_at"`
}

type getClubActiveSportsmenCountsForm struct {
	StartAt string `validate:"required" form:"start_at"`
	EndAt   string `validate:"required" form:"end_at"`
}

type getTournamentsForPeriodForm struct {
	StartAt     string `validate:"required" form:"start_at"`
	EndAt       string `validate:"required" form:"end_at"`
	OrganizerID int64  `form:"organizer_id"`
}

type getTrainerSportsmenForm struct {
	TrainerID int64 `validate:"required" form:"trainer_id"`
	MinRank   int16 `form:"min_rank"`
}

type createSportForm struct {
	Name string `validate:"required" form:"name"`
}

type createSportsmanForm struct {
	Name      string  `validate:"required" form:"name"`
	BirthDate string  `validate:"required" form:"birth_date"`
	HeightCm  string  `validate:"required" form:"height_cm"`
	WeightKg  string  `validate:"required" form:"weight_kg"`
	ClubID    int64   `validate:"required" form:"club_id"`
	SportIDs  []int64 `form:"sport_ids"`
}

type updateSportForm struct {
	Name string `validate:"required" form:"name"`
}

type updateSportsmanForm struct {
	Name      string  `validate:"required" form:"name"`
	BirthDate string  `validate:"required" form:"birth_date"`
	HeightCm  string  `validate:"required" form:"height_cm"`
	WeightKg  string  `validate:"required" form:"weight_kg"`
	ClubID    int64   `validate:"required" form:"club_id"`
	SportIDs  []int64 `form:"sport_ids"`
}
