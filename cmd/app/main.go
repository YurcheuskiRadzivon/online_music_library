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
	defer func() {
		if rec := recover(); rec != nil {
			lgr.ErrorLogger.Printf("Caught panic: %v", rec)
		}
	}()
	conf := config.NewConfig()
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.DB.DB_USER, conf.DB.DB_PASSWORD, conf.DB.DB_HOST, strconv.Itoa(conf.DB.DB_PORT), conf.DB.DB_NAME)
	songHandler, err := initialization.InitializeComponentsSong(connectionStr)
	if err != nil {
		panic(fmt.Errorf("Initialization has failed: %s\n", err))
	}
	lgr.InfoLogger.Println("Initialization components for router has successfully")
	app := router.NewFiberRouter(songHandler)
	lgr.DebugLogger.Println("Launching the application.....")
	app.Listen(fmt.Sprintf(":%s", strconv.Itoa(conf.API.API_PORT)))

}
