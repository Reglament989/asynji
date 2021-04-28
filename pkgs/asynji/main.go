package main

import (
	"github.com/Reglament989/asynji/pkgs/asynji/rdb"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	rdb.VerifyRdbConnection()

	r := InitGin()

	r.Run()

}
