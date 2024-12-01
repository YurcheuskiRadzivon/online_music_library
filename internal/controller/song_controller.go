package controller

import (
	"context"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/model"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/repository"
)

type SongController interface {
	GetSongs(ctx context.Context) (*model.Song, error)
	GetSong(ctx context.Context, songId int) (*model.Song, error)
	InsertSong(ctx context.Context, SongRequest model.SongRequest) error
	UpdateUser(ctx context.Context, songId int, Song model.Song) error
	DeleteUser(ctx context.Context, songId int) error
}
type songController struct {
	repo repository.SongRepository
}

func NewSongController(repo repository.SongRepository) SongController {
	return &songController{repo: repo}
}
func (sc *songController) GetSongs(ctx context.Context) (*model.Song, error) {
	return nil, nil
}
func (sc *songController) GetSong(ctx context.Context, songId int) (*model.Song, error) {
	return nil, nil
}
func (sc *songController) InsertSong(ctx context.Context, SongRequest model.SongRequest) error {
	return nil
}
func (sc *songController) UpdateUser(ctx context.Context, songId int, Song model.Song) error {
	return nil
}
func (sc *songController) DeleteUser(ctx context.Context, songId int) error {
	return nil
}
