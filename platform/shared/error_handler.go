package shared

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case *ServiceError:
		return c.
			Status(e.StatusCode).
			JSON(&HTTPError{
				Code:    e.ErrorCode,
				Message: e.Error(),
				Details: e.Response,
			})
	default:
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(ServerErrorResponse(err, ""))
	}
}
