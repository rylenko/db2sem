package delivery

import "github.com/gofiber/fiber/v2"

type transport interface {
	RenderIndexPage(fiberCtx *fiber.Ctx) error
}
