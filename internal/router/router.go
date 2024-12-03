package router

import (
	"github.com/YurcheuskiRadzivon/online_music_library/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func NewFiberRouter(songHandler handler.SongHandler) *fiber.App {
	app := fiber.New()
	app.Static("/docs", "./docs")
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL: "http://localhost:8080/docs/swagger.json",
	}))
	app.Get("/songs", songHandler.GetSongs)
	app.Get("/songs/:song_id/text", songHandler.GetSongText)
	app.Delete("/songs/:song_id", songHandler.DeleteSong)
	app.Put("/songs/:song_id", songHandler.UpdateSong)
	app.Post("/songs", songHandler.InsertSong)
	return app
}
