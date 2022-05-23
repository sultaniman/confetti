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

// ListCards godoc
// @Summary List cards
// @Description List cards
// @Tags cards
// @Produce json
// @Success 200 {object} []schema.CardResponse
// @Router / [post]
func (h *Handler) ListCards(ctx *fiber.Ctx) error {
	userId, err := h.Params.GetUserIdFromLocals(ctx)
	if err != nil {
		return err
	}

	cards, err := h.CardService.List(*userId)
	if err != nil {
		return err
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(cards)
}

// GetCard godoc
// @Summary Get card by id
// @Description Get card by id
// @Tags cards
// @Produce json
// @Success 200 {object} schema.CardResponse
// @Router /{id} [get]
func (h *Handler) GetCard(ctx *fiber.Ctx) error {
	claim, err := h.Params.EnsureCardClaim(ctx)
	if err != nil {
		return err
	}

	card, err := h.CardService.Get(claim.CardId)
	if err != nil {
		return err
	}

	return ctx.JSON(card)
}

// UpdateCard godoc
// @Summary Update card
// @Description Update card
// @Tags cards
// @Produce json
// @Success 204 {string} nil update succeeded
// @Router /{id} [put]
func (h *Handler) UpdateCard(ctx *fiber.Ctx) error {
	claim, err := h.Params.EnsureCardClaim(ctx)
	if err != nil {
		return err
	}

	updateCardRequest, err := h.Params.UpdateCardPayload(ctx)
	if err != nil {
		return err
	}

	err = h.CardService.Update(claim.CardId, updateCardRequest)
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
// @Router /{id} [delete]
func (h *Handler) DeleteCard(ctx *fiber.Ctx) error {
	claim, err := h.Params.EnsureCardClaim(ctx)
	if err != nil {
		return err
	}

	err = h.CardService.Delete(claim.CardId)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// DecryptCard godoc
// @Summary Decrypt card by id
// @Description Decrypt card by id
// @Tags cards
// @Produce json
// @Success 200 {object} schema.PlainCardResponse
// @Router /{id}/decrypt [get]
func (h *Handler) DecryptCard(ctx *fiber.Ctx) error {
	claim, err := h.Params.EnsureCardClaim(ctx)
	if err != nil {
		return err
	}

	plainCard, err := h.CardService.Decrypt(claim.CardId)
	if err != nil {
		return err
	}

	return ctx.JSON(plainCard)
}
