package domain

import "time"

type Sportsman struct {
	ID        int64
	Name      string
	BirthDate time.Time
	HeightCm  uint16
	WeightKg  float64
	Club      Club
	Sports    []Sport
}
