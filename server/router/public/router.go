package public

import (
	"ShellHooks.net/server/router/midst"
	"github.com/gofiber/fiber/v2"
)

/*
/api/public
*/
func Router(api fiber.Router) {
	r := api.Group("/public", MiddleWare)

	r.Get("/ping", midst.Ping)

	r.Post("/githubHooks", GithubHooks)
}
