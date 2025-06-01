package transport

import (
	"fmt"
	"strconv"
	"time"

	servicedto "db2sem/internal/service/dto"
	"db2sem/internal/transport/models"

	"github.com/gofiber/fiber/v2"
)

type Transport struct {
	requestReader requestReader
	service       service
}

func New(requestReader requestReader, service service) *Transport {
	return &Transport{
		requestReader: requestReader,
		service:       service,
	}
}

func (t *Transport) CreateSport(fiberCtx *fiber.Ctx) error {
	var form createSportForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	err := t.service.CreateSport(fiberCtx.Context(), form.Name)
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect("/sports/", fiber.StatusFound)
}

func (t *Transport) CreateSportsman(fiberCtx *fiber.Ctx) error {
	var form createSportsmanForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	birthDate, err := time.Parse("2006-01-02", form.BirthDate)
	if err != nil {
		return fmt.Errorf("parse birth date: %w", err)
	}

	heightCm, err := strconv.ParseUint(form.HeightCm, 10, 16)
	if err != nil {
		return fmt.Errorf("parse height cm: %w", err)
	}

	weightKg, err := strconv.ParseFloat(form.WeightKg, 64)
	if err != nil {
		return fmt.Errorf("parse weight kg: %w", err)
	}

	err = t.service.CreateSportsman(fiberCtx.Context(), servicedto.CreateSportsmanRequest{
		Name:      form.Name,
		BirthDate: birthDate,
		HeightCm:  uint16(heightCm),
		WeightKg:  weightKg,
		ClubID:    form.ClubID,
		SportIDs:  form.SportIDs,
	})
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect("/sportsmen/", fiber.StatusFound)
}

func (t *Transport) DeleteSport(fiberCtx *fiber.Ctx) error {
	sportID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sport ID: %w", err)
	}

	if err := t.service.DeleteSportByID(fiberCtx.Context(), sportID); err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect("/sports/", fiber.StatusFound)
}

func (t *Transport) DeleteSportsman(fiberCtx *fiber.Ctx) error {
	sportsmanID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sportsman ID: %w", err)
	}

	if err := t.service.DeleteSportsmanByID(fiberCtx.Context(), sportsmanID); err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect("/sportsmen/", fiber.StatusFound)
}

func (t *Transport) RenderIndexPage(fiberCtx *fiber.Ctx) error {
	return fiberCtx.Render("index", fiber.Map{})
}

