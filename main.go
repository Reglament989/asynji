package main

import (
	"fmt"
	"gin_msg/ws"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
)
var (
	g errgroup.Group
)

func main() {
	err := godotenv.Load()
  if err != nil {
    panic("Error loading .env file")
  }
	InitMongo()
	
	r := InitGin()

	g.Go(func () error{
		h := ws.InitWs()
		s := &http.Server{
			Addr:         ":8081",
			Handler:      h,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		fmt.Println("\033[32mWs now online on 8081!\033[39m")
		err := s.ListenAndServe()
		return err
	})

	g.Go(func () error {
		s := &http.Server{
			Addr:         ":8080",
			Handler:      r,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		fmt.Println("\033[32mBackend now online on 8080!\033[39m")
		err := s.ListenAndServe()
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
	
}
