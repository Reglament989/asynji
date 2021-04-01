package main

import (
	"fmt"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo() {
	// Setup the mgm default config 
	if err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout:12 * time.Second}, "mgm_lab", options.Client().ApplyURI("mongodb://localhost:27017")); err != nil {
		panic(fmt.Sprintf("Mongo not connected. %s", err.Error()))
	} else {
		fmt.Println("\033[32mMongo has connected!\033[39m")
	}
}