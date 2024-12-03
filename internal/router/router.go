package router

import (
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func NewFiberRouter(songHandler handler.SongHandler, port int) *fiber.App {
	app := fiber.New()
	app.Static("/docs", "./docs")
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL: fmt.Sprintf("http://localhost:%v/docs/swagger.json", port),
	}))
	app.Get("/songs", songHandler.GetSongs)
	app.Get("/songs/:song_id/text", songHandler.GetSongText)
	app.Delete("/songs/:song_id", songHandler.DeleteSong)
	app.Put("/songs/:song_id", songHandler.UpdateSong)
	app.Post("/songs", songHandler.InsertSong)
	return app
}
