package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imanhodjaev/getout/platform/shared"
)

func App(handler *Handler) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          shared.ErrorHandler,
	})

	system := app.Group("/system")
	system.Get("/health", handler.Health)

	admin := app.Group("/admin")

	users := admin.Group("/users")
	users.Post("/", handler.CreateUser)
	users.Get("/:user_id", handler.GetUser)
	users.Put("/:user_id", handler.UpdateUser)
	users.Put("/:user_id/email", handler.UpdateEmail)
	users.Put("/:user_id/password", handler.UpdatePassword)
	users.Delete("/:user_id", handler.DeleteUser)

	auth := app.Group("/auth")
	auth.Get("/jwks", handler.JWKS)
	auth.Post("/register", handler.Register)
	auth.Post("/token", handler.AuthTokenFlow)
	auth.Post("/token/refresh", handler.RefreshToken)
	auth.Delete("/token", handler.LogOut)

	return app
}
