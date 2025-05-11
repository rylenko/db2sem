package transport

import (
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

func (t *Transport) RenderIndexPage(fiberCtx *fiber.Ctx) error {
	return fiberCtx.Render("index", fiber.Map{})
}
