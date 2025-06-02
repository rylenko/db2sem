package models

import "db2sem/internal/domain"

type Gym struct {
	ID             int64
	Name           string
	Location       string
	TrainersCount  int16
	DumbbellsCount int16
	HasBathhouse   bool
}

func ConvertFromServiceGym(sStadium domain.Gym) Gym {
	return Gym{
		ID:             sStadium.ID,
		Name:           sStadium.Name,
		Location:       sStadium.Location,
		TrainersCount:  sStadium.TrainersCount,
		DumbbellsCount: sStadium.DumbbellsCount,
		HasBathhouse:   sStadium.HasBathhouse,
	}
}

func ConvertFromServiceGyms(sArenas []domain.Gym) []Gym {
	arenas := make([]Gym, 0, len(sArenas))

	for _, sArena := range sArenas {
		arenas = append(arenas, ConvertFromServiceGym(sArena))
	}

	return arenas
}

type Court struct {
	ID        int64
	Name      string
	Location  string
	WidthCm   int64
	LengthCm  int64
	IsOutdoor bool
}

func ConvertFromServiceCourt(sStadium domain.Court) Court {
	return Court{
		ID:        sStadium.ID,
		Name:      sStadium.Name,
		Location:  sStadium.Location,
		WidthCm:   sStadium.WidthCm,
		LengthCm:  sStadium.LengthCm,
		IsOutdoor: sStadium.IsOutdoor,
	}
}

func ConvertFromServiceCourts(sArenas []domain.Court) []Court {
	arenas := make([]Court, 0, len(sArenas))

	for _, sArena := range sArenas {
		arenas = append(arenas, ConvertFromServiceCourt(sArena))
	}

	return arenas
}

type Stadium struct {
	ID            int64
	Name          string
	Location      string
	WidthCm       int64
	LengthCm      int64
	MaxSpectators int16
	IsOutdoor     bool
	Coating       string
}

func ConvertFromServiceStadium(sStadium domain.Stadium) Stadium {
	return Stadium{
		ID:            sStadium.ID,
		Name:          sStadium.Name,
		Location:      sStadium.Location,
		WidthCm:       sStadium.WidthCm,
		LengthCm:      sStadium.LengthCm,
		MaxSpectators: sStadium.MaxSpectators,
		IsOutdoor:     sStadium.IsOutdoor,
		Coating:       sStadium.Coating,
	}
}

func ConvertFromServiceStadiums(sArenas []domain.Stadium) []Stadium {
	arenas := make([]Stadium, 0, len(sArenas))

	for _, sArena := range sArenas {
		arenas = append(arenas, ConvertFromServiceStadium(sArena))
	}

	return arenas
}

type Arena struct {
	ID                int64
	Name              string
	Location          string
	RefereesCount     int16
	TreadmillLengthCm int64
}

func ConvertFromServiceArena(sArena domain.Arena) Arena {
	return Arena{
		ID:                sArena.ID,
		Name:              sArena.Name,
		Location:          sArena.Location,
		RefereesCount:     sArena.RefereesCount,
		TreadmillLengthCm: sArena.TreadmillLengthCm,
	}
}

func ConvertFromServiceArenas(sArenas []domain.Arena) []Arena {
	arenas := make([]Arena, 0, len(sArenas))

	for _, sArena := range sArenas {
		arenas = append(arenas, ConvertFromServiceArena(sArena))
	}

	return arenas
}

type ClubSportsmenCount struct {
	Club
	SportsmenCount uint64
}

func ConvertFromServiceClubSportsmenCount(serviceClub domain.ClubSportsmenCount) ClubSportsmenCount {
	return ClubSportsmenCount{
		Club:           ConvertFromServiceClub(serviceClub.Club),
		SportsmenCount: serviceClub.SportsmenCount,
	}
}

func ConvertFromServiceClubSportsmenCounts(serviceClubs []domain.ClubSportsmenCount) []ClubSportsmenCount {
	clubs := make([]ClubSportsmenCount, 0, len(serviceClubs))

	for _, serviceClub := range serviceClubs {
		club := ConvertFromServiceClubSportsmenCount(serviceClub)
		clubs = append(clubs, club)
	}

	return clubs
}

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
	SportNames    []string
	StartAtString string
}

func ConvertFromServiceTournament(serviceTournament domain.Tournament) Tournament {
	return Tournament{
		ID:            serviceTournament.ID,
		OrganizerName: serviceTournament.OrganizerName,
		PlaceName:     serviceTournament.PlaceName,
		SportNames:    serviceTournament.SportNames,
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

type OrganizerTournamentsCount struct {
	Organizer
	TournamentsCount uint64
}

func ConvertFromServiceOrganizerTournamentsCount(
	serviceOrganizer domain.OrganizerTournamentsCount) OrganizerTournamentsCount {
	return OrganizerTournamentsCount{
		Organizer:        ConvertFromServiceOrganizer(serviceOrganizer.Organizer),
		TournamentsCount: serviceOrganizer.TournamentsCount,
	}
}

func ConvertFromServiceOrganizerTournamentCounts(
	serviceOrganizers []domain.OrganizerTournamentsCount) []OrganizerTournamentsCount {
	organizers := make([]OrganizerTournamentsCount, 0, len(serviceOrganizers))

	for _, serviceOrganizer := range serviceOrganizers {
		organizer := ConvertFromServiceOrganizerTournamentsCount(serviceOrganizer)
		organizers = append(organizers, organizer)
	}

	return organizers
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

type PlaceWithTournamentDates struct {
	Place
	TournamentDateStrings []string
}

func ConvertFromServicePlaceWithTournamentDates(servicePlace domain.PlaceWithTournamentDates) PlaceWithTournamentDates {
	dates := make([]string, 0, len(servicePlace.TournamentDates))

	for _, date := range servicePlace.TournamentDates {
		dates = append(dates, date.Format("02.01.2006 15:04:05"))
	}

	return PlaceWithTournamentDates{
		Place:                 ConvertFromServicePlace(servicePlace.Place),
		TournamentDateStrings: dates,
	}
}

func ConvertFromServicePlacesWithTournamentDates(
	servicePlaces []domain.PlaceWithTournamentDates) []PlaceWithTournamentDates {
	places := make([]PlaceWithTournamentDates, 0, len(servicePlaces))

	for _, servicePlace := range servicePlaces {
		place := ConvertFromServicePlaceWithTournamentDates(servicePlace)
		places = append(places, place)
	}

	return places
}

type Place struct {
	ID       int64
	Name     string
	Location string
	TypeName string
}

func ConvertFromServicePlace(servicePlace domain.Place) Place {
	return Place{
		ID:       servicePlace.ID,
		Name:     servicePlace.Name,
		Location: servicePlace.Location,
		TypeName: servicePlace.TypeName,
	}
}

func ConvertFromServicePlaces(servicePlaces []domain.Place) []Place {
	places := make([]Place, 0, len(servicePlaces))

	for _, servicePlace := range servicePlaces {
		place := ConvertFromServicePlace(servicePlace)
		places = append(places, place)
	}

	return places
}
