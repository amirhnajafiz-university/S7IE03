package middleware

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	errEmptyId         = errors.New("id cannot be empty")
	errInvalidEndpoint = errors.New("endpoint not found")
)

// UserEndpoint middleware checks that right user is
// accessing right endpoint stats.
func (m *Middleware) UserEndpoint(ctx *fiber.Ctx) error {
	// get username
	username := ctx.Locals("username")

	// get endpoint id
	id := ctx.Params("id", "nil")
	if id == "nil" {
		return errEmptyId
	}

	// get endpoint
	ep := m.Repositories.Endpoints.GetSingle(id)
	if ep == nil {
		return errInvalidEndpoint
	}

	// check the validation
	if ep.Username == username {
		ctx.Locals("id", id)

		return ctx.Next()
	}

	return ctx.SendStatus(http.StatusForbidden)
}
