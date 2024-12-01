package repository

import (
	"context"
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SongRepository interface {
	GetSongs() ([]model.Song, error)
	GetSong(songId int) (*model.Song, error)
	InsertSong(User model.Song) error
	UpdateSong(songId int, Song model.Song) error
	DeleteSong(songId int) error
}

type songRepository struct {
	db *pgxpool.Pool
}

func NewSongRepository(dsnStr string) (SongRepository, error) {
	dsn := fmt.Sprintf(dsnStr)
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return &songRepository{db: db}, nil
}
func (sr *songRepository) GetSongs() ([]model.Song, error) {
	return nil, nil
}
func (sr *songRepository) GetSong(songId int) (*model.Song, error) {
	return nil, nil
}
func (sr *songRepository) InsertSong(User model.Song) error {
	return nil
}
func (sr *songRepository) UpdateSong(songId int, Song model.Song) error {

	return nil
}
func (sr *songRepository) DeleteSong(songId int) error {
	return nil
}