func (t *Transport) RenderTournamentPrizeWinnersGetPage(fiberCtx *fiber.Ctx) error {
	serviceTournaments, err := t.service.GetTournaments(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	tournaments := models.ConvertFromServiceTournaments(serviceTournaments)

	return fiberCtx.Render("queries/tournament_prize_winners", fiber.Map{
		"Tournaments": tournaments,
	})
}

func (t *Transport) RenderTournamentPrizeWinnersPostPage(fiberCtx *fiber.Ctx) error {
	serviceTournaments, err := t.service.GetTournaments(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	tournaments := models.ConvertFromServiceTournaments(serviceTournaments)

	var form getTournamentPrizeWinnersForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	serviceWinners, err := t.service.GetPrizeWinnersByTournamentID(fiberCtx.Context(), form.TournamentID)
	if err != nil {
		return fmt.Errorf("get trainers: %w", err)
	}

	winners := models.ConvertFromServicePrizeWinners(serviceWinners)

	return fiberCtx.Render("queries/tournament_prize_winners", fiber.Map{
		"Tournaments":  tournaments,
		"PrizeWinners": winners,
	})
}

func (t *Transport) RenderQueriesPage(fiberCtx *fiber.Ctx) error {
	return fiberCtx.Render("queries/index", fiber.Map{})
}

func (t *Transport) RenderSportsmanPage(fiberCtx *fiber.Ctx) error {
	sportsmanID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sportsman ID: %w", err)
	}

	serviceSportsman, err := t.service.GetSportsmanByID(fiberCtx.Context(), sportsmanID)
	if err != nil {
		return fmt.Errorf("get sportsman: %w", err)
	}

	if serviceSportsman == nil {
		return fiberCtx.Status(fiber.StatusNotFound).SendString("Sportsman not found")
	}

	sportsman := models.ConvertFromServiceSportsman(*serviceSportsman)

	serviceSports, err := t.service.GetSports(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("get sport names: %w", err)
	}

	sports := models.ConvertFromServiceSports(serviceSports)

	serviceClubs, err := t.service.GetClubs(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("get clubs: %w", err)
	}

	clubs := models.ConvertFromServiceClubs(serviceClubs)

	return fiberCtx.Render("sportsman", fiber.Map{
		"Sportsman": sportsman,
		"Clubs":     clubs,
		"Sports":    sports,
	})
}

func (t *Transport) RenderSportsmenPage(fiberCtx *fiber.Ctx) error {
	serviceSportsmen, err := t.service.GetSportsmen(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	sportsmen := models.ConvertFromServiceSportsmen(serviceSportsmen)

	serviceSports, err := t.service.GetSports(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("get sport names: %w", err)
	}

	sports := models.ConvertFromServiceSports(serviceSports)

	serviceClubs, err := t.service.GetClubs(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("get clubs: %w", err)
	}

	clubs := models.ConvertFromServiceClubs(serviceClubs)

	return fiberCtx.Render("sportsmen", fiber.Map{
		"Sportsmen": sportsmen,
		"Sports":    sports,
		"Clubs":     clubs,
	})
}

func (t *Transport) RenderSportPage(fiberCtx *fiber.Ctx) error {
	sportID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sportsman ID: %w", err)
	}

	serviceSport, err := t.service.GetSportByID(fiberCtx.Context(), sportID)
	if err != nil {
		return fmt.Errorf("get sportsman: %w", err)
	}

	if serviceSport == nil {
		return fiberCtx.Status(fiber.StatusNotFound).SendString("Sport not found")
	}

	sport := models.ConvertFromServiceSport(*serviceSport)

	return fiberCtx.Render("sport", fiber.Map{
		"Sport": sport,
	})
}

func (t *Transport) RenderSportsPage(fiberCtx *fiber.Ctx) error {
	serviceSports, err := t.service.GetSports(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	sports := models.ConvertFromServiceSports(serviceSports)

	return fiberCtx.Render("sports", fiber.Map{
		"Sports": sports,
	})
}

func (t *Transport) RenderSportsmanTrainersGetPage(fiberCtx *fiber.Ctx) error {
	serviceSportsmen, err := t.service.GetSportsmen(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	sportsmen := models.ConvertFromServiceSportsmen(serviceSportsmen)

	return fiberCtx.Render("queries/sportsman_trainers", fiber.Map{
		"Sportsmen": sportsmen,
	})
}

func (t *Transport) RenderSportsmanTrainersPostPage(fiberCtx *fiber.Ctx) error {
	serviceSportsmen, err := t.service.GetSportsmen(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("get sportsman: %w", err)
	}

	sportsmen := models.ConvertFromServiceSportsmen(serviceSportsmen)

	var form getSportsmanTrainersForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	serviceTrainers, err := t.service.GetTrainersBySportsmanID(fiberCtx.Context(), form.SportsmanID)
	if err != nil {
		return fmt.Errorf("get trainers: %w", err)
	}

	trainers := models.ConvertFromServiceTrainers(serviceTrainers)

	return fiberCtx.Render("queries/sportsman_trainers", fiber.Map{
		"Sportsmen": sportsmen,
		"Trainers":  trainers,
	})
}

func (t *Transport) RenderPlaceTournamentsGetPage(fiberCtx *fiber.Ctx) error {
	servicePlaces, err := t.service.GetPlaces(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	places := models.ConvertFromServicePlaces(servicePlaces)

	serviceSports, err := t.service.GetSports(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	sports := models.ConvertFromServiceSports(serviceSports)

	return fiberCtx.Render("queries/place_tournaments", fiber.Map{
		"Places": places,
		"Sports": sports,
	})
}

func (t *Transport) RenderPlaceTournamentsPostPage(fiberCtx *fiber.Ctx) error {
	servicePlaces, err := t.service.GetPlaces(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	places := models.ConvertFromServicePlaces(servicePlaces)

	serviceSports, err := t.service.GetSports(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	sports := models.ConvertFromServiceSports(serviceSports)

	var form getPlaceTournamentsForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	var sportID *int64
	if form.SportID > 0 {
		sportID = &form.SportID
	}

	serviceTournaments, err := t.service.GetTournamentsByPlaceID(fiberCtx.Context(), form.PlaceID, sportID)
	if err != nil {
		return fmt.Errorf("get tournaments: %w", err)
	}

	tournaments := models.ConvertFromServiceTournaments(serviceTournaments)

	return fiberCtx.Render("queries/place_tournaments", fiber.Map{
		"Places":      places,
		"Sports":      sports,
		"Tournaments": tournaments,
	})
}

func (t *Transport) RenderInactiveSportsmenGetPage(fiberCtx *fiber.Ctx) error {
	return fiberCtx.Render("queries/inactive_sportsmen", fiber.Map{})
}

func (t *Transport) RenderInactiveSportsmenPostPage(fiberCtx *fiber.Ctx) error {
	var form getInactiveSportsmenForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	startAt, err := time.Parse("2006-01-02T15:04", form.StartAt)
	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).SendString("Invalid start date format")
	}

	endAt, err := time.Parse("2006-01-02T15:04", form.EndAt)
	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).SendString("Invalid end date format")
	}

	serviceSportsmen, err := t.service.GetInactiveSportsmenForPeriod(fiberCtx.Context(), startAt, endAt)
	if err != nil {
		return fmt.Errorf("get sportsmen: %w", err)
	}

	sportsmen := models.ConvertFromServiceSportsmen(serviceSportsmen)

	return fiberCtx.Render("queries/inactive_sportsmen", fiber.Map{
		"Sportsmen": sportsmen,
	})
}

func (t *Transport) RenderOrganizerTournamentCountsGetPage(fiberCtx *fiber.Ctx) error {
	return fiberCtx.Render("queries/organizer_tournament_counts", fiber.Map{})
}

func (t *Transport) RenderOrganizerTournamentCountsPostPage(fiberCtx *fiber.Ctx) error {
	var form getOrganizerTournamentCountsForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	startAt, err := time.Parse("2006-01-02T15:04", form.StartAt)
	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).SendString("Invalid start date format")
	}

	endAt, err := time.Parse("2006-01-02T15:04", form.EndAt)
	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).SendString("Invalid end date format")
	}

	serviceOrganizers, err := t.service.GetOrganizerTournamentCountsForPeriod(fiberCtx.Context(), startAt, endAt)
	if err != nil {
		return fmt.Errorf("get tournaments: %w", err)
	}

	organizer := models.ConvertFromServiceOrganizerTournamentCounts(serviceOrganizers)

	return fiberCtx.Render("queries/organizer_tournament_counts", fiber.Map{
		"OrganizerTournamentCounts": organizer,
	})
}

