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
// @Router /auth/login [post]
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
// @Router /auth/login [post]
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

// Register godoc
// @Summary Register using email and password
// @Description Register using email and password
// @Tags auth
// @Produce json
// @Success 204 {string} nil registration is successful
// @Router /auth/register [post]
func (h *Handler) Register(ctx *fiber.Ctx) error {
	registerPayload, err := h.Params.RegisterPayload(ctx)
	if err != nil {
		return err
	}

	err = h.AuthService.Register(registerPayload)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// ResetPasswordRequest godoc
// @Summary Reset password using email
// @Description Reset password using email
// @Tags auth
// @Produce json
// @Success 204 {string} nil reset link was sent
// @Router /auth/reset-password [post]
func (h *Handler) ResetPasswordRequest(ctx *fiber.Ctx) error {
	resetPasswordPayload, err := h.Params.ResetPasswordPayload(ctx)
	if err != nil {
		return err
	}

	_ = h.AuthService.ResetPasswordRequest(resetPasswordPayload)
	return ctx.SendStatus(fiber.StatusNoContent)
}

// Confirm godoc
// @Summary Confirmation for user accounts
// @Description Confirmation for user accounts
// @Tags auth
// @Produce json
// @Success 204 {string} nil registration is successful
// @Router /confirm/{code} [get]
func (h *Handler) Confirm(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusNoContent)
}

// ResendConfirmation godoc
// @Summary Confirmation for user accounts
// @Description Confirmation for user accounts
// @Tags auth
// @Produce json
// @Success 204 {string} nil confirmation code resent
// @Router /confirm/resend [post]
func (h *Handler) ResendConfirmation(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusNoContent)
}
