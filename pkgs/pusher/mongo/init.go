package mongo

import (
	"fmt"
	"os"

	"github.com/go-bongo/bongo"
)

var Conn *bongo.Connection

func init() {
	config := &bongo.Config{
		ConnectionString: os.Getenv("MONGO_URI"),
		Database:         "bongotest",
	}
	var err error
	Conn, _ = bongo.Connect(config)
	err = Conn.Session.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Mongo connected!")
}
