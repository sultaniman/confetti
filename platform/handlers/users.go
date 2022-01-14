package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// GetUser godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Produce json
// @Failure 404 {object} shared.HTTPError User not found
// @Success 200 {object} schema.UserResponse
// @Router /user/{user_id} [get]
func (h *Handler) GetUser(ctx *fiber.Ctx) error {
	user, err := h.Params.GetUser(ctx)
	if err != nil {
		return err
	}

	return ctx.JSON(user)
}

// CreateUser godoc
// @Summary Create user
// @Description Create user
// @Tags users
// @Produce json
// @Success 201 {object} schema.UserResponse
// @Router /user [post]
func (h *Handler) CreateUser(ctx *fiber.Ctx) error {
	newUser, err := h.Params.CreateUserPayload(ctx)
	if err != nil {
		return err
	}

	user, err := h.UserService.Create(newUser)
	if err != nil {
		return err
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(user)
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user
// @Tags users
// @Produce json
// @Success 202 {object} schema.UserResponse
// @Router /user [put]
func (h *Handler) UpdateUser(ctx *fiber.Ctx) error {
	user, err := h.Params.GetUser(ctx)
	if err != nil {
		return err
	}

	updateUserPayload, err := h.Params.UpdateUserPayload(ctx)
	if err != nil {
		return err
	}

	user, err = h.UserService.Update(user.ID, updateUserPayload)
	if err != nil {
		return err
	}

	return ctx.
		Status(fiber.StatusAccepted).
		JSON(user)
}

// UpdateEmail godoc
// @Summary Update email for user
// @Description Update email for user
// @Tags users
// @Produce json
// @Success 202 {object} schema.UserResponse
// @Router /user/{user_id}/email [put]
func (h *Handler) UpdateEmail(ctx *fiber.Ctx) error {
	user, err := h.Params.GetUser(ctx)
	if err != nil {
		return err
	}

	updateUserEmailPayload, err := h.Params.UpdateUserEmailPayload(ctx)
	if err != nil {
		return err
	}

	user, err = h.UserService.UpdateEmail(user.ID, updateUserEmailPayload)
	if err != nil {
		return err
	}

	return ctx.
		Status(fiber.StatusAccepted).
		JSON(user)
}

// UpdatePassword godoc
// @Summary Update password for user
// @Description Update password for user
// @Tags users
// @Produce json
// @Success 202 {object} schema.UserResponse
// @Router /user/{user_id}/password [put]
func (h *Handler) UpdatePassword(ctx *fiber.Ctx) error {
	user, err := h.Params.GetUser(ctx)
	if err != nil {
		return err
	}

	updateUserPasswordPayload, err := h.Params.UpdateUserPasswordPayload(ctx)
	if err != nil {
		return err
	}

	user, err = h.UserService.UpdatePassword(user.ID, updateUserPasswordPayload)
	if err != nil {
		return err
	}

	return ctx.
		Status(fiber.StatusAccepted).
		JSON(user)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user
// @Tags users
// @Produce json
// @Success 200 {object} schema.UserResponse
// @Router /user [delete]
func (h *Handler) DeleteUser(ctx *fiber.Ctx) error {
	user, err := h.Params.GetUser(ctx)
	if err != nil {
		return err
	}

	return ctx.JSON(user)
}
