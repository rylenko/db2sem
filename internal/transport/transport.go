package transport

import (
	"fmt"
	"strconv"
	"time"

	servicedto "db2sem/internal/service/dto"

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

func (t *Transport) DeleteSportsman(fiberCtx *fiber.Ctx) error {
	sportsmanID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sportsman ID: %w", err)
	}

	if err := t.service.DeleteSportsmanByID(fiberCtx.Context(), sportsmanID); err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect("/sportsmen/", fiber.StatusFound)
}

func (t *Transport) RenderIndexPage(fiberCtx *fiber.Ctx) error {
	return fiberCtx.Render("index", fiber.Map{})
}

func (t *Transport) RenderQueriesPage(fiberCtx *fiber.Ctx) error {
	return fiberCtx.Render("queries/index", fiber.Map{})
}

func (t *Transport) RenderSportsmanPage(fiberCtx *fiber.Ctx) error {
	sportsmanID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sportsman ID: %w", err)
	}

	serviceSportsman, err := t.service.GetSportsmanByID(fiberCtx.Context(), sportsmanID)
	if err != nil {
		return fmt.Errorf("get sportsman: %w", err)
	}

	if serviceSportsman == nil {
		return fiberCtx.Status(fiber.StatusNotFound).SendString("Sportsman not found")
	}

	sportsman := convertFromServiceSportsman(*serviceSportsman)

	sportNames, err := t.service.GetSportNames(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("get sport names: %w", err)
	}

	return fiberCtx.Render("sportsman", fiber.Map{
		"Sportsman":  sportsman,
		"SportNames": sportNames,
	})
}

func (t *Transport) RenderSportsmenPage(fiberCtx *fiber.Ctx) error {
	serviceSportsmen, err := t.service.GetSportsmen(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	sportsmen := convertFromServiceSportsmen(serviceSportsmen)

	return fiberCtx.Render("sportsmen", fiber.Map{
		"Sportsmen": sportsmen,
	})
}

func (t *Transport) RenderSportsmenInvolvedInSeveralSportsPage(fiberCtx *fiber.Ctx) error {
	serviceSportsmen, err := t.service.GetSportsmenInvolvedInSeveralSports(fiberCtx.Context())
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	sportsmen := convertFromServiceSportsmen(serviceSportsmen)

	return fiberCtx.Render("queries/sportsmen_involved_in_several_sports", fiber.Map{
		"Sportsmen": sportsmen,
	})
}

func (t *Transport) UpdateSportsman(fiberCtx *fiber.Ctx) error {
	sportsmanID, err := strconv.ParseInt(fiberCtx.Params("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse sportsman ID: %w", err)
	}

	var form updateSportsmanForm
	if err := fiberCtx.BodyParser(&form); err != nil {
		return fmt.Errorf("parse body: %w", err)
	}

	birthDate, err := time.Parse("2006-01-02", form.BirthDate)
	if err != nil {
		return fmt.Errorf("parse birth date: %w", err)
	}

	heightCm, err := strconv.ParseUint(form.HeightCm, 10, 16)
	if err != nil {
		return fmt.Errorf("parse height cm: %w", err)
	}

	weightKg, err := strconv.ParseFloat(form.WeightKg, 64)
	if err != nil {
		return fmt.Errorf("parse weight kg: %w", err)
	}

	fmt.Println(form.SportNames)

	err = t.service.UpdateSportsmanByID(fiberCtx.Context(), servicedto.UpdateSportsmanByIDRequest{
		ID:         sportsmanID,
		Name:       form.Name,
		BirthDate:  birthDate,
		HeightCm:   uint16(heightCm),
		WeightKg:   weightKg,
		SportNames: form.SportNames,
	})
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}

	return fiberCtx.Redirect(fmt.Sprintf("/sportsmen/%d", sportsmanID), fiber.StatusFound)
}
