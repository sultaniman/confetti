package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sultaniman/confetti/platform/http"
	"github.com/sultaniman/confetti/platform/schema"
	"github.com/sultaniman/confetti/platform/services"
	"github.com/sultaniman/confetti/platform/shared"
)

type ParamHandler struct {
	UserService services.UserService
	CardService services.CardService
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

func (p *ParamHandler) GetUserIdFromLocals(c *fiber.Ctx) (*uuid.UUID, error) {
	userID, err := uuid.Parse(c.Locals("user_id").(string))
	if err != nil {
		return nil, &shared.ServiceError{
			Response:   err,
			StatusCode: fiber.StatusInternalServerError,
			ErrorCode:  shared.ServerError,
		}
	}

	return &userID, nil
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

func (p *ParamHandler) ResetPasswordRequestPayload(ctx *fiber.Ctx) (*schema.ResetPasswordRequest, error) {
	resetPasswordRequest := &schema.ResetPasswordRequest{}
	if err := ctx.BodyParser(resetPasswordRequest); err != nil {
		return nil, &shared.ServiceError{
			Response:   err,
			StatusCode: fiber.StatusInternalServerError,
			ErrorCode:  shared.ServerError,
		}
	}

	return resetPasswordRequest, nil
}

func (p *ParamHandler) NewPasswordPayload(ctx *fiber.Ctx) (*schema.NewPasswordRequest, error) {
	newPasswordRequest := &schema.NewPasswordRequest{}
	if err := ctx.BodyParser(newPasswordRequest); err != nil {
		return nil, &shared.ServiceError{
			Response:   err,
			StatusCode: fiber.StatusInternalServerError,
			ErrorCode:  shared.ServerError,
		}
	}

	return newPasswordRequest, nil
}

// Card params

func (p *ParamHandler) CardOptionsPayload(c *fiber.Ctx) (*schema.CardOptions, error) {
	cardOptions := new(schema.CardOptions)
	if err := c.BodyParser(cardOptions); err != nil {
		return nil, &shared.ServiceError{
			Response:   err,
			StatusCode: fiber.StatusBadRequest,
			ErrorCode:  shared.BadRequest,
		}
	}

	return cardOptions, nil
}

func (p *ParamHandler) CreateCardPayload(c *fiber.Ctx) (*schema.NewCardRequest, error) {
	newCard := new(schema.NewCardRequest)
	if err := c.BodyParser(newCard); err != nil {
		return nil, &shared.ServiceError{
			Response:   err,
			StatusCode: fiber.StatusInternalServerError,
			ErrorCode:  shared.ServerError,
		}
	}

	return newCard, nil
}

func (p *ParamHandler) UpdateCardPayload(c *fiber.Ctx) (*schema.UpdateCardRequest, error) {
	updatePayload := new(schema.UpdateCardRequest)
	if err := c.BodyParser(updatePayload); err != nil {
		return nil, &shared.ServiceError{
			Response:   err,
			StatusCode: fiber.StatusInternalServerError,
			ErrorCode:  shared.ServerError,
		}
	}

	return updatePayload, nil
}

func (p *ParamHandler) EnsureCardClaim(c *fiber.Ctx) (*schema.CardClaim, error) {
	cardId, err := p.GetUUIDParam(c, "card_id")
	if err != nil {
		return nil, err
	}

	userId, err := p.GetUserIdFromLocals(c)
	if err != nil {
		return nil, err
	}

	if !p.CardService.ClaimExists(*cardId, *userId) {
		return nil, http.NotFoundError("Card not found")
	}

	return &schema.CardClaim{
		CardId: *cardId,
		UserId: *userId,
	}, nil
}

// Generic handlers

func (p *ParamHandler) GetUUIDParam(c *fiber.Ctx, paramName string) (*uuid.UUID, error) {
	idParam, err := uuid.Parse(c.Params(paramName))
	if err != nil {
		return nil, &shared.ServiceError{
			Response:   err,
			StatusCode: fiber.StatusInternalServerError,
			ErrorCode:  shared.ServerError,
		}
	}

	return &idParam, nil
}
