package handler

import (
	"github.com/ceit-aut/policeman/internal/port/http/middleware"
	"github.com/ceit-aut/policeman/internal/repositories"
	"github.com/ceit-aut/policeman/pkg/auth"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	JWT          auth.Auth
	Repositories repositories.Repositories
	Threshold    int
}

// CreateRoutes will generate endpoints of application
func (h *Handler) CreateRoutes(app *fiber.Group, cfg auth.Config) {
	// creating middleware
	mid := middleware.Middleware{
		Auth: auth.New(cfg),
	}

	// user endpoints
	app.Post("/api/register", h.Register)
	app.Post("/api/login", h.Login)

	app.Use(mid.Authentication)

	// endpoints endpoints
	app.Post("/api/endpoints", h.RegisterEndpoint)
	app.Get("/api/endpoints", h.GetAllEndpoints)

	app.Use(mid.UserEndpoint)

	app.Get("/api/endpoint/:id", h.GetEndpointStatus)
	app.Get("/api/endpoint/:id/warnings", h.GetEndpointWarnings)
}
