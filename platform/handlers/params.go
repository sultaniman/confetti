package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/imanhodjaev/getout/platform/schema"
	"github.com/imanhodjaev/getout/platform/services"
	"github.com/imanhodjaev/getout/platform/shared"
)

type ParamHandler struct {
	UserService services.UserService
}

// User params

func (p *ParamHandler) GetUser(c *fiber.Ctx) (*schema.UserResponse, error) {
	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return nil, &shared.ServiceError{
			Response:   err,
			StatusCode: fiber.StatusInternalServerError,
			ErrorCode:  shared.ServerError,
		}
	}

	return p.UserService.Get(userID)
}

func (p *ParamHandler) GetUserFromLocals(c *fiber.Ctx) (*schema.UserResponse, error) {
	userID, err := uuid.Parse(c.Locals("user_id").(string))
	if err != nil {
		return nil, &shared.ServiceError{
			Response:   err,
			StatusCode: fiber.StatusInternalServerError,
			ErrorCode:  shared.ServerError,
		}
	}

	return p.UserService.Get(userID)
}

func (p *ParamHandler) CreateUserPayload(c *fiber.Ctx) (*schema.NewUserRequest, error) {
	newUser := new(schema.NewUserRequest)
	if err := c.BodyParser(newUser); err != nil {
		return nil, &shared.ServiceError{
			Response:   err,
			StatusCode: fiber.StatusInternalServerError,
			ErrorCode:  shared.ServerError,
		}
	}

	return newUser, nil
}

func (p *ParamHandler) UpdateUserPayload(c *fiber.Ctx) (*schema.UpdateUserRequest, error) {
	updateUserPayload := new(schema.UpdateUserRequest)
	if err := c.BodyParser(updateUserPayload); err != nil {
		return nil, &shared.ServiceError{
			Response:   err,
			StatusCode: fiber.StatusInternalServerError,
			ErrorCode:  shared.ServerError,
		}
	}

	return updateUserPayload, nil
}

func (p *ParamHandler) UpdateUserEmailPayload(c *fiber.Ctx) (*schema.UpdateUserEmailRequest, error) {
	updateUserEmailPayload := new(schema.UpdateUserEmailRequest)
	if err := c.BodyParser(updateUserEmailPayload); err != nil {
		return nil, &shared.ServiceError{
			Response:   err,
			StatusCode: fiber.StatusInternalServerError,
			ErrorCode:  shared.ServerError,
		}
	}

	return updateUserEmailPayload, nil
}

func (p *ParamHandler) UpdateUserPasswordPayload(c *fiber.Ctx) (*schema.UpdateUserPasswordRequest, error) {
	updateUserPasswordPayload := new(schema.UpdateUserPasswordRequest)
	if err := c.BodyParser(updateUserPasswordPayload); err != nil {
		return nil, &shared.ServiceError{
			Response:   err,
			StatusCode: fiber.StatusInternalServerError,
			ErrorCode:  shared.ServerError,
		}
	}

	return updateUserPasswordPayload, nil
}

func (p *ParamHandler) LoginPayload(ctx *fiber.Ctx) (*schema.LoginRequest, error) {
	loginRequestPayload := &schema.LoginRequest{}
	if err := ctx.BodyParser(loginRequestPayload); err != nil {
		return nil, &shared.ServiceError{
			Response:   err,
			StatusCode: fiber.StatusInternalServerError,
			ErrorCode:  shared.ServerError,
		}
	}

	return loginRequestPayload, nil
}

func (p *ParamHandler) RegisterPayload(ctx *fiber.Ctx) (*schema.RegisterRequest, error) {
	registerRequestPayload := &schema.RegisterRequest{}
	if err := ctx.BodyParser(registerRequestPayload); err != nil {
		return nil, &shared.ServiceError{
			Response:   err,
			StatusCode: fiber.StatusInternalServerError,
			ErrorCode:  shared.ServerError,
		}
	}

	return registerRequestPayload, nil
}
