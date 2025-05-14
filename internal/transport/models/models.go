package models

import "db2sem/internal/domain"

type Club struct {
	ID   int64
	Name string
}

func ConvertFromServiceClub(serviceClub domain.Club) Club {
	return Club{
		ID:   serviceClub.ID,
		Name: serviceClub.Name,
	}
}

type Sport struct {
	ID   int64
	Name string
}

func ConvertFromServiceSport(serviceSport domain.Sport) Sport {
	return Sport{
		ID:   serviceSport.ID,
		Name: serviceSport.Name,
	}
}

func ConvertFromServiceSports(serviceSports []domain.Sport) []Sport {
	sports := make([]Sport, 0, len(serviceSports))

	for _, serviceSport := range serviceSports {
		sport := ConvertFromServiceSport(serviceSport)
		sports = append(sports, sport)
	}

	return sports
}

type Sportsman struct {
	ID              int64
	Name            string
	BirthDateString string
	HeightCm        uint16
	WeightKg        float64
	Club            Club
	Sports          []Sport
}

func ConvertFromServiceSportsmen(serviceSportsmen []domain.Sportsman) []Sportsman {
	sportsmen := make([]Sportsman, 0, len(serviceSportsmen))

	for _, serviceSportsman := range serviceSportsmen {
		sportsman := ConvertFromServiceSportsman(serviceSportsman)
		sportsmen = append(sportsmen, sportsman)
	}

	return sportsmen
}

func ConvertFromServiceSportsman(serviceSportsman domain.Sportsman) Sportsman {
	club := ConvertFromServiceClub(serviceSportsman.Club)
	sports := ConvertFromServiceSports(serviceSportsman.Sports)

	return Sportsman{
		ID:              serviceSportsman.ID,
		Name:            serviceSportsman.Name,
		BirthDateString: serviceSportsman.BirthDate.Format("2006-01-02"),
		HeightCm:        serviceSportsman.HeightCm,
		WeightKg:        serviceSportsman.WeightKg,
		Club:            club,
		Sports:          sports,
	}
}
