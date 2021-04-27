package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Reglament989/asynji/pkgs/pusher/api/types"

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
	requestPayload, err := json.Marshal(types.Notification{
		Notifications: []types.NotificationPayload{
			{
				Tokens:   tokens,
				Platform: 2,
				Message:  message,
				Title:    title,
			},
		},
	})
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	resp, err := http.Post(url+"/api/push", "application/json", bytes.NewBuffer(requestPayload))
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return err1
	}
	response := types.Response{}
	err2 := json.Unmarshal(body, &response)
	if err2 != nil {
		return err2
	}
	if response.Success == "ok" {
		return nil
	}
	return &errorString{response.Logs[0]}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