func (t *Transport) RenderClubActiveSportsmenCountsGetPage(fiberCtx *fiber.Ctx) error {
	return fiberCtx.Render("queries/club_active_sportsmen_counts", fiber.Map{})
}

func (t *Transport) RenderClubActiveSportsmenCountsPostPage(fiberCtx *fiber.Ctx) error {
	var form getClubActiveSportsmenCountsForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	startAt, err := time.Parse("2006-01-02T15:04", form.StartAt)
	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).SendString("Invalid start date format")
	}

	endAt, err := time.Parse("2006-01-02T15:04", form.EndAt)
	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).SendString("Invalid end date format")
	}

	serviceClubs, err := t.service.GetClubActiveSportsmenCountsForPeriod(fiberCtx.Context(), startAt, endAt)
	if err != nil {
		return fmt.Errorf("get tournaments: %w", err)
	}

	clubs := models.ConvertFromServiceClubSportsmenCounts(serviceClubs)

	return fiberCtx.Render("queries/club_active_sportsmen_counts", fiber.Map{
		"ClubSportsmenCounts": clubs,
	})
}

func (t *Transport) RenderPlacesWithTournamentDatesGetPage(fiberCtx *fiber.Ctx) error {
	return fiberCtx.Render("queries/places_with_tournament_dates", fiber.Map{})
}

