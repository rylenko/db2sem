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

func (t *Transport) CreateOrganizer(fiberCtx *fiber.Ctx) error {
	var form createOrganizerForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	var location *string
	if form.Location != "" {
		location = &form.Location
	}

	err := t.service.CreateOrganizer(fiberCtx.Context(), form.Name, location)
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect("/organizers/", fiber.StatusFound)
}

func (t *Transport) CreateTrainer(fiberCtx *fiber.Ctx) error {
	var form createTrainerForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	err := t.service.CreateTrainer(fiberCtx.Context(), form.Name)
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect("/trainers/", fiber.StatusFound)
}

func (t *Transport) CreateClub(fiberCtx *fiber.Ctx) error {
	var form createClubForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	err := t.service.CreateClub(fiberCtx.Context(), form.Name)
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect("/clubs/", fiber.StatusFound)
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

func (t *Transport) CreateStadium(fiberCtx *fiber.Ctx) error {
	var form createStadiumForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	err := t.service.CreateStadium(fiberCtx.Context(), servicedto.CreateStadiumRequest{
		Name:          form.Name,
		Location:      form.Location,
		WidthCm:       form.WidthCm,
		LengthCm:      form.LengthCm,
		MaxSpectators: form.MaxSpectators,
		IsOutdoor:     form.IsOutdoor,
		Coating:       form.Coating,
	})
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect("/queries/stadiums/", fiber.StatusFound)
}

func (t *Transport) CreateArena(fiberCtx *fiber.Ctx) error {
	var form createArenaForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	err := t.service.CreateArena(fiberCtx.Context(), servicedto.CreateArenaRequest{
		Name:              form.Name,
		Location:          form.Location,
		RefereesCount:     form.RefereesCount,
		TreadmillLengthCm: form.TreadmillLengthCm,
	})
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect("/queries/arenas/", fiber.StatusFound)
}

func (t *Transport) CreateTournament(fiberCtx *fiber.Ctx) error {
	var form createTournamentForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	startAt, err := time.Parse("2006-01-02T15:04", form.StartAt)
	if err != nil {
		return fmt.Errorf("parse birth date: %w", err)
	}

	err = t.service.CreateTournament(fiberCtx.Context(), servicedto.CreateTournamentRequest{
		OrganizerID: form.OrganizerID,
		PlaceID:     form.PlaceID,
		StartAt:     startAt,
		SportIDs:    form.SportIDs,
	})
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect("/tournaments/", fiber.StatusFound)
}

func (t *Transport) CreateParticipation(fiberCtx *fiber.Ctx) error {
	var form createParticipationForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	var results *string
	if form.Results != "" {
		results = &form.Results
	}

	err := t.service.CreateParticipation(fiberCtx.Context(), servicedto.CreateParticipationRequest{
		TournamentSportID: form.TournamentSportID,
		SportsmanID:       form.SportsmanID,
		Rank:              form.Rank,
		Results:           results,
	})
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect("/participations/", fiber.StatusFound)
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

func (t *Transport) DeletePlace(fiberCtx *fiber.Ctx) error {
	placeID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse place ID: %w", err)
	}

	if err := t.service.DeletePlaceByID(fiberCtx.Context(), placeID); err != nil {
		return fmt.Errorf("service: %w", err)
	}

	referer := fiberCtx.Get("Referer") // получаем URL страницы с которой пришли
	if referer == "" {
		referer = "/" // fallback, если Referer нет
	}

	return fiberCtx.Redirect(referer, fiber.StatusFound)
}

func (t *Transport) DeleteOrganizer(fiberCtx *fiber.Ctx) error {
	sportID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sport ID: %w", err)
	}

	if err := t.service.DeleteOrganizerByID(fiberCtx.Context(), sportID); err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect("/organizers/", fiber.StatusFound)
}

func (t *Transport) DeleteTrainer(fiberCtx *fiber.Ctx) error {
	sportID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sport ID: %w", err)
	}

	if err := t.service.DeleteTrainerByID(fiberCtx.Context(), sportID); err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect("/trainers/", fiber.StatusFound)
}

func (t *Transport) DeleteClub(fiberCtx *fiber.Ctx) error {
	sportID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sport ID: %w", err)
	}

	if err := t.service.DeleteClubByID(fiberCtx.Context(), sportID); err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect("/clubs/", fiber.StatusFound)
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

