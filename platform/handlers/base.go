package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imanhodjaev/confetti/platform/middleware"
	"github.com/imanhodjaev/confetti/platform/shared"
)

func App(handler *Handler) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          shared.ErrorHandler,
	})

	authMiddleware := middleware.AuthMiddleware(
		"Authorization",
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

	cards := app.Group("/c")
	cards.Use(authMiddleware)
	cards.Post("/", handler.CreateCard)
	cards.Get("/:card_id", handler.CreateCard)
	cards.Delete("/:card_id", handler.CreateCard)
	cards.Put("/:card_id", handler.CreateCard)
	cards.Get("/:card_id/decrypt", handler.CreateCard)
	cards.Post("/new", handler.GenerateCard)

	auth := app.Group("/auth")
	auth.Get("/jwks", handler.JWKS)
	auth.Post("/register", handler.Register)
	auth.Post("/token", handler.AuthTokenFlow)
	auth.Post("/token/refresh", handler.RefreshToken)
	auth.Delete("/token", handler.LogOut)

	return app
}
