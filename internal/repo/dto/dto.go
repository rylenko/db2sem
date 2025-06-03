package dto

import "time"

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

type InsertSportsmanRequest struct {
	Name      string
	BirthDate time.Time
	HeightCm  uint16
	WeightKg  float64
	ClubID    int64
	SportIDs  []int64
}

type InsertArenaRequest struct {
	Name              string
	Location          string
	RefereesCount     int16
	TreadmillLengthCm int64
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
