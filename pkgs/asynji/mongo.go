package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo() {
	// Setup the mgm default config 
	if err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout:12 * time.Second}, "mgm_lab", options.Client().ApplyURI(os.Getenv("MONGO_URI"))); err != nil {
		panic(fmt.Sprintf("Mongo not connected. %s\nGetten %s", err.Error(), os.Getenv("MONGO_URI")))
	} else {
		fmt.Println("\033[32mMongo has connected!\033[39m")
	}
}