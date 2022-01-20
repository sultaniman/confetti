package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imanhodjaev/confetti/platform/http"
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

	card, err := h.CardService.Generate(cardOptions)
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

// GetCard godoc
// @Summary Get card by id
// @Description Get card by id
// @Tags cards
// @Produce json
// @Success 200 {object} schema.CardResponse
// @Router / [get]
func (h *Handler) GetCard(ctx *fiber.Ctx) error {
	cardId, err := h.Params.GetUUIDParam(ctx, "card_id")
	if err != nil {
		return err
	}

	userId, err := h.Params.GetUserIdFromLocals(ctx)
	if err != nil {
		return err
	}

	card, err := h.CardService.Get(*cardId)
	if err != nil {
		return err
	}

	if card.UserId.String() != userId.String() {
		return http.NotFoundError("Card not found")
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(card)
}

// UpdateCard godoc
// @Summary Update card
// @Description Update card
// @Tags cards
// @Produce json
// @Success 204 {string} nil update succeeded
// @Router / [put]
func (h *Handler) UpdateCard(ctx *fiber.Ctx) error {
	updateCardRequest, err := h.Params.UpdateCardPayload(ctx)
	if err != nil {
		return err
	}

	userId, err := h.Params.GetUserIdFromLocals(ctx)
	if err != nil {
		return err
	}

	cardId, err := h.Params.GetUUIDParam(ctx, "card_id")
	if err != nil {
		return err
	}

	card, err := h.CardService.Get(*cardId)
	if err != nil {
		return err
	}

	if card.UserId.String() != userId.String() {
		return http.NotFoundError("Card not found")
	}

	err = h.CardService.Update(*cardId, updateCardRequest)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// DeleteCard godoc
// @Summary Delete card by id
// @Description Delete card by id
// @Tags cards
// @Produce json
// @Success 204 {string} nil deletion is successful
// @Router / [delete]
func (h *Handler) DeleteCard(ctx *fiber.Ctx) error {
	cardId, err := h.Params.GetUUIDParam(ctx, "card_id")
	if err != nil {
		return err
	}

	userId, err := h.Params.GetUserIdFromLocals(ctx)
	if err != nil {
		return err
	}

	card, err := h.CardService.Get(*cardId)
	if err != nil {
		return err
	}

	if card.UserId.String() != userId.String() {
		return http.NotFoundError("Card not found")
	}

	err = h.CardService.Delete(*cardId)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
