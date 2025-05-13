package transport

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"db2sem/internal/domain"
)

type requestReader interface {
	ReadAndValidateFiberBody(fiberCtx *fiber.Ctx, request any) error
}

type service interface {
	DeleteSportsmanByID(ctx context.Context, sportsmanID int64) error
	GetSportsmanByID(ctx context.Context, sportsmanID int64) (*domain.Sportsman, error)
	GetSportsmen(ctx context.Context) ([]domain.Sportsman, error)
	GetSportsmenInvolvedInSeveralSports(ctx context.Context) ([]domain.Sportsman, error)
	GetSportNames(ctx context.Context) ([]string, error)
}
