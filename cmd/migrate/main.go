package main

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/config"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/migrator"
	"github.com/joho/godotenv"
)

const migrationsDir = "migrations"

//go:embed migrations/*.sql
var MigrationsFS embed.FS

func init() {

	if err := godotenv.Load(); err != nil {
		////		log.Println("Not found .env file")
	}

}
func main() {
	mgrtr := migrator.NewMigrator(MigrationsFS, migrationsDir)
	conf := config.NewConfig()
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.DB.DB_USER, conf.DB.DB_PASSWORD, conf.DB.DB_HOST, conf.DB.DB_PORT, conf.DB.DB_NAME)
	connection, err := sql.Open("postgres", connectionStr)
	if err != nil {
		////
	}
	defer connection.Close()
	if err = mgrtr.ApplyMigrations(connection); err != nil {
		////
	}

}
