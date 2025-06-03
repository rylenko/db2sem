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

type getGymsForm struct {
	TrainersCount  int16 `query:"trainers_count"`
	DumbbellsCount int16 `query:"dumbbells_count"`
	HasBathhouse   bool  `query:"has_bathhouse"`
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

type createClubForm struct {
	Name string `validate:"required" form:"name"`
}

type createSportForm struct {
	Name string `validate:"required" form:"name"`
}

type createOrganizerForm struct {
	Name     string `validate:"required" form:"name"`
	Location string `form:"location"`
}

type createSportsmanForm struct {
	Name      string  `validate:"required" form:"name"`
	BirthDate string  `validate:"required" form:"birth_date"`
	HeightCm  string  `validate:"required" form:"height_cm"`
	WeightKg  string  `validate:"required" form:"weight_kg"`
	ClubID    int64   `validate:"required" form:"club_id"`
	SportIDs  []int64 `form:"sport_ids"`
}

type updateClubForm struct {
	Name string `validate:"required" form:"name"`
}

type updateSportForm struct {
	Name string `validate:"required" form:"name"`
}

type updateOrganizerForm struct {
	Name     string `validate:"required" form:"name"`
	Location string `form:"location"`
}

type updateSportsmanForm struct {
	Name      string  `validate:"required" form:"name"`
	BirthDate string  `validate:"required" form:"birth_date"`
	HeightCm  string  `validate:"required" form:"height_cm"`
	WeightKg  string  `validate:"required" form:"weight_kg"`
	ClubID    int64   `validate:"required" form:"club_id"`
	SportIDs  []int64 `form:"sport_ids"`
}

type createArenaForm struct {
	Name              string `validate:"required" form:"name"`
	Location          string `validate:"required" form:"location"`
	RefereesCount     int16  `validate:"required" form:"referees_count"`
	TreadmillLengthCm int64  `validate:"required" form:"treadmill_length_cm"`
}

type updateArenaForm struct {
	Name              string `validate:"required" form:"name"`
	Location          string `validate:"required" form:"location"`
	RefereesCount     int16  `validate:"required" form:"referees_count"`
	TreadmillLengthCm int64  `validate:"required" form:"treadmill_length_cm"`
}

type updateGymForm struct {
	Name           string `validate:"required" form:"name"`
	Location       string `validate:"required" form:"location"`
	TrainersCount  int16  `validate:"required" form:"trainers_count"`
	DumbbellsCount int16  `validate:"required" form:"dumbbells_count"`
	HasBathhouse   bool   `form:"has_bathhouse"`
}

type updateCourtForm struct {
	Name      string `validate:"required" form:"name"`
	Location  string `validate:"required" form:"location"`
	WidthCm   int64  `validate:"required" form:"width_cm"`
	LengthCm  int64  `validate:"required" form:"length_cm"`
	IsOutdoor bool   `form:"is_outdoor"`
}

type createStadiumForm struct {
	Name          string `validate:"required" form:"name"`
	Location      string `validate:"required" form:"location"`
	WidthCm       int64  `validate:"required" form:"width_cm"`
	LengthCm      int64  `validate:"required" form:"length_cm"`
	MaxSpectators int16  `validate:"required" form:"max_spectators"`
	IsOutdoor     bool   `form:"is_outdoor"`
	Coating       string `validate:"required" form:"coating"`
}

type updateStadiumForm struct {
	Name          string `validate:"required" form:"name"`
	Location      string `validate:"required" form:"location"`
	WidthCm       int64  `validate:"required" form:"width_cm"`
	LengthCm      int64  `validate:"required" form:"length_cm"`
	MaxSpectators int16  `validate:"required" form:"max_spectators"`
	IsOutdoor     bool   `form:"is_outdoor"`
	Coating       string `validate:"required" form:"coating"`
}
