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

type UpdateSportByIDRequest struct {
	ID   int64
	Name string
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
