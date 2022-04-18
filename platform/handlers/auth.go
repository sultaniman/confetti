package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// AuthTokenFlow godoc
// @Summary Authenticate user with email, password and issue access tokens
// @Description Authenticate user with email, password and issue access tokens
// @Tags auth
// @Produce json
// @Failure 403 {object} shared.HTTPError Invalid auth details
// @Success 200 {object} schema.TokenResponse
// @Router /token [post]
func (h *Handler) AuthTokenFlow(ctx *fiber.Ctx) error {
	loginPayload, err := h.Params.LoginPayload(ctx)
	if err != nil {
		return err
	}

	tokenResponse, err := h.AuthService.AccessTokenAuthFlow(ctx, loginPayload)
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
// @Router /token/refresh [post]
func (h *Handler) RefreshToken(ctx *fiber.Ctx) error {
	tokenResponse, err := h.AuthService.RefreshAuthToken(ctx)
	if err != nil {
		return err
	}

	return ctx.JSON(tokenResponse)
}

// LogOut godoc
// @Summary Logout user
// @Description Logout user
// @Tags auth
// @Produce json
// @Success 204 {string} nil Log out is successful
// @Router /auth/logout [delete]
func (h *Handler) LogOut(ctx *fiber.Ctx) error {
	return h.AuthService.Logout(ctx)
}

// JWKS godoc
// @Summary Returns jwks details
// @Description Returns jwks details
// @Tags auth
// @Produce json
// @Success 200 {object} nil
// @Router /auth/jwks [get]
func (h *Handler) JWKS(ctx *fiber.Ctx) error {
	return ctx.JSON(h.JWXService.JWKS())
}
