package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Login godoc
// @Summary Login user
// @Description Login user
// @Tags auth
// @Produce json
// @Failure 403 {object} shared.HTTPError Invalid auth details
// @Success 200 {object} schema.TokenResponse
// @Router /auth/login [post]
func (h *Handler) Login(ctx *fiber.Ctx) error {
	loginPayload, err := h.Params.LoginPayload(ctx)
	if err != nil {
		return err
	}

	tokenResponse, err := h.AuthService.Login(ctx, loginPayload)
	if err != nil {
		return err
	}

	return ctx.JSON(tokenResponse)
}

// RefreshToken godoc
// @Summary Refresh access token
// @Description Refresh access token
// @Tags auth
// @Produce json
// @Failure 403 {object} shared.HTTPError Invalid auth details
// @Success 200 {object} schema.TokenResponse
// @Router /auth/login [post]
func (h *Handler) RefreshToken(ctx *fiber.Ctx) error {
	tokenResponse, err := h.AuthService.RefreshAuthToken(ctx)
	if err != nil {
		return err
	}

	return ctx.JSON(tokenResponse)
}

// LoginOut godoc
// @Summary Logout user
// @Description Logout user
// @Tags auth
// @Produce json
// @Success 204 {string} nil Log out is successful
// @Router /auth/logout [get]
func (h *Handler) LoginOut(ctx *fiber.Ctx) error {
	return h.AuthService.Logout(ctx)
}
