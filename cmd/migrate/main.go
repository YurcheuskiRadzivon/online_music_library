package main

import (
	"github.com/YurcheuskiRadzivon/online_music_library/internal/config"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	infoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	_ = debugLogger

	if err := godotenv.Load(); err != nil {
		infoLogger.Println("Not found .env file")
	}
}
func main() {
	conf := config.NewConfig()
	_ = conf

}
