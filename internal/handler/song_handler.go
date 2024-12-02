package handler

import (
	"context"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/controller"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type SongHandler interface {
	GetSongs(c *fiber.Ctx) error
	GetSong(c *fiber.Ctx) error
	GetSongText(c *fiber.Ctx) error
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

	sortParam := c.Query("sort", "sound_id")

	page := getPage(c, 1)
	pageSize := getPageSize(c, 10)

	paginatedSongs, err := sh.controller.GetSongs(c.Context(), sortParam, page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(paginatedSongs)
}
func (sh *songHandler) GetSong(c *fiber.Ctx) error {
	return nil
}
func (sh *songHandler) GetSongText(c *fiber.Ctx) error {
	songIDStr := c.Params("song_id")
	songId, err := strconv.Atoi(songIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid song_id"})
	}

	page := getPage(c, 1)
	pageSize := getPageSize(c, 1)

	verses, err := sh.controller.GetSongText(c.Context(), songId, pageSize, page)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(verses)
}
func (sh *songHandler) InsertSong(c *fiber.Ctx) error {
	var songRequest model.SongRequest
	if err := c.BodyParser(&songRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if songRequest.Group == "" || songRequest.Song == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Group and Song fields are required",
		})
	}

	if err := sh.controller.InsertSong(c.Context(), songRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Song inserted successfully",
	})
}
func (sh *songHandler) UpdateSong(c *fiber.Ctx) error {
	songIDStr := c.Params("song_id")
	songID, err := strconv.Atoi(songIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid song ID",
		})
	}

	// Parse JSON request body into Song struct
	var song model.Song
	err = c.BodyParser(&song)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON body",
		})
	}

	// Call the controller's UpdateUser method
	err = sh.controller.UpdateUser(c.Context(), songID, song)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Song updated successfully",
	})
}
func (sh *songHandler) DeleteSong(c *fiber.Ctx) error {
	songIDStr := c.Params("song_id")
	songID, err := strconv.Atoi(songIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid song ID",
		})
	}
	err = sh.controller.DeleteUser(c.Context(), songID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Song deleted successfully",
	})
}
func getPage(c *fiber.Ctx, defaultValue int) int {
	pageStr := c.Query("page", strconv.Itoa(defaultValue))
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return defaultValue
	}
	return page
}

func getPageSize(c *fiber.Ctx, defaultValue int) int {
	pageSizeStr := c.Query("page_size", strconv.Itoa(defaultValue))
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		return defaultValue
	}
	return pageSize
}
