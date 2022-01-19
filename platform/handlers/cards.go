package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// GenerateCard godoc
// @Summary Generate card preview
// @Description Generate card preview
// @Tags cards
// @Produce json
// @Success 200 {object} schema.NewCardResponse
// @Router /new [post]
func (h *Handler) GenerateCard(ctx *fiber.Ctx) error {
	cardOptions, err := h.Params.CardOptionsPayload(ctx)
	if err != nil {
		return err
	}

	card, err := h.CardService.GenerateCard(cardOptions)
	if err != nil {
		return err
	}

	return ctx.JSON(card)
}

// CreateCard godoc
// @Summary Create card
// @Description Create card
// @Tags cards
// @Produce json
// @Success 201 {object} schema.CardResponse
// @Router / [post]
func (h *Handler) CreateCard(ctx *fiber.Ctx) error {
	newCardRequest, err := h.Params.CreateCardPayload(ctx)
	if err != nil {
		return err
	}

	userId, err := h.Params.GetUserIdFromLocals(ctx)
	if err != nil {
		return err
	}

	card, err := h.CardService.Create(*userId, newCardRequest)
	if err != nil {
		return err
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(card)
}
