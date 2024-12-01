package main

import (
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/config"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/router"
	"github.com/YurcheuskiRadzivon/online_music_library/internal/utils/initialization"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
	"github.com/joho/godotenv"
	"path/filepath"
	"strconv"
)

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
	conf := config.NewConfig()
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.DB.DB_USER, conf.DB.DB_PASSWORD, conf.DB.DB_HOST, conf.DB.DB_PORT, conf.DB.DB_NAME)
	songHandler, err := initialization.InitializeComponentsSong(connectionStr)
	if err != nil {
		lgr.ErrorLogger.Printf("Initialization has failed: %s\n", err)
	}
	app := router.NewFiberRouter(songHandler)
	go func() {
		app.Listen(fmt.Sprintf(":%s", strconv.Itoa(conf.API.API_PORT)))
	}()

}
