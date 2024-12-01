package router

import (
	"github.com/YurcheuskiRadzivon/online_music_library/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func NewFiberRouter(songHandler handler.SongHandler) *fiber.App {
	app := fiber.New()
	app.Get("/songs", songHandler.GetSongs)
	app.Get("/songs/:song_id/text", songHandler.GetSong)
	app.Delete("/songs/:song_id", songHandler.DeleteSong)
	app.Put("/songs/:song_id", songHandler.UpdateSong)
	app.Post("/songs", songHandler.InsertSong)
	return app
}
