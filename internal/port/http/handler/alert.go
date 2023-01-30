package handler

import (
	"github.com/gofiber/fiber/v2"
)

// GetAlerts for a single endpoint.
func (h *Handler) GetAlerts(ctx *fiber.Ctx) error {
	// get alerts
	alerts := h.Repositories.Alerts.GetAll(ctx.Locals("id").(string))

	return ctx.JSON(alerts)
}
