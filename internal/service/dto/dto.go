package dto

import "time"

type CreateSportsmanRequest struct {
	Name      string
	BirthDate time.Time
	HeightCm  uint16
	WeightKg  float64
	ClubID    int64
	SportIDs  []int64
}

type CreateTournamentRequest struct {
	PlaceID     int64
	OrganizerID int64
	StartAt     time.Time
	SportIDs    []int64
}

type CreateArenaRequest struct {
	Name              string
	Location          string
	RefereesCount     int16
	TreadmillLengthCm int64
}

type UpdateOrganizerByIDRequest struct {
	ID       int64
	Name     string
	Location *string
}

type UpdateClubByIDRequest struct {
	ID   int64
	Name string
}

type UpdateTrainerByIDRequest struct {
	ID   int64
	Name string
}

type UpdateSportByIDRequest struct {
	ID   int64
	Name string
}

type UpdateArenaByIDRequest struct {
	ID                int64
	Name              string
	Location          string
	RefereesCount     int16
	TreadmillLengthCm int64
}

type UpdateGymByIDRequest struct {
	ID             int64
	Name           string
	Location       string
	TrainersCount  int16
	DumbbellsCount int16
	HasBathhouse   bool
}

type UpdateCourtByIDRequest struct {
	ID        int64
	Name      string
	Location  string
	WidthCm   int64
	LengthCm  int64
	IsOutdoor bool
}

type CreateStadiumRequest struct {
	Name          string
	Location      string
	WidthCm       int64
	LengthCm      int64
	MaxSpectators int16
	IsOutdoor     bool
	Coating       string
}

type UpdateStadiumByIDRequest struct {
	ID            int64
	Name          string
	Location      string
	WidthCm       int64
	LengthCm      int64
	MaxSpectators int16
	IsOutdoor     bool
	Coating       string
}

type UpdateSportsmanByIDRequest struct {
	ID        int64
	Name      string
	BirthDate time.Time
	HeightCm  uint16
	WeightKg  float64
	ClubID    int64
	SportIDs  []int64
}

type GetArenasRequest struct {
	RefereesCount     *int16
	TreadmillLengthCm *int64
}

type GetStadiumsRequest struct {
	WidthCm       *int64
	LengthCm      *int64
	MaxSpectators *int16
	IsOutdoor     *bool
	Coating       *string
}

type GetGymsRequest struct {
	TrainersCount  *int16
	DumbbellsCount *int16
	HasBathhouse   *bool
}

type GetCourtsRequest struct {
	WidthCm   *int64
	LengthCm  *int64
	IsOutdoor *bool
}

type CreateParticipationRequest struct {
	TournamentSportID int64
	SportsmanID       int64
	Rank              int16
	Results           *string
}
