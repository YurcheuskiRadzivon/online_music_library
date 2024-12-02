package initialization

import (
	"errors"
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/controller"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/handler"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/repository"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
)

func InitializeComponentsSong(dsnStr string, lgr *logger.Logger) (handler.SongHandler, error) {
	songRepo, err := repository.NewSongRepository(dsnStr, lgr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("DB connection has failed: %v", err))
	}
	songController := controller.NewSongController(songRepo, lgr)
	userHandler := handler.NewSongHandler(songController, lgr)
	return userHandler, nil
}
