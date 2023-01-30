package cmd

import (
	"fmt"

	"github.com/ceit-aut/S7IE03/internal/config"
	"github.com/ceit-aut/S7IE03/internal/port/http/handler"
	"github.com/ceit-aut/S7IE03/internal/repositories"
	"github.com/ceit-aut/S7IE03/pkg/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

type HTTP struct {
	Cfg          config.Config
	Auth         *auth.Auth
	Repositories repositories.Repositories
}

func (h HTTP) Command() *cobra.Command {
	run := func(_ *cobra.Command, _ []string) { h.main() }
	return &cobra.Command{Use: "http", Short: "start http server", Run: run}
}

func (h HTTP) main() {
	// create fiber app
	app := fiber.New()

	// create api group
	api := app.Group("/api")

	// create handler
	hdl := handler.Handler{
		JWT:          h.Auth,
		Repositories: h.Repositories,
		Threshold:    h.Cfg.Threshold,
		Limit:        h.Cfg.UserEndpoints,
	}

	// register routes
	hdl.CreateRoutes(api)

	// start http server
	if err := app.Listen(fmt.Sprintf(":%d", h.Cfg.HttpPort)); err != nil {
		panic(err)
	}
}
