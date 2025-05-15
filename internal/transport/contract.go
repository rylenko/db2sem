package transport

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"db2sem/internal/domain"
	servicedto "db2sem/internal/service/dto"
)

type requestReader interface {
	ReadAndValidateFiberBody(fiberCtx *fiber.Ctx, request any) error
}

type service interface {
	CreateSportsman(ctx context.Context, req servicedto.CreateSportsmanRequest) error
	DeleteSportsmanByID(ctx context.Context, sportsmanID int64) error
	GetClubs(ctx context.Context) ([]domain.Club, error)
	GetSportsmanByID(ctx context.Context, sportsmanID int64) (*domain.Sportsman, error)
	GetSportsmen(ctx context.Context) ([]domain.Sportsman, error)
	GetSportsmenInvolvedInSeveralSports(ctx context.Context) ([]domain.Sportsman, error)
	GetSports(ctx context.Context) ([]domain.Sport, error)
	GetTrainersBySportsmanID(ctx context.Context, sportsmanID int64) ([]domain.Trainer, error)
	UpdateSportsmanByID(ctx context.Context, req servicedto.UpdateSportsmanByIDRequest) error
}
