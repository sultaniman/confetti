package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imanhodjaev/confetti/platform/schema"
)

// Health get system health status
// @Summary Get system health status
// @Description Get system health status
// @Tags system
// @Produce json
// @Success 200 {object} schema.HealthResponse
// @Router /system/health [get]
func (h *Handler) Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).
		JSON(&schema.HealthResponse{OK: true})
}
