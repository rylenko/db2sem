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
		queries.Get("/sportsmen-involved-in-several-sports", transport.RenderSportsmenInvolvedInSeveralSportsPage)

		queries.Get("/sportsman-trainers", transport.RenderSportsmanTrainersGetPage)
		queries.Post("/sportsman-trainers", transport.RenderSportsmanTrainersPostPage)

		queries.Get("/tournament-prize-winners", transport.RenderTournamentPrizeWinnersGetPage)
		queries.Post("/tournament-prize-winners", transport.RenderTournamentPrizeWinnersPostPage)

		queries.Get("/sport-trainers", transport.RenderSportTrainersGetPage)
		queries.Post("/sport-trainers", transport.RenderSportTrainersPostPage)
	}

	sportsmen := app.Group("sportsmen")
	{
		sportsmen.Get("/", transport.RenderSportsmenPage)
		sportsmen.Post("/", transport.CreateSportsman)
		sportsmen.Get("/:id", transport.RenderSportsmanPage)
		sportsmen.Post("/:id/update", transport.UpdateSportsman)
		sportsmen.Post("/:id/delete", transport.DeleteSportsman)
	}
}
