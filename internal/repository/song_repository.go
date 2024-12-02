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
	InsertSong(song model.Song) error
	UpdateSong(songId int, song model.Song) error
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
	var songs []model.Song

	query := `SELECT id, "group", song, release_date, text, link FROM songs`
	rows, err := sr.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song model.Song
		err := rows.Scan(&song.SoundId, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return songs, nil
}
func (sr *songRepository) GetSong(songId int) (*model.Song, error) {
	var song model.Song
	query := `SELECT id, "group", song, release_date, text, link FROM songs WHERE id = $1;`
	err := sr.db.QueryRow(context.Background(), query, songId).Scan(&song.SoundId, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link)
	if err != nil {
		return nil, err
	}
	return &song, nil
}
func (sr *songRepository) InsertSong(song model.Song) error {
	query := `INSERT INTO songs("group", song, release_date, text, link) VALUES ($1, $2, $3, $4, $5);`
	_, err := sr.db.Exec(context.Background(), query, song.Group, song.Song, song.ReleaseDate, song.Text, song.Link)
	if err != nil {
		return err
	}
	return nil
}
func (sr *songRepository) UpdateSong(songId int, song model.Song) error {
	query := `UPDATE songs SET "group"=$1, song=$2, release_date=$3, text=$4, link=$5 WHERE id=$6;;`
	_, err := sr.db.Exec(context.Background(), query, song.Group, song.Song, song.ReleaseDate, song.Text, song.Link, songId)
	if err != nil {
		return err
	}
	return nil
}
func (sr *songRepository) DeleteSong(songId int) error {
	query := `DELETE FROM songs WHERE id=$1;`
	_, err := sr.db.Exec(context.Background(), query, songId)
	if err != nil {
		return err
	}
	return nil
}
