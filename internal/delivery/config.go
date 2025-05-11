package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"

	"db2sem/internal/utils/duration"
)

type Config struct {
	Server ServerConfig `validate:"required"`
	Cors   CorsConfig   `validate:"required"`
	Serve  ServeConfig  `validate:"required"`
}

type CorsConfig struct {
	AllowOrigins     string `validate:"required"`
	AllowHeaders     string `validate:"required"`
	ExposeHeaders    string
	AllowCredentials bool
}

func (cfg CorsConfig) convertToForeign() cors.Config {
	return cors.Config{
		AllowOrigins:     cfg.AllowOrigins,
		AllowHeaders:     cfg.AllowHeaders,
		ExposeHeaders:    cfg.ExposeHeaders,
		AllowCredentials: cfg.AllowCredentials,
	}
}

type ServeConfig struct {
	Address string `validate:"required"`
}

type ServerConfig struct {
	ProxyHeader           string           `validate:"required" default:"X-Real-IP"`
	ReadTimeoutSeconds    duration.Seconds `validate:"required"`
	WriteTimeoutSeconds   duration.Seconds `validate:"required"`
	DisableStartupMessage bool
	ViewsDir              string `validate:"required" default:"./views"`
}

func (cfg ServerConfig) convertToForeign() fiber.Config {
	return fiber.Config{
		ProxyHeader:           cfg.ProxyHeader,
		ReadTimeout:           cfg.ReadTimeoutSeconds.Duration,
		WriteTimeout:          cfg.WriteTimeoutSeconds.Duration,
		DisableStartupMessage: cfg.DisableStartupMessage,
		Views:                 html.New(cfg.ViewsDir, ".html"),
	}
}
