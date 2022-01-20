package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imanhodjaev/confetti/platform/shared"
)

func NotFoundError(message string) *shared.ServiceError {
	return &shared.ServiceError{
		Response:             message,
		StatusCode:           fiber.StatusNotFound,
		ErrorCode:            shared.NotFound,
		UseResponseAsMessage: shared.Bool(true),
	}
}

func UnauthorizedError(message string) *shared.ServiceError {
	return &shared.ServiceError{
		Response:             message,
		StatusCode:           fiber.StatusUnauthorized,
		ErrorCode:            shared.Unauthorized,
		UseResponseAsMessage: shared.Bool(true),
	}
}

func ForbiddenError(message string) *shared.ServiceError {
	return &shared.ServiceError{
		Response:             message,
		StatusCode:           fiber.StatusForbidden,
		ErrorCode:            shared.Forbidden,
		UseResponseAsMessage: shared.Bool(true),
	}
}

func InternalError(err error) *shared.ServiceError {
	return &shared.ServiceError{
		Response:   err,
		StatusCode: fiber.StatusInternalServerError,
		ErrorCode:  shared.ServerError,
	}
}

func Conflict(message string) *shared.ServiceError {
	return &shared.ServiceError{
		Response:             message,
		StatusCode:           fiber.StatusConflict,
		ErrorCode:            shared.Conflict,
		UseResponseAsMessage: shared.Bool(true),
	}
}

func BadRequestWithMessage(message string) *shared.ServiceError {
	return &shared.ServiceError{
		Response:             message,
		StatusCode:           fiber.StatusBadRequest,
		ErrorCode:            shared.BadRequest,
		UseResponseAsMessage: shared.Bool(true),
	}
}

func FetchError(err error, message string) *shared.ServiceError {
	if err != nil {
		return &shared.ServiceError{
			Response:             err,
			StatusCode:           fiber.StatusInternalServerError,
			ErrorCode:            shared.ServerError,
			UseResponseAsMessage: shared.Bool(true),
		}
	}

	return &shared.ServiceError{
		Response:   message,
		StatusCode: fiber.StatusInternalServerError,
		ErrorCode:  shared.ServerError,
	}
}

func UpdateError(err error, message string) *shared.ServiceError {
	if err != nil {
		return &shared.ServiceError{
			Response:             err,
			StatusCode:           fiber.StatusInternalServerError,
			ErrorCode:            shared.UpdateError,
			UseResponseAsMessage: shared.Bool(true),
		}
	}

	return &shared.ServiceError{
		Response:   message,
		StatusCode: fiber.StatusInternalServerError,
		ErrorCode:  shared.UpdateError,
	}
}

func InvalidPasswordError() *shared.ServiceError {
	return &shared.ServiceError{
		Response:             "Invalid password",
		StatusCode:           fiber.StatusForbidden,
		ErrorCode:            shared.Forbidden,
		UseResponseAsMessage: shared.Bool(true),
	}
}

func InactiveUserError() *shared.ServiceError {
	return &shared.ServiceError{
		Response:             "User is not active",
		StatusCode:           fiber.StatusForbidden,
		ErrorCode:            shared.InactiveUser,
		UseResponseAsMessage: shared.Bool(true),
	}
}

func InsecurePasswordError() *shared.ServiceError {
	return &shared.ServiceError{
		Response:             "Insecure password",
		StatusCode:           fiber.StatusBadRequest,
		ErrorCode:            shared.BadRequest,
		UseResponseAsMessage: shared.Bool(true),
	}
}

func EncryptionError(err error) *shared.ServiceError {
	return &shared.ServiceError{
		Response:   err,
		StatusCode: fiber.StatusInternalServerError,
		ErrorCode:  shared.EncryptionError,
	}
}

func DecryptionError(err error) *shared.ServiceError {
	return &shared.ServiceError{
		Response:   err,
		StatusCode: fiber.StatusInternalServerError,
		ErrorCode:  shared.DecryptionError,
	}
}

func DecodingError(err error) *shared.ServiceError {
	return &shared.ServiceError{
		Response:   err,
		StatusCode: fiber.StatusInternalServerError,
		ErrorCode:  shared.DecodingError,
	}
}