func (t *Transport) RenderTournamentsPage(fiberCtx *fiber.Ctx) error {
	serviceTournaments, err := t.service.GetTournaments(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("get clubs: %w", err)
	}

	tournaments := models.ConvertFromServiceTournaments(serviceTournaments)

	serviceOrganizers, err := t.service.GetOrganizers(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("get clubs: %w", err)
	}

	organizers := models.ConvertFromServiceOrganizers(serviceOrganizers)

	servicePlaces, err := t.service.GetPlaces(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("get clubs: %w", err)
	}

	places := models.ConvertFromServicePlaces(servicePlaces)

	serviceSports, err := t.service.GetSports(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("get sport names: %w", err)
	}

	sports := models.ConvertFromServiceSports(serviceSports)

	return fiberCtx.Render("tournaments", fiber.Map{
		"Tournaments": tournaments,
		"Organizers":  organizers,
		"Places":      places,
		"Sports":      sports,
	})
}

func (t *Transport) RenderParticipationsPage(fiberCtx *fiber.Ctx) error {
	participations, err := t.service.GetParticipations(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	tournamentSports, err := t.service.GetTournamentSports(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("get sport names: %w", err)
	}

	serviceSportsmen, err := t.service.GetSportsmen(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("get clubs: %w", err)
	}

	sportsmen := models.ConvertFromServiceSportsmen(serviceSportsmen)

	return fiberCtx.Render("participations", fiber.Map{
		"Sportsmen":        sportsmen,
		"TournamentSports": tournamentSports,
		"Participations":   participations,
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

func (t *Transport) RenderGymPage(fiberCtx *fiber.Ctx) error {
	courtID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sportsman ID: %w", err)
	}

	serviceCourt, err := t.service.GetGymByID(fiberCtx.Context(), courtID)
	if err != nil {
		return fmt.Errorf("get sportsman: %w", err)
	}

	if serviceCourt == nil {
		return fiberCtx.Status(fiber.StatusNotFound).SendString("Arena not found")
	}

	court := models.ConvertFromServiceGym(*serviceCourt)

	return fiberCtx.Render("gym", court)
}

func (t *Transport) RenderCourtPage(fiberCtx *fiber.Ctx) error {
	courtID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sportsman ID: %w", err)
	}

	serviceCourt, err := t.service.GetCourtByID(fiberCtx.Context(), courtID)
	if err != nil {
		return fmt.Errorf("get sportsman: %w", err)
	}

	if serviceCourt == nil {
		return fiberCtx.Status(fiber.StatusNotFound).SendString("Arena not found")
	}

	court := models.ConvertFromServiceCourt(*serviceCourt)

	return fiberCtx.Render("court", court)
}

func (t *Transport) RenderStadiumPage(fiberCtx *fiber.Ctx) error {
	arenaID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sportsman ID: %w", err)
	}

	serviceArena, err := t.service.GetStadiumByID(fiberCtx.Context(), arenaID)
	if err != nil {
		return fmt.Errorf("get sportsman: %w", err)
	}

	if serviceArena == nil {
		return fiberCtx.Status(fiber.StatusNotFound).SendString("Arena not found")
	}

	arena := models.ConvertFromServiceStadium(*serviceArena)

	return fiberCtx.Render("stadium", arena)
}

func (t *Transport) RenderArenaPage(fiberCtx *fiber.Ctx) error {
	arenaID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sportsman ID: %w", err)
	}

	serviceArena, err := t.service.GetArenaByID(fiberCtx.Context(), arenaID)
	if err != nil {
		return fmt.Errorf("get sportsman: %w", err)
	}

	if serviceArena == nil {
		return fiberCtx.Status(fiber.StatusNotFound).SendString("Arena not found")
	}

	arena := models.ConvertFromServiceArena(*serviceArena)

	return fiberCtx.Render("arena", arena)
}

func (t *Transport) RenderOrganizerPage(fiberCtx *fiber.Ctx) error {
	sportID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sportsman ID: %w", err)
	}

	serviceSport, err := t.service.GetOrganizerByID(fiberCtx.Context(), sportID)
	if err != nil {
		return fmt.Errorf("get sportsman: %w", err)
	}

	if serviceSport == nil {
		return fiberCtx.Status(fiber.StatusNotFound).SendString("Sport not found")
	}

	sport := models.ConvertFromServiceOrganizer(*serviceSport)

	return fiberCtx.Render("organizer", fiber.Map{
		"Organizer": sport,
	})
}

func (t *Transport) RenderTrainerPage(fiberCtx *fiber.Ctx) error {
	sportID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sportsman ID: %w", err)
	}

	serviceSport, err := t.service.GetTrainerByID(fiberCtx.Context(), sportID)
	if err != nil {
		return fmt.Errorf("get sportsman: %w", err)
	}

	if serviceSport == nil {
		return fiberCtx.Status(fiber.StatusNotFound).SendString("Sport not found")
	}

	sport := models.ConvertFromServiceTrainer(*serviceSport)

	return fiberCtx.Render("trainer", fiber.Map{
		"Trainer": sport,
	})
}

