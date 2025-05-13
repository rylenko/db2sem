package dto

import "time"

type UpdateSportsmanByIDRequest struct {
	ID         int64
	Name       string
	BirthDate  time.Time
	HeightCm   uint16
	WeightKg   float64
	SportNames []string
}
