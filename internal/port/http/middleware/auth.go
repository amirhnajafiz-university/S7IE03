package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) Authentication(ctx *fiber.Ctx) error {
	if token, ok := ctx.GetReqHeaders()["token"]; ok {
		if user, pass, err := m.Auth.ParseJWT(token); err == nil {
			ctx.Locals("username", user)
			ctx.Locals("password", pass)

			return ctx.Next()
		} else {
			return ctx.SendStatus(http.StatusUnauthorized)
		}
	} else {
		return ctx.SendStatus(http.StatusBadRequest)
	}
}
