package main

import (
	"github.com/YurcheuskiRadzivon/online_music_library/internal/config"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
	"github.com/joho/godotenv"
	"path/filepath"
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
	_ = conf
}