func (t *Transport) RenderPlacesWithTournamentDatesPostPage(fiberCtx *fiber.Ctx) error {
	var form getPlacesWithTournamentDatesForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	startAt, err := time.Parse("2006-01-02T15:04", form.StartAt)
	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).SendString("Invalid start date format")
	}

	endAt, err := time.Parse("2006-01-02T15:04", form.EndAt)
	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).SendString("Invalid end date format")
	}

	servicePlaces, err := t.service.GetPlacesWithTournamentDatesForPeriod(fiberCtx.Context(), startAt, endAt)
	if err != nil {
		return fmt.Errorf("get places: %w", err)
	}

	places := models.ConvertFromServicePlacesWithTournamentDates(servicePlaces)

	return fiberCtx.Render("queries/places_with_tournament_dates", fiber.Map{
		"Places": places,
	})
}

func (t *Transport) RenderTournamentsForPeriodGetPage(fiberCtx *fiber.Ctx) error {
	serviceOrganizers, err := t.service.GetOrganizers(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	organizers := models.ConvertFromServiceOrganizers(serviceOrganizers)

	return fiberCtx.Render("queries/tournaments_for_period", fiber.Map{
		"Organizers": organizers,
	})
}

func (t *Transport) RenderTournamentsForPeriodPostPage(fiberCtx *fiber.Ctx) error {
	serviceOrganizers, err := t.service.GetOrganizers(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	organizers := models.ConvertFromServiceOrganizers(serviceOrganizers)

	var form getTournamentsForPeriodForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	var organizerID *int64
	if form.OrganizerID > 0 {
		organizerID = &form.OrganizerID
	}

	startAt, err := time.Parse("2006-01-02T15:04", form.StartAt)
	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).SendString("Invalid start date format")
	}

	endAt, err := time.Parse("2006-01-02T15:04", form.EndAt)
	if err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).SendString("Invalid end date format")
	}

	serviceTournaments, err := t.service.GetTournamentsForPeriod(fiberCtx.Context(), startAt, endAt, organizerID)
	if err != nil {
		return fmt.Errorf("get tournaments: %w", err)
	}

	tournaments := models.ConvertFromServiceTournaments(serviceTournaments)

	return fiberCtx.Render("queries/tournaments_for_period", fiber.Map{
		"Organizers":  organizers,
		"Tournaments": tournaments,
	})
}

