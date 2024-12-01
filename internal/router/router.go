package router

import "github.com/gofiber/fiber/v2"

func NewFiberRouter() *fiber.App {
	app := fiber.New()
	app.Get("/songs")
	app.Get("/songs/:song_id/text")
	app.Delete("/songs/:song_id")
	app.Put("/songs/:song_id")
	app.Post("/songs")
	return app
}
