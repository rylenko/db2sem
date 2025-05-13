package transport

import (
	"strings"

	"db2sem/internal/domain"
)

type Sportsman struct {
	ID               int64
	Name             string
	BirthDateString  string
	HeightCm         uint16
	WeightKg         float64
	SportNames       []string
	SportNamesString string
}

func convertFromServiceSportsmen(serviceSportsmen []domain.Sportsman) []Sportsman {
	sportsmen := make([]Sportsman, 0, len(serviceSportsmen))

	for _, serviceSportsman := range serviceSportsmen {
		sportsman := convertFromServiceSportsman(serviceSportsman)
		sportsmen = append(sportsmen, sportsman)
	}

	return sportsmen
}

func convertFromServiceSportsman(serviceSportsman domain.Sportsman) Sportsman {
	return Sportsman{
		ID:               serviceSportsman.ID,
		Name:             serviceSportsman.Name,
		BirthDateString:  serviceSportsman.BirthDate.Format("2006-01-02"),
		HeightCm:         serviceSportsman.HeightCm,
		WeightKg:         serviceSportsman.WeightKg,
		SportNames:       serviceSportsman.SportNames,
		SportNamesString: strings.Join(serviceSportsman.SportNames, ", "),
	}
}
