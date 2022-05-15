package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/sultaniman/confetti/platform/http"
	"github.com/sultaniman/confetti/platform/shared"
	"time"
)

const authScheme = "Bearer"

func AuthMiddleware(authHeader string, authRequired bool, jwks *jwk.Set) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		auth := ctx.Get(authHeader)
		if len(auth) <= len(authScheme) || auth[:len(authScheme)] != authScheme {
			if authRequired {
				return http.ForbiddenError("Please authenticate")
			}

			return ctx.Next()
		}

		tokenStr := auth[len(authScheme)+1:]
		payload, err := jwt.Parse([]byte(tokenStr), jwt.WithKeySet(*jwks))
		if err != nil {
			return &shared.ServiceError{
				Response:             "failed to verify token",
				StatusCode:           fiber.StatusUnauthorized,
				ErrorCode:            shared.Unauthorized,
				UseResponseAsMessage: shared.Bool(false),
			}
		}

		if payload.Expiration().Before(time.Now()) {
			return &shared.ServiceError{
				Response:             "token has expired",
				StatusCode:           fiber.StatusUnauthorized,
				ErrorCode:            shared.TokenExpired,
				UseResponseAsMessage: shared.Bool(false),
			}
		}

		subject := payload.Subject()
		if subject == "" {
			return http.ForbiddenError("Invalid token")
		}

		ctx.Locals("user_id", subject)
		return ctx.Next()
	}
}