func (t *Transport) RenderClubPage(fiberCtx *fiber.Ctx) error {
	sportID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sportsman ID: %w", err)
	}

	serviceSport, err := t.service.GetClubByID(fiberCtx.Context(), sportID)
	if err != nil {
		return fmt.Errorf("get sportsman: %w", err)
	}

	if serviceSport == nil {
		return fiberCtx.Status(fiber.StatusNotFound).SendString("Sport not found")
	}

	sport := models.ConvertFromServiceClub(*serviceSport)

	return fiberCtx.Render("club", fiber.Map{
		"Club": sport,
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

func (t *Transport) RenderOrganizersPage(fiberCtx *fiber.Ctx) error {
	serviceSports, err := t.service.GetOrganizers(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	sports := models.ConvertFromServiceOrganizers(serviceSports)

	return fiberCtx.Render("organizers", fiber.Map{
		"Organizers": sports,
	})
}

func (t *Transport) RenderTrainersPage(fiberCtx *fiber.Ctx) error {
	serviceSports, err := t.service.GetTrainers(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	sports := models.ConvertFromServiceTrainers(serviceSports)

	return fiberCtx.Render("trainers", fiber.Map{
		"Trainers": sports,
	})
}

func (t *Transport) RenderClubsPage(fiberCtx *fiber.Ctx) error {
	serviceSports, err := t.service.GetClubs(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	sports := models.ConvertFromServiceClubs(serviceSports)

	return fiberCtx.Render("clubs", fiber.Map{
		"Clubs": sports,
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
		"PlacesWithTournamentDates": places,
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

func (t *Transport) RenderGymsPage(fiberCtx *fiber.Ctx) error {
	var form getGymsForm

	err := t.requestReader.ReadAndValidateFiberQuery(fiberCtx, &form)
	if err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	var req servicedto.GetGymsRequest

	if form.TrainersCount > 0 {
		req.TrainersCount = &form.TrainersCount
	}

	if form.DumbbellsCount > 0 {
		req.DumbbellsCount = &form.DumbbellsCount
	}

	req.HasBathhouse = &form.HasBathhouse

	serviceGyms, err := t.service.GetGyms(fiberCtx.Context(), req)
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	gyms := models.ConvertFromServiceGyms(serviceGyms)

	return fiberCtx.Render("queries/gyms", fiber.Map{
		"Gyms": gyms,
	})
}

func (t *Transport) RenderCourtsPage(fiberCtx *fiber.Ctx) error {
	var form getCourtsForm

	err := t.requestReader.ReadAndValidateFiberQuery(fiberCtx, &form)
	if err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	var req servicedto.GetCourtsRequest

	if form.WidthCm > 0 {
		req.WidthCm = &form.WidthCm
	}

	if form.LengthCm > 0 {
		req.LengthCm = &form.LengthCm
	}

	req.IsOutdoor = &form.IsOutdoor

	serviceCourts, err := t.service.GetCourts(fiberCtx.Context(), req)
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	courts := models.ConvertFromServiceCourts(serviceCourts)

	return fiberCtx.Render("queries/courts", fiber.Map{
		"Courts": courts,
	})
}

func (t *Transport) RenderStadiumsPage(fiberCtx *fiber.Ctx) error {
	var form getStadiumsForm

	err := t.requestReader.ReadAndValidateFiberQuery(fiberCtx, &form)
	if err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	var req servicedto.GetStadiumsRequest

	if form.WidthCm > 0 {
		req.WidthCm = &form.WidthCm
	}

	if form.LengthCm > 0 {
		req.LengthCm = &form.LengthCm
	}

	if form.MaxSpectators > 0 {
		req.MaxSpectators = &form.MaxSpectators
	}

	req.IsOutdoor = &form.IsOutdoor

	if form.Coating != "" {
		req.Coating = &form.Coating
	}

	serviceStadiums, err := t.service.GetStadiums(fiberCtx.Context(), req)
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	stadiums := models.ConvertFromServiceStadiums(serviceStadiums)

	return fiberCtx.Render("queries/stadiums", fiber.Map{
		"Stadiums": stadiums,
	})
}

func (t *Transport) RenderArenasPage(fiberCtx *fiber.Ctx) error {
	var form getArenasForm
	if err := t.requestReader.ReadAndValidateFiberQuery(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	var req servicedto.GetArenasRequest

	if form.RefereesCount > 0 {
		req.RefereesCount = &form.RefereesCount
	}

	if form.TreadmillLengthCm > 0 {
		req.TreadmillLengthCm = &form.TreadmillLengthCm
	}

	serviceArenas, err := t.service.GetArenas(fiberCtx.Context(), req)
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	arenas := models.ConvertFromServiceArenas(serviceArenas)

	return fiberCtx.Render("queries/arenas", fiber.Map{
		"Arenas": arenas,
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

func (t *Transport) UpdateGym(fiberCtx *fiber.Ctx) error {
	stadiumID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sport ID: %w", err)
	}

	var form updateGymForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	err = t.service.UpdateGymByID(fiberCtx.Context(), servicedto.UpdateGymByIDRequest{
		ID:             stadiumID,
		Name:           form.Name,
		Location:       form.Location,
		TrainersCount:  form.TrainersCount,
		DumbbellsCount: form.DumbbellsCount,
		HasBathhouse:   form.HasBathhouse,
	})
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect(fmt.Sprintf("/gyms/%d", stadiumID), fiber.StatusFound)
}

func (t *Transport) UpdateCourt(fiberCtx *fiber.Ctx) error {
	stadiumID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sport ID: %w", err)
	}

	var form updateCourtForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	err = t.service.UpdateCourtByID(fiberCtx.Context(), servicedto.UpdateCourtByIDRequest{
		ID:        stadiumID,
		Name:      form.Name,
		Location:  form.Location,
		WidthCm:   form.WidthCm,
		LengthCm:  form.LengthCm,
		IsOutdoor: form.IsOutdoor,
	})
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect(fmt.Sprintf("/courts/%d", stadiumID), fiber.StatusFound)
}

func (t *Transport) UpdateStadium(fiberCtx *fiber.Ctx) error {
	stadiumID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sport ID: %w", err)
	}

	var form updateStadiumForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	err = t.service.UpdateStadiumByID(fiberCtx.Context(), servicedto.UpdateStadiumByIDRequest{
		ID:            stadiumID,
		Name:          form.Name,
		Location:      form.Location,
		WidthCm:       form.WidthCm,
		LengthCm:      form.LengthCm,
		MaxSpectators: form.MaxSpectators,
		IsOutdoor:     form.IsOutdoor,
		Coating:       form.Coating,
	})
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect(fmt.Sprintf("/stadiums/%d", stadiumID), fiber.StatusFound)
}

func (t *Transport) UpdateArena(fiberCtx *fiber.Ctx) error {
	arenaID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sport ID: %w", err)
	}

	var form updateArenaForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	err = t.service.UpdateArenaByID(fiberCtx.Context(), servicedto.UpdateArenaByIDRequest{
		ID:                arenaID,
		Name:              form.Name,
		Location:          form.Location,
		RefereesCount:     form.RefereesCount,
		TreadmillLengthCm: form.TreadmillLengthCm,
	})
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect(fmt.Sprintf("/arenas/%d", arenaID), fiber.StatusFound)
}

func (t *Transport) UpdateOrganizer(fiberCtx *fiber.Ctx) error {
	sportID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sport ID: %w", err)
	}

	var form updateOrganizerForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	var location *string
	if form.Location != "" {
		location = &form.Location
	}

	err = t.service.UpdateOrganizerByID(fiberCtx.Context(), servicedto.UpdateOrganizerByIDRequest{
		ID:       sportID,
		Name:     form.Name,
		Location: location,
	})
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect(fmt.Sprintf("/organizers/%d", sportID), fiber.StatusFound)
}

func (t *Transport) UpdateTrainer(fiberCtx *fiber.Ctx) error {
	sportID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sport ID: %w", err)
	}

	var form updateTrainerForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	err = t.service.UpdateTrainerByID(fiberCtx.Context(), servicedto.UpdateTrainerByIDRequest{
		ID:   sportID,
		Name: form.Name,
	})
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect(fmt.Sprintf("/trainers/%d", sportID), fiber.StatusFound)
}

func (t *Transport) UpdateClub(fiberCtx *fiber.Ctx) error {
	sportID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sport ID: %w", err)
	}

	var form updateClubForm
	if err := t.requestReader.ReadAndValidateFiberBody(fiberCtx, &form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	err = t.service.UpdateClubByID(fiberCtx.Context(), servicedto.UpdateClubByIDRequest{
		ID:   sportID,
		Name: form.Name,
	})
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect(fmt.Sprintf("/clubs/%d", sportID), fiber.StatusFound)
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
