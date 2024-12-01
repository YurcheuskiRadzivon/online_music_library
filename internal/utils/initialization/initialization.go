package initialization

import (
	"errors"
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/controller"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/handler"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/repository"
)

func InitializeComponentsSong(dsnStr string) (handler.SongHandler, error) {
	songRepo, err := repository.NewSongRepository(dsnStr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("DB connection has failed: %v", err))
	}
	songController := controller.NewSongController(songRepo)
	userHandler := handler.NewSongHandler(songController)
	return userHandler, nil
}
