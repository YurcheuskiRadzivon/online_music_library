package repository

import (
	"context"
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/model"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
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
	db  *pgxpool.Pool
	lgr *logger.Logger
}

func NewSongRepository(dsnStr string, lgr *logger.Logger) (SongRepository, error) {
	dsn := fmt.Sprintf(dsnStr)
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		lgr.ErrorLogger.Println("Failed to connect to the database:", err)
		return nil, err
	}
	lgr.InfoLogger.Println("SongRepository created successfully.")
	return &songRepository{
		db:  db,
		lgr: lgr,
	}, nil
}
func (sr *songRepository) GetSongs() ([]model.Song, error) {
	sr.lgr.DebugLogger.Println("Getting all songs from the database.")
	var songs []model.Song
	query := `SELECT id, "group", song, release_date, text, link FROM songs`
	rows, err := sr.db.Query(context.Background(), query)
	if err != nil {
		sr.lgr.ErrorLogger.Println("Error querying songs:", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var song model.Song
		err := rows.Scan(&song.SoundId, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link)
		if err != nil {
			sr.lgr.ErrorLogger.Println("Error scanning song row:", err)
			return nil, err
		}
		songs = append(songs, song)
	}
	if rows.Err() != nil {
		sr.lgr.ErrorLogger.Println("Row iteration error:", rows.Err())
		return nil, rows.Err()
	}
	sr.lgr.InfoLogger.Printf("Retrieved %d songs from the database.\n", len(songs))
	return songs, nil
}
func (sr *songRepository) GetSong(songId int) (*model.Song, error) {

	var song model.Song
	query := `SELECT id, "group", song, release_date, text, link FROM songs WHERE id = $1;`
	err := sr.db.QueryRow(context.Background(), query, songId).Scan(&song.SoundId, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link)
	if err != nil {
		sr.lgr.ErrorLogger.Printf("Error querying song with ID %d: %v\n", songId, err)
		return nil, err
	}
	sr.lgr.InfoLogger.Printf("Retrieved song with ID %d.\n", songId)
	return &song, nil
}
func (sr *songRepository) InsertSong(song model.Song) error {
	sr.lgr.DebugLogger.Printf("Inserting song: %+v\n", song)
	query := `INSERT INTO songs("group", song, release_date, text, link) VALUES ($1, $2, $3, $4, $5);`
	_, err := sr.db.Exec(context.Background(), query, song.Group, song.Song, song.ReleaseDate, song.Text, song.Link)
	if err != nil {
		sr.lgr.ErrorLogger.Printf("Error inserting song %+v: %v\n", song, err)
		return err
	}
	sr.lgr.InfoLogger.Printf("Inserted song with ID %d.\n", song.SoundId)
	return nil
}
func (sr *songRepository) UpdateSong(songId int, song model.Song) error {
	sr.lgr.DebugLogger.Printf("Updating song with ID %d: %+v\n", songId, song)
	query := `UPDATE songs SET "group"=$1, song=$2, release_date=$3, text=$4, link=$5 WHERE id=$6;`
	_, err := sr.db.Exec(context.Background(), query, song.Group, song.Song, song.ReleaseDate, song.Text, song.Link, songId)
	if err != nil {
		sr.lgr.ErrorLogger.Printf("Error updating song with ID %d: %v\n", songId, err)
		return err
	}
	sr.lgr.InfoLogger.Printf("Updated song with ID %d.\n", songId)
	return nil
}
func (sr *songRepository) DeleteSong(songId int) error {
	sr.lgr.DebugLogger.Printf("Deleting song with ID %d from the database.\n", songId)
	query := `DELETE FROM songs WHERE id=$1;`
	_, err := sr.db.Exec(context.Background(), query, songId)
	if err != nil {
		sr.lgr.ErrorLogger.Printf("Error deleting song with ID %d: %v\n", songId, err)
		return err
	}
	sr.lgr.InfoLogger.Printf("Deleted song with ID %d.\n", songId)
	return nil
}
