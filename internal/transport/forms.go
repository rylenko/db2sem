package transport

type getArenasForm struct {
	RefereesCount     int16 `query:"referees_count"`
	TreadmillLengthCm int64 `query:"treadmill_length_cm"`
}

type getStadiumsForm struct {
	WidthCm       int64  `query:"width_cm"`
	LengthCm      int64  `query:"length_cm"`
	MaxSpectators int16  `query:"max_spectators"`
	IsOutdoor     bool   `query:"is_outdoor"`
	Coating       string `query:"coating"`
}

type getCourtsForm struct {
	WidthCm   int64 `query:"width_cm"`
	LengthCm  int64 `query:"length_cm"`
	IsOutdoor bool  `query:"is_outdoor"`
}

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

type getPlacesWithTournamentDatesForm struct {
	StartAt string `validate:"required" form:"start_at"`
	EndAt   string `validate:"required" form:"end_at"`
}

type getOrganizerTournamentCountsForm struct {
	StartAt string `validate:"required" form:"start_at"`
	EndAt   string `validate:"required" form:"end_at"`
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
