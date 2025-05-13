package delivery

import "github.com/gofiber/fiber/v2"

type transport interface {
	DeleteSportsman(fiberCtx *fiber.Ctx) error
	RenderIndexPage(fiberCtx *fiber.Ctx) error
	RenderQueriesPage(fibercCtx *fiber.Ctx) error
	RenderSportsmanPage(fibercCtx *fiber.Ctx) error
	RenderSportsmenPage(fibercCtx *fiber.Ctx) error
	RenderSportsmenInvolvedInSeveralSportsPage(fibercCtx *fiber.Ctx) error
	UpdateSportsman(fiberCtx *fiber.Ctx) error
}
