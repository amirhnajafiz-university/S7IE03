package handler

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterEndpoint will add an endpoint to application.
// takes a endpoint info request.
func (h *Handler) RegisterEndpoint(ctx *fiber.Ctx) error {
	return nil
}

// GetAllEndpoints for a user.
func (h *Handler) GetAllEndpoints(ctx *fiber.Ctx) error {
	return nil
}

// GetEndpointStatus will return one endpoint status.
func (h *Handler) GetEndpointStatus(ctx *fiber.Ctx) error {
	return nil
}

// GetEndpointWarnings will return all the warnings for an endpoint.
func (h *Handler) GetEndpointWarnings(ctx *fiber.Ctx) error {
	return nil
}
