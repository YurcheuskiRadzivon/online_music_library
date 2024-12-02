package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/model"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/repository"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

type SongController interface {
	GetSongs(ctx context.Context, sortParam string, page int, pageSize int) ([]model.Song, error)
	GetSong(ctx context.Context, songId int) (*model.Song, error)
	GetSongText(ctx context.Context, songId int, pageSize int, page int) ([]string, error)
	InsertSong(ctx context.Context, songRequest model.SongRequest) error
	UpdateUser(ctx context.Context, songId int, Song model.Song) error
	DeleteUser(ctx context.Context, songId int) error
}
type songController struct {
	repo repository.SongRepository
}

func NewSongController(repo repository.SongRepository) SongController {
	return &songController{repo: repo}
}
func (sc *songController) GetSongs(ctx context.Context, sortParam string, page int, pageSize int) ([]model.Song, error) {
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
	}

	switch sortParam {
	case "sound_id":
		sort.Slice(songs, func(i, j int) bool {
			return songs[i].SoundId < songs[j].SoundId
		})
	case "text_length":
		sort.Slice(songs, func(i, j int) bool {
			return len(songs[i].Text) < len(songs[j].Text)
		})
	case "song":
		sort.Slice(songs, func(i, j int) bool {
			return strings.ToLower(songs[i].Song) < strings.ToLower(songs[j].Song)
		})
	case "release_date":
		sort.Slice(songs, func(i, j int) bool {
			dateI, errI := time.Parse("02.01.2006", songs[i].ReleaseDate)
			dateJ, errJ := time.Parse("02.01.2006", songs[j].ReleaseDate)
			if errI != nil || errJ != nil {
				return false
			}
			return dateI.Before(dateJ)
		})
	}

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

	return paginatedSongs, nil
}
func (sc *songController) GetSong(ctx context.Context, songId int) (*model.Song, error) {
	song, err := sc.repo.GetSong(songId)
	if err != nil {
		return nil, err
	}
	return song, nil
}
func (sc *songController) GetSongText(ctx context.Context, songId int, pageSize int, page int) ([]string, error) {
	if pageSize < 1 {
		pageSize = 1
	}
	if page < 1 {
		page = 1
	}

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
	return paginatedVerses, nil
}
func (sc *songController) InsertSong(ctx context.Context, songRequest model.SongRequest) error {
	apiUrl := os.Getenv("EXTERNAL_API_URL") + "/info?group=" + songRequest.Group + "&song=" + songRequest.Song
	resp, err := http.Get(apiUrl)
	if err != nil {
		return fmt.Errorf("External api error:%s, status:%v", err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var songDetail model.SongDetail
		err = json.NewDecoder(resp.Body).Decode(&songDetail)
		if err != nil {
			return fmt.Errorf("External api error:%s, status:%v", err, http.StatusInternalServerError)
		}
		song := model.NewSong(songRequest, songDetail)
		if err := sc.repo.InsertSong(song); err != nil {
			return fmt.Errorf("Insert method: %s", err)
		}
		return nil

	} else {
		return fmt.Errorf("External api error: bad request, status:%v", resp.StatusCode)

	}

}
func (sc *songController) UpdateUser(ctx context.Context, songId int, song model.Song) error {
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
func (sc *songController) DeleteUser(ctx context.Context, songId int) error {
	if err := sc.repo.DeleteSong(songId); err != nil {
		return fmt.Errorf("Delete method: %s", err)
	}
	return nil
}