func (t *Transport) RenderSportSportsmenGetPage(fiberCtx *fiber.Ctx) error {
	serviceSports, err := t.service.GetSports(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	sports := models.ConvertFromServiceSports(serviceSports)

	return fiberCtx.Render("queries/sport_sportsmen", fiber.Map{
		"Sports": sports,
	})
}

func (t *Transport) RenderSportSportsmenPostPage(fiberCtx *fiber.Ctx) error {
	serviceSports, err := t.service.GetSports(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	sports := models.ConvertFromServiceSports(serviceSports)

	var form getSportSportsmenForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	var minRank *int16
	if form.MinRank > 0 {
		minRank = &form.MinRank
	}

	serviceSportsmen, err := t.service.GetSportsmenBySportID(fiberCtx.Context(), form.SportID, minRank)
	if err != nil {
		return fmt.Errorf("get trainers: %w", err)
	}

	rankedSportsmen := models.ConvertFromServiceRankedSportsmen(serviceSportsmen)

	return fiberCtx.Render("queries/sport_sportsmen", fiber.Map{
		"Sports":          sports,
		"RankedSportsmen": rankedSportsmen,
	})
}

func (t *Transport) RenderTrainerSportsmenGetPage(fiberCtx *fiber.Ctx) error {
	serviceTrainers, err := t.service.GetTrainers(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	trainers := models.ConvertFromServiceTrainers(serviceTrainers)

	return fiberCtx.Render("queries/trainer_sportsmen", fiber.Map{
		"Trainers": trainers,
	})
}

func (t *Transport) RenderTrainerSportsmenPostPage(fiberCtx *fiber.Ctx) error {
	serviceTrainers, err := t.service.GetTrainers(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	trainers := models.ConvertFromServiceTrainers(serviceTrainers)

	var form getTrainerSportsmenForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	var minRank *int16
	if form.MinRank > 0 {
		minRank = &form.MinRank
	}

	serviceSportsmen, err := t.service.GetSportsmenBySportID(fiberCtx.Context(), form.TrainerID, minRank)
	if err != nil {
		return fmt.Errorf("get trainers: %w", err)
	}

	rankedSportsmen := models.ConvertFromServiceRankedSportsmen(serviceSportsmen)

	return fiberCtx.Render("queries/trainer_sportsmen", fiber.Map{
		"Trainers":        trainers,
		"RankedSportsmen": rankedSportsmen,
	})
}

func (t *Transport) RenderSportTrainersGetPage(fiberCtx *fiber.Ctx) error {
	serviceSports, err := t.service.GetSports(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	sports := models.ConvertFromServiceSports(serviceSports)

	return fiberCtx.Render("queries/sport_trainers", fiber.Map{
		"Sports": sports,
	})
}

func (t *Transport) RenderSportTrainersPostPage(fiberCtx *fiber.Ctx) error {
	serviceSports, err := t.service.GetSports(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	sports := models.ConvertFromServiceSports(serviceSports)

	var form getSportTrainersForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	serviceTrainers, err := t.service.GetTrainersBySportID(fiberCtx.Context(), form.SportID)
	if err != nil {
		return fmt.Errorf("get trainers: %w", err)
	}

	trainers := models.ConvertFromServiceTrainers(serviceTrainers)

	return fiberCtx.Render("queries/sport_trainers", fiber.Map{
		"Sports":   sports,
		"Trainers": trainers,
	})
}

func (t *Transport) RenderSportsmenInvolvedInSeveralSportsPage(fiberCtx *fiber.Ctx) error {
	serviceSportsmen, err := t.service.GetSportsmenInvolvedInSeveralSports(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	sportsmen := models.ConvertFromServiceSportsmen(serviceSportsmen)

	return fiberCtx.Render("queries/sportsmen_involved_in_several_sports", fiber.Map{
		"Sportsmen": sportsmen,
	})
}

func (t *Transport) UpdateSport(fiberCtx *fiber.Ctx) error {
	sportID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sport ID: %w", err)
	}

	var form updateSportForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	err = t.service.UpdateSportByID(fiberCtx.Context(), servicedto.UpdateSportByIDRequest{
		ID:   sportID,
		Name: form.Name,
	})
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect(fmt.Sprintf("/sports/%d", sportID), fiber.StatusFound)
}

func (t *Transport) UpdateSportsman(fiberCtx *fiber.Ctx) error {
	sportsmanID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sportsman ID: %w", err)
	}

	var form updateSportsmanForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	birthDate, err := time.Parse("2006-01-02", form.BirthDate)
	if err != nil {
		return fmt.Errorf("parse birth date: %w", err)
	}

	heightCm, err := strconv.ParseUint(form.HeightCm, 10, 16)
	if err != nil {
		return fmt.Errorf("parse height cm: %w", err)
	}

	weightKg, err := strconv.ParseFloat(form.WeightKg, 64)
	if err != nil {
		return fmt.Errorf("parse weight kg: %w", err)
	}

	err = t.service.UpdateSportsmanByID(fiberCtx.Context(), servicedto.UpdateSportsmanByIDRequest{
		ID:        sportsmanID,
		Name:      form.Name,
		BirthDate: birthDate,
		HeightCm:  uint16(heightCm),
		WeightKg:  weightKg,
		ClubID:    form.ClubID,
		SportIDs:  form.SportIDs,
	})
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect(fmt.Sprintf("/sportsmen/%d", sportsmanID), fiber.StatusFound)
}
