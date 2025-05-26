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

func ConvertFromServiceClubs(serviceClubs []domain.Club) []Club {
	clubs := make([]Club, 0, len(serviceClubs))

	for _, serviceClub := range serviceClubs {
		club := ConvertFromServiceClub(serviceClub)
		clubs = append(clubs, club)
	}

	return clubs
}

type PrizeWinner struct {
	Sportsman
	Rank int16
}

func ConvertFromServicePrizeWinner(serviceWinner domain.PrizeWinner) PrizeWinner {
	return PrizeWinner{
		Sportsman: ConvertFromServiceSportsman(serviceWinner.Sportsman),
		Rank:      serviceWinner.Rank,
	}
}

func ConvertFromServicePrizeWinners(serviceWinners []domain.PrizeWinner) []PrizeWinner {
	winners := make([]PrizeWinner, 0, len(serviceWinners))

	for _, serviceWinner := range serviceWinners {
		winner := ConvertFromServicePrizeWinner(serviceWinner)
		winners = append(winners, winner)
	}

	return winners
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

type RankedSportsman struct {
	Sportsman
	Rank *int16
}

func ConvertFromServiceRankedSportsman(serviceSportsman domain.RankedSportsman) RankedSportsman {
	return RankedSportsman{
		Sportsman: ConvertFromServiceSportsman(serviceSportsman.Sportsman),
		Rank:      serviceSportsman.Rank,
	}
}

func ConvertFromServiceRankedSportsmen(serviceSportsmen []domain.RankedSportsman) []RankedSportsman {
	sportsmen := make([]RankedSportsman, 0, len(serviceSportsmen))

	for _, serviceSportsman := range serviceSportsmen {
		sportsman := ConvertFromServiceRankedSportsman(serviceSportsman)
		sportsmen = append(sportsmen, sportsman)
	}

	return sportsmen
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

type Tournament struct {
	ID            int64
	OrganizerName string
	PlaceName     string
	StartAtString string
}

func ConvertFromServiceTournament(serviceTournament domain.Tournament) Tournament {
	return Tournament{
		ID:            serviceTournament.ID,
		OrganizerName: serviceTournament.OrganizerName,
		PlaceName:     serviceTournament.PlaceName,
		StartAtString: serviceTournament.StartAt.Format("02.01.2006 15:04:05"),
	}
}

func ConvertFromServiceTournaments(serviceTournaments []domain.Tournament) []Tournament {
	tournaments := make([]Tournament, 0, len(serviceTournaments))

	for _, serviceTournament := range serviceTournaments {
		tournament := ConvertFromServiceTournament(serviceTournament)
		tournaments = append(tournaments, tournament)
	}

	return tournaments
}

type Organizer struct {
	ID       int64
	Name     string
	Location *string
}

func ConvertFromServiceOrganizer(serviceOrganizer domain.Organizer) Organizer {
	return Organizer{
		ID:       serviceOrganizer.ID,
		Name:     serviceOrganizer.Name,
		Location: serviceOrganizer.Location,
	}
}

func ConvertFromServiceOrganizers(serviceOrganizers []domain.Organizer) []Organizer {
	organizers := make([]Organizer, 0, len(serviceOrganizers))

	for _, serviceOrganizer := range serviceOrganizers {
		organizer := ConvertFromServiceOrganizer(serviceOrganizer)
		organizers = append(organizers, organizer)
	}

	return organizers
}

type Trainer struct {
	ID   int64
	Name string
}

func ConvertFromServiceTrainer(serviceTrainer domain.Trainer) Trainer {
	return Trainer{
		ID:   serviceTrainer.ID,
		Name: serviceTrainer.Name,
	}
}

func ConvertFromServiceTrainers(serviceTrainers []domain.Trainer) []Trainer {
	trainers := make([]Trainer, 0, len(serviceTrainers))

	for _, serviceTrainer := range serviceTrainers {
		trainer := ConvertFromServiceTrainer(serviceTrainer)
		trainers = append(trainers, trainer)
	}

	return trainers
}
