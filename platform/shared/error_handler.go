package shared

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case *ServiceError:
		details := e.Response
		errMsg := e.Error()
		if e.Response == errMsg {
			details = nil
		}

		return c.
			Status(e.StatusCode).
			JSON(&HTTPError{
				Code:    e.ErrorCode,
				Message: errMsg,
				Details: details,
			})
	default:
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(ServerErrorResponse(err, ServerError))
	}
}
