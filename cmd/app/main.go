package main

import (
	"github.com/YurcheuskiRadzivon/online_music_library/internal/config"
	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(); err != nil {
		////		log.Println("Not found .env file")
	}
}
func main() {
	conf := config.NewConfig()
	_ = conf
}
