package handler

import (
	"net/http"

	"github.com/ceit-aut/S7IE03/internal/port/http/middleware"
	"github.com/ceit-aut/S7IE03/internal/repositories"
	"github.com/ceit-aut/S7IE03/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	JWT          *auth.Auth
	Repositories repositories.Repositories
	Threshold    int
	Limit        int
}

// Health will return the status of application.
func (h *Handler) Health(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusOK)
}

// CreateRoutes will generate endpoints of application
func (h *Handler) CreateRoutes(api fiber.Router) {
	// creating middleware
	mid := middleware.Middleware{
		Repositories: h.Repositories,
		Auth:         h.JWT,
	}

	// status route
	api.Get("/hlz", h.Health)

	// user routes
	api.Post("/register", h.Register)
	api.Post("/login", h.Login)

	api.Use(mid.Authentication)

	// endpoints routes
	api.Post("/endpoints", h.RegisterEndpoint)
	api.Get("/endpoints", h.GetAllEndpoints)

	v1 := api.Group("/endpoint/:id")

	v1.Use(mid.UserEndpoint)

	v1.Get("/", h.GetEndpointStatus)
	v1.Get("/warnings", h.GetEndpointWarnings)
	v1.Delete("/", h.RemoveUserEndpoint)
}
