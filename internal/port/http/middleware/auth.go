package middleware

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Authentication is used to check users tokens in header.
func (m *Middleware) Authentication(ctx *fiber.Ctx) error {
	if token := ctx.Get("token", ""); token != "" {
		if user, pass, err := m.Auth.ParseJWT(token); err == nil {
			// set variables into request context
			ctx.Locals("username", user)
			ctx.Locals("password", pass)

			return ctx.Next()
		} else {
			log.Println(err)

			// 401
			return ctx.SendStatus(http.StatusUnauthorized)
		}
	} else {
		// 400
		return ctx.SendStatus(http.StatusBadRequest)
	}
}
