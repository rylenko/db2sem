package dto

import "time"

type InsertSportsmanRequest struct {
	Name      string
	BirthDate time.Time
	HeightCm  uint16
	WeightKg  float64
	ClubID    int64
	SportIDs  []int64
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
