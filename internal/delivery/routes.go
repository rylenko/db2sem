package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func registerRoutes(app *fiber.App, transport transport) {
	app.Use(recover.New())

	app.Get("/", transport.RenderIndexPage)

	queries := app.Group("queries")
	{
		queries.Get("/", transport.RenderQueriesPage)

		queries.Get("/arenas", transport.RenderArenasPage)
		queries.Get("/stadiums", transport.RenderStadiumsPage)
		queries.Get("/courts", transport.RenderCourtsPage)
		queries.Get("/gyms", transport.RenderGymsPage)

		queries.Get("/sportsmen-involved-in-several-sports", transport.RenderSportsmenInvolvedInSeveralSportsPage)

		queries.Get("/sportsman-trainers", transport.RenderSportsmanTrainersGetPage)
		queries.Post("/sportsman-trainers", transport.RenderSportsmanTrainersPostPage)

		queries.Get("/tournament-prize-winners", transport.RenderTournamentPrizeWinnersGetPage)
		queries.Post("/tournament-prize-winners", transport.RenderTournamentPrizeWinnersPostPage)

		queries.Get("/sport-trainers", transport.RenderSportTrainersGetPage)
		queries.Post("/sport-trainers", transport.RenderSportTrainersPostPage)

		queries.Get("/sport-sportsmen", transport.RenderSportSportsmenGetPage)
		queries.Post("/sport-sportsmen", transport.RenderSportSportsmenPostPage)

		queries.Get("/trainer-sportsmen", transport.RenderTrainerSportsmenGetPage)
		queries.Post("/trainer-sportsmen", transport.RenderTrainerSportsmenPostPage)

		queries.Get("/tournaments-for-period", transport.RenderTournamentsForPeriodGetPage)
		queries.Post("/tournaments-for-period", transport.RenderTournamentsForPeriodPostPage)

		queries.Get("/place-tournaments", transport.RenderPlaceTournamentsGetPage)
		queries.Post("/place-tournaments", transport.RenderPlaceTournamentsPostPage)

		queries.Get("/club-active-sportsmen-counts", transport.RenderClubActiveSportsmenCountsGetPage)
		queries.Post("/club-active-sportsmen-counts", transport.RenderClubActiveSportsmenCountsPostPage)

		queries.Get("/inactive-sportsmen", transport.RenderInactiveSportsmenGetPage)
		queries.Post("/inactive-sportsmen", transport.RenderInactiveSportsmenPostPage)

		queries.Get("/organizer-tournament-counts", transport.RenderOrganizerTournamentCountsGetPage)
		queries.Post("/organizer-tournament-counts", transport.RenderOrganizerTournamentCountsPostPage)

		queries.Get("/places-with-tournament-dates", transport.RenderPlacesWithTournamentDatesGetPage)
		queries.Post("/places-with-tournament-dates", transport.RenderPlacesWithTournamentDatesPostPage)
	}

	sports := app.Group("sports")
	{
		sports.Get("/", transport.RenderSportsPage)
		sports.Post("/", transport.CreateSport)
		sports.Get("/:id", transport.RenderSportPage)
		sports.Post("/:id/update", transport.UpdateSport)
		sports.Post("/:id/delete", transport.DeleteSport)
	}

	sportsmen := app.Group("sportsmen")
	{
		sportsmen.Get("/", transport.RenderSportsmenPage)
		sportsmen.Post("/", transport.CreateSportsman)
		sportsmen.Get("/:id", transport.RenderSportsmanPage)
		sportsmen.Post("/:id/update", transport.UpdateSportsman)
		sportsmen.Post("/:id/delete", transport.DeleteSportsman)
	}

	places := app.Group("places")
	{
		places.Post("/:id/delete", transport.DeletePlace)
	}

	arenas := app.Group("arenas")
	{
		arenas.Get("/:id", transport.RenderArenaPage)
		arenas.Post("/:id/update", transport.UpdateArena)
	}

	stadiums := app.Group("stadiums")
	{
		stadiums.Get("/:id", transport.RenderStadiumPage)
		stadiums.Post("/:id/update", transport.UpdateStadium)
	}

	courts := app.Group("courts")
	{
		courts.Get("/:id", transport.RenderCourtPage)
		courts.Post("/:id/update", transport.UpdateCourt)
	}

	gyms := app.Group("gyms")
	{
		gyms.Get("/:id", transport.RenderGymPage)
		gyms.Post("/:id/update", transport.UpdateGym)
	}
}
