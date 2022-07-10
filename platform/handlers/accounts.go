package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sultaniman/confetti/platform/http"
)

// Register godoc
// @Summary Register using email and password
// @Description Register using email and password
// @Tags accounts
// @Produce json
// @Success 204 {string} nil registration is successful
// @Router /accounts/register [post]
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
// @Tags accounts
// @Produce json
// @Success 204 {string} nil reset link was sent
// @Router /accounts/reset-password [post]
func (h *Handler) ResetPasswordRequest(ctx *fiber.Ctx) error {
	resetPasswordPayload, err := h.Params.ResetPasswordRequestPayload(ctx)
	if err != nil {
		return err
	}

	_ = h.AuthService.ResetPasswordRequest(resetPasswordPayload)
	return ctx.SendStatus(fiber.StatusNoContent)
}

// ResetPassword godoc
// @Summary Reset password for a given reset code
// @Description Reset password for a given reset code
// @Tags accounts
// @Produce json
// @Success 204 {string} nil reset link was sent
// @Router /accounts/reset-password/{code} [post]
func (h *Handler) ResetPassword(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	passwordResetCode, err := h.UserService.GetResetPasswordCode(code)
	if err != nil {
		return http.BadRequestWithMessage("Invalid password reset code")
	}

	newPasswordRequest, err := h.Params.NewPasswordPayload(ctx)
	if err != nil {
		return err
	}

	// TODO: check if password is the same
	err = h.UserService.ResetPassword(passwordResetCode.UserId, newPasswordRequest.Password)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// Confirm godoc
// @Summary Confirmation for user accounts
// @Description Confirmation for user accounts
// @Tags accounts
// @Produce json
// @Success 204 {string} nil registration is successful
// @Router /accounts/confirm/{code} [get]
func (h *Handler) Confirm(ctx *fiber.Ctx) error {
	code := ctx.Params("code")

	userId, err := h.Params.GetUserIdFromLocals(ctx)
	if err != nil {
		return err
	}

	err = h.UserService.ConfirmUser(*userId, code)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// ResendConfirmation godoc
// @Summary Confirmation for user accounts
// @Description Confirmation for user accounts
// @Tags accounts
// @Produce json
// @Success 204 {string} nil confirmation code resent
// @Router /accounts/resend-confirmation [post]
func (h *Handler) ResendConfirmation(ctx *fiber.Ctx) error {
	userId, err := h.Params.GetUserIdFromLocals(ctx)
	if err != nil {
		return err
	}

	err = h.UserService.ResendConfirmation(*userId)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
