package main

import (
	"database/sql"
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/config"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/migrator"
	"github.com/YurcheuskiRadzivon/online_music_library/migration"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
)

const migrationsDir = "sql_files"

var lgr *logger.Logger = logger.NewLogger()

func init() {

	envPath := filepath.Join("..", "..", ".env")
	if err := godotenv.Load(envPath); err != nil {
		lgr.DebugLogger.Println("Not found .env file")
	} else {
		lgr.InfoLogger.Println(".env file was found")
	}

}
func main() {
	defer func() {
		if rec := recover(); rec != nil {
			log.Printf("Caught panic: %v", rec)
		}
	}()

	mgrtr, err := migrator.NewMigrator(migration.MigrationsFS, migrationsDir)
	if err != nil {
		panic(fmt.Errorf("Creating migrator has failed: %s\n", err))
	}
	lgr.InfoLogger.Println("Creating migrator has successfully")

	conf := config.NewConfig()
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.DB.DB_USER, conf.DB.DB_PASSWORD, conf.DB.DB_HOST, conf.DB.DB_PORT, conf.DB.DB_NAME)
	connection, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(fmt.Errorf("Connection has failed: %s\n", err))
	}
	lgr.InfoLogger.Println("Connection has successfully")
	defer connection.Close()
	if err = mgrtr.ApplyMigrations(connection); err != nil {
		panic(fmt.Errorf("Applying migrations has failed: %s\n", err))
	}
	lgr.InfoLogger.Println("Applying migrations has successfully")

}
