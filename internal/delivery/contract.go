package delivery

import "github.com/gofiber/fiber/v2"

type transport interface {
	CreateSport(fiberCtx *fiber.Ctx) error
	CreateSportsman(fiberCtx *fiber.Ctx) error
	DeleteSport(fiberCtx *fiber.Ctx) error
	DeleteSportsman(fiberCtx *fiber.Ctx) error
	RenderIndexPage(fiberCtx *fiber.Ctx) error
	RenderQueriesPage(fibercCtx *fiber.Ctx) error
	RenderArenasPage(fiberCtx *fiber.Ctx) error
	RenderStadiumsPage(fiberCtx *fiber.Ctx) error
	RenderCourtsPage(fiberCtx *fiber.Ctx) error
	RenderPlaceTournamentsGetPage(fibercCtx *fiber.Ctx) error
	RenderPlaceTournamentsPostPage(fibercCtx *fiber.Ctx) error
	RenderSportsmanPage(fibercCtx *fiber.Ctx) error
	RenderPlacesWithTournamentDatesGetPage(fiberCtx *fiber.Ctx) error
	RenderPlacesWithTournamentDatesPostPage(fiberCtx *fiber.Ctx) error
	RenderSportsmanTrainersGetPage(fiberCtx *fiber.Ctx) error
	RenderSportsmanTrainersPostPage(fiberCtx *fiber.Ctx) error
	RenderSportPage(fiberCtx *fiber.Ctx) error
	RenderSportsPage(fiberCtx *fiber.Ctx) error
	RenderSportTrainersGetPage(fiberCtx *fiber.Ctx) error
	RenderSportTrainersPostPage(fiberCtx *fiber.Ctx) error
	RenderClubActiveSportsmenCountsGetPage(fiberCtx *fiber.Ctx) error
	RenderClubActiveSportsmenCountsPostPage(fiberCtx *fiber.Ctx) error
	RenderOrganizerTournamentCountsGetPage(fiberCtx *fiber.Ctx) error
	RenderOrganizerTournamentCountsPostPage(fiberCtx *fiber.Ctx) error
	RenderInactiveSportsmenGetPage(fiberCtx *fiber.Ctx) error
	RenderInactiveSportsmenPostPage(fiberCtx *fiber.Ctx) error
	RenderTournamentPrizeWinnersGetPage(fiberCtx *fiber.Ctx) error
	RenderTournamentPrizeWinnersPostPage(fiberCtx *fiber.Ctx) error
	RenderTournamentsForPeriodGetPage(fiberCtx *fiber.Ctx) error
	RenderTournamentsForPeriodPostPage(fiberCtx *fiber.Ctx) error
	RenderTrainerSportsmenGetPage(fibercCtx *fiber.Ctx) error
	RenderTrainerSportsmenPostPage(fibercCtx *fiber.Ctx) error
	RenderSportSportsmenGetPage(fibercCtx *fiber.Ctx) error
	RenderSportSportsmenPostPage(fibercCtx *fiber.Ctx) error
	RenderSportsmenPage(fibercCtx *fiber.Ctx) error
	RenderSportsmenInvolvedInSeveralSportsPage(fibercCtx *fiber.Ctx) error
	UpdateSport(fiberCtx *fiber.Ctx) error
	UpdateSportsman(fiberCtx *fiber.Ctx) error
}
