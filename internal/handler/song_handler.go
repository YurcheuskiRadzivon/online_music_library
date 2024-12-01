package handler

import (
	"context"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/controller"

	"github.com/gofiber/fiber/v2"
)

type SongHandler interface {
	GetSongs(c *fiber.Ctx) error
	GetSong(c *fiber.Ctx) error
	InsertSong(c *fiber.Ctx) error
	UpdateSong(c *fiber.Ctx) error
	DeleteSong(c *fiber.Ctx) error
}
type songHandler struct {
	ctx        context.Context
	controller controller.SongController
}

func NewSongHandler(controller controller.SongController) SongHandler {
	return &songHandler{
		controller: controller,
		ctx:        context.Background(),
	}

}
func (sh *songHandler) GetSongs(c *fiber.Ctx) error {
	return nil
}
func (sh *songHandler) GetSong(c *fiber.Ctx) error {
	return nil
}
func (sh *songHandler) InsertSong(c *fiber.Ctx) error {
	return nil
}
func (sh *songHandler) UpdateSong(c *fiber.Ctx) error {
	return nil
}
func (sh *songHandler) DeleteSong(c *fiber.Ctx) error {
	return nil
}
