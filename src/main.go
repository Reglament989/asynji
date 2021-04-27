package main

import (
	"asynji/src/rdb"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	InitMongo()
	rdb.VerifyRdbConnection()

	r := InitGin()

	r.Run()

}
