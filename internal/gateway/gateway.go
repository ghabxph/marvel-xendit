package gateway

import (
	"github.com/ghabxph/marvel-xendit/internal/marvel"
	"github.com/gofiber/fiber/v2"
)

type gateway struct {
	db interface{}
}
var instance *gateway

func GetInstance(db ...interface{}) *gateway {
	if instance == nil {
		instance = &gateway{db:db[0]}
	}
	return instance
}

func (g *gateway) Fiber() *fiber.App {
	app := fiber.New()

	bl := marvel.GetInstance(g.db)

	app.Get("/characters", func(c *fiber.Ctx) error {
		return c.SendString(bl.GetAllCharacters())
	})

	app.Get("/characters/:id", func(c *fiber.Ctx) error {
		return c.SendString(bl.GetCharacter(c.Params("id")))
	})

	return app
}

func (g *gateway) Serve() {
	g.Fiber().Listen(":8080")
}
