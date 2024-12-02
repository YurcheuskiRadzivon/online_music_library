package handler

import (
	"context"
	"strconv"

	"github.com/YurcheuskiRadzivon/online_music_library/internal/controller"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/model"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
	"github.com/gofiber/fiber/v2"
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
	lgr        *logger.Logger
}

func NewSongHandler(controller controller.SongController, lgr *logger.Logger) SongHandler {
	return &songHandler{
		controller: controller,
		ctx:        context.Background(),
		lgr:        lgr,
	}
}

// @Summary      Get all songs
// @Description  Retrieve a list of songs with pagination and sorting
// @Tags         songs
// @Param        sort      query    string  false  "Field to sort by" Enums(sound_id,text_length,song,release_date)
// @Param        page      query    int     true   "Page number"
// @Param        page_size query    int     true   "Number of items per page"
// @Success      200  {array}  model.Song
// @Failure      500  {object} map[string]interface{}
// @Router       /songs [get]
func (sh *songHandler) GetSongs(c *fiber.Ctx) error {
	sortParam := c.Query("sort", "sound_id")
	page := getPage(c, 1, sh.lgr)
	pageSize := getPageSize(c, 10, sh.lgr)

	paginatedSongs, err := sh.controller.GetSongs(c.Context(), sortParam, page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	sh.lgr.InfoLogger.Printf("Returned %d songs\n", len(paginatedSongs))
	return c.JSON(paginatedSongs)
}

func (sh *songHandler) GetSong(c *fiber.Ctx) error {
	sh.lgr.DebugLogger.Println("GetSong is not implemented")
	return nil
}

// @Summary      Get song text
// @Description  Retrieve the text of a song with pagination
// @Tags         songs
// @Param        song_id   path     int     true   "ID of the song"
// @Param        page      query    int     true   "Page number"
// @Param        page_size query    int     true   "Number of items per page"
// @Success      200  {array}  string
// @Failure      400  {object} map[string]interface{}
// @Failure      500  {object} map[string]interface{}
// @Router       /songs/{song_id}/text [get]
func (sh *songHandler) GetSongText(c *fiber.Ctx) error {
	songIDStr := c.Params("song_id")
	songId, err := strconv.Atoi(songIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid song_id"})
	}

	page := getPage(c, 1, sh.lgr)
	pageSize := getPageSize(c, 1, sh.lgr)

	verses, err := sh.controller.GetSongText(c.Context(), songId, pageSize, page)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(verses)
}

// @Summary      Insert a new song
// @Description  Insert a new song from a SongRequest
// @Tags         songs
// @Param        songRequest body    model.SongRequest true "Song request object"
// @Success      201  {object} map[string]interface{}
// @Failure      400  {object} map[string]interface{}
// @Failure      500  {object} map[string]interface{}
// @Router       /songs [post]
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

	sh.lgr.DebugLogger.Printf("InsertSong called with group: %s, song: %s\n", songRequest.Group, songRequest.Song)

	if err := sh.controller.InsertSong(c.Context(), songRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	sh.lgr.InfoLogger.Printf("Song inserted successfully\n")
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Song inserted successfully",
	})
}

// @Summary      Update an existing song
// @Description  Update a song by ID
// @Tags         songs
// @Param        song_id path     int     true   "ID of the song"
// @Param        song    body     model.Song true "Updated song object"
// @Success      200  {object} map[string]interface{}
// @Failure      400  {object} map[string]interface{}
// @Failure      500  {object} map[string]interface{}
// @Router       /songs/{song_id} [put]
func (sh *songHandler) UpdateSong(c *fiber.Ctx) error {
	songIDStr := c.Params("song_id")
	songID, err := strconv.Atoi(songIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid song ID",
		})
	}

	var song model.Song
	err = c.BodyParser(&song)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON body",
		})
	}

	if err := sh.controller.UpdateSong(c.Context(), songID, song); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	sh.lgr.InfoLogger.Printf("Song updated successfully\n")
	return c.JSON(fiber.Map{
		"message": "Song updated successfully",
	})
}

// @Summary      Delete a song
// @Description  Delete a song by ID
// @Tags         songs
// @Param        song_id path     int     true   "ID of the song"
// @Success      200  {object} map[string]interface{}
// @Failure      400  {object} map[string]interface{}
// @Failure      500  {object} map[string]interface{}
// @Router       /songs/{song_id} [delete]
func (sh *songHandler) DeleteSong(c *fiber.Ctx) error {
	songIDStr := c.Params("song_id")
	songID, err := strconv.Atoi(songIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid song ID",
		})
	}

	if err := sh.controller.DeleteSong(c.Context(), songID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	sh.lgr.InfoLogger.Printf("Song deleted successfully\n")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Song deleted successfully",
	})
}

func getPage(c *fiber.Ctx, defaultValue int, lgr *logger.Logger) int {
	pageStr := c.Query("page", strconv.Itoa(defaultValue))
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		lgr.DebugLogger.Printf("Invalid page parameter, using default: %d\n", defaultValue)
		return defaultValue
	}
	return page
}

func getPageSize(c *fiber.Ctx, defaultValue int, lgr *logger.Logger) int {
	pageSizeStr := c.Query("page_size", strconv.Itoa(defaultValue))
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		lgr.DebugLogger.Printf("Invalid page_size parameter, using default: %d\n", defaultValue)
		return defaultValue
	}
	return pageSize
}
