package api

import (
	"os"

	"github.com/Reglament989/asynji/pkgs/pusher/fire"

	"github.com/joho/godotenv"
)

var url string

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	url = os.Getenv("GORUSH_URL")
	if url == "" {
		panic("Cannot find GORUSH_URL")
	}
}

func SendNotify(tokens []string, message string, title string) error {
	fire.PushNotify(map[string]string{
		"sd": "sad",
	}, tokens)

	return nil
}
