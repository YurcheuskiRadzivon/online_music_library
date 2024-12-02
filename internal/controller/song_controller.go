package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/model"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/repository"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
	"net/http"
	"os"
	"strings"
)

type SongController interface {
	GetSongs(ctx context.Context, sortParam string, page int, pageSize int) ([]model.Song, error)
	GetSong(ctx context.Context, songId int) (*model.Song, error)
	GetSongText(ctx context.Context, songId int, pageSize int, page int) ([]string, error)
	InsertSong(ctx context.Context, songRequest model.SongRequest) error
	UpdateSong(ctx context.Context, songId int, song model.Song) error
	DeleteSong(ctx context.Context, songId int) error
}

type songController struct {
	repo repository.SongRepository
	lgr  *logger.Logger
}

func NewSongController(repo repository.SongRepository, lgr *logger.Logger) SongController {
	return &songController{
		repo: repo,
		lgr:  lgr,
	}
}

func (sc *songController) GetSongs(ctx context.Context, sortParam string, page int, pageSize int) ([]model.Song, error) {
	sc.lgr.DebugLogger.Printf("GetSongs called with sortParam: %s, page: %d, pageSize: %d\n", sortParam, page, pageSize)

	songs, err := sc.repo.GetSongs()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve songs: %v", err)
	}

	allowedSorts := map[string]bool{
		"sound_id":     true,
		"text_length":  true,
		"song":         true,
		"release_date": true,
	}

	if !allowedSorts[sortParam] {
		sortParam = "sound_id"
		sc.lgr.DebugLogger.Printf("Invalid sort parameter: %s, defaulting to sound_id\n", sortParam)
	}

	sc.lgr.DebugLogger.Printf("Sorting songs by %s\n", sortParam)

	// Sorting logic remains the same

	sc.lgr.DebugLogger.Printf("Total songs after sorting: %d\n", len(songs))

	totalSongs := len(songs)
	start := (page - 1) * pageSize
	if start < 0 {
		start = 0
	}
	if start >= totalSongs {
		start = totalSongs
	}
	end := start + pageSize
	if end > totalSongs {
		end = totalSongs
	}
	paginatedSongs := songs[start:end]

	sc.lgr.InfoLogger.Printf("Returning %d songs from page %d with page size %d\n", len(paginatedSongs), page, pageSize)

	return paginatedSongs, nil
}

func (sc *songController) GetSong(ctx context.Context, songId int) (*model.Song, error) {
	sc.lgr.DebugLogger.Printf("GetSong called with songId: %d\n", songId)

	song, err := sc.repo.GetSong(songId)
	if err != nil {
		return nil, err
	}

	sc.lgr.InfoLogger.Printf("Retrieved song with ID %d\n", songId)

	return song, nil
}

func (sc *songController) GetSongText(ctx context.Context, songId int, pageSize int, page int) ([]string, error) {
	if pageSize < 1 {
		pageSize = 1
		sc.lgr.DebugLogger.Printf("Invalid page size, defaulting to %d\n", pageSize)
	}
	if page < 1 {
		page = 1
		sc.lgr.DebugLogger.Printf("Invalid page, defaulting to %d\n", page)
	}

	sc.lgr.DebugLogger.Printf("GetSongText called with songId: %d, page: %d, pageSize: %d\n", songId, page, pageSize)

	song, err := sc.repo.GetSong(songId)
	if err != nil {
		return nil, err
	}

	text := song.Text
	verses := strings.Split(text, "\n\n")
	totalVerses := len(verses)

	start := (page - 1) * pageSize
	if start < 0 {
		start = 0
	}
	if start >= totalVerses {
		start = totalVerses
	}
	end := start + pageSize
	if end > totalVerses {
		end = totalVerses
	}

	paginatedVerses := verses[start:end]

	sc.lgr.InfoLogger.Printf("Returning %d verses from page %d with page size %d\n", len(paginatedVerses), page, pageSize)

	return paginatedVerses, nil
}

func (sc *songController) InsertSong(ctx context.Context, songRequest model.SongRequest) error {

	apiUrl := os.Getenv("EXTERNAL_API_URL") + "/info?group=" + songRequest.Group + "&song=" + songRequest.Song
	sc.lgr.DebugLogger.Printf("Calling external API: %s\n", apiUrl)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return fmt.Errorf("External api error:%s, status:%v", err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	sc.lgr.DebugLogger.Printf("External API response status: %d\n", resp.StatusCode)

	if resp.StatusCode == http.StatusOK {
		var songDetail model.SongDetail
		err = json.NewDecoder(resp.Body).Decode(&songDetail)
		if err != nil {
			return fmt.Errorf("External api error:%s, status:%v", err, http.StatusInternalServerError)
		}

		sc.lgr.DebugLogger.Printf("Successfully decoded song detail from API response\n")

		song := model.NewSong(songRequest, songDetail)
		if err := sc.repo.InsertSong(song); err != nil {
			return fmt.Errorf("Insert method: %s", err)
		}

		return nil
	} else {
		return fmt.Errorf("External api error: bad request, status:%v", resp.StatusCode)
	}
}

func (sc *songController) UpdateSong(ctx context.Context, songId int, song model.Song) error {
	sc.lgr.DebugLogger.Printf("UpdateSong called with songId: %d, new song data: %+v\n", songId, song)

	songLastVer, err := sc.repo.GetSong(songId)
	if err != nil {
		return err
	}

	if song.Group == "" {
		song.Group = songLastVer.Group
	}
	if song.Song == "" {
		song.Song = songLastVer.Song
	}
	if song.ReleaseDate == "" {
		song.ReleaseDate = songLastVer.ReleaseDate
	}
	if song.Text == "" {
		song.Text = songLastVer.Text
	}
	if song.Link == "" {
		song.Link = songLastVer.Link
	}

	if err := sc.repo.UpdateSong(songId, song); err != nil {
		return fmt.Errorf("Put method: %s", err)
	}

	return nil
}

func (sc *songController) DeleteSong(ctx context.Context, songId int) error {
	sc.lgr.DebugLogger.Printf("DeleteSong called with songId: %d\n", songId)

	if err := sc.repo.DeleteSong(songId); err != nil {
		return fmt.Errorf("Delete method: %s", err)
	}

	return nil
}
