package handler

import (
	"net/http"

	"github.com/ceit-aut/policeman/internal/port/http/middleware"
	"github.com/ceit-aut/policeman/internal/repositories"
	"github.com/ceit-aut/policeman/pkg/auth"
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
func (h *Handler) CreateRoutes(app fiber.Router) {
	// creating middleware
	mid := middleware.Middleware{
		Repositories: h.Repositories,
		Auth:         h.JWT,
	}

	// status route
	app.Get("/hlz", h.Health)

	// user routes
	app.Post("/register", h.Register)
	app.Post("/login", h.Login)

	app.Use(mid.Authentication)

	// endpoints routes
	app.Post("/endpoints", h.RegisterEndpoint)
	app.Get("/endpoints", h.GetAllEndpoints)

	v1 := app.Group("/endpoint/:id")

	v1.Use(mid.UserEndpoint)

	v1.Get("/", h.GetEndpointStatus)
	v1.Get("/warnings", h.GetEndpointWarnings)
	v1.Delete("/", h.RemoveUserEndpoint)
}
