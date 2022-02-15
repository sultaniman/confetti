package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sultaniman/confetti/platform/middleware"
	"github.com/sultaniman/confetti/platform/shared"
)

func App(handler *Handler) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          shared.ErrorHandler,
	})
	app.Use(cors.New())

	authMiddleware := middleware.AuthMiddleware(
		"Authorization",
		true,
		handler.JWXService.JWKS(),
	)

	system := app.Group("/system")
	system.Get("/health", handler.Health)

	admin := app.Group("/admin")
	admin.Use(authMiddleware)

	users := admin.Group("/users")
	users.Post("/", handler.CreateUser)
	users.Get("/:user_id", handler.GetUser)
	users.Put("/:user_id", handler.UpdateUser)
	users.Put("/:user_id/email", handler.UpdateEmail)
	users.Put("/:user_id/password", handler.UpdatePassword)
	users.Delete("/:user_id", handler.DeleteUser)

	cards := app.Group("/cards")
	cards.Use(authMiddleware)
	cards.Get("/", handler.ListCards)
	cards.Post("/", handler.CreateCard)
	cards.Get("/:card_id", handler.GetCard)
	cards.Delete("/:card_id", handler.DeleteCard)
	cards.Put("/:card_id", handler.UpdateCard)
	cards.Get("/:card_id/decrypt", handler.DecryptCard)
	cards.Post("/new", handler.GenerateCard)

	accounts := app.Group("/accounts")
	accounts.Post("/register", handler.Register)
	accounts.Get("/confirm/:code", handler.Confirm)
	accounts.Post("/resend-confirmation", authMiddleware, handler.ResendConfirmation)
	accounts.Post("/reset-password", handler.ResetPasswordRequest)
	accounts.Post("/reset-password/:code", handler.ResetPassword)

	auth := app.Group("/auth")
	auth.Get("/jwks", handler.JWKS)
	auth.Post("/token", handler.AuthTokenFlow)
	auth.Post("/token/refresh", handler.RefreshToken)
	auth.Delete("/token", handler.LogOut)

	return app
}
