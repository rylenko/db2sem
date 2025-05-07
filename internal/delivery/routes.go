package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func registerRoutes(app *fiber.App, _ transport) {
	app.Use(recover.New())
}
