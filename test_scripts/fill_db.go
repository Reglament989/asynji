package main

import (
	"fmt"

	"github.com/go-bongo/bongo"

	"github.com/Reglament989/asynji/pkgs/asynji/models"
)

func main() {
	config := &bongo.Config{
		ConnectionString: "mongodb://localhost:27017",
		Database:         "bongotest",
	}
	Conn, _ := bongo.Connect(config)
	err := Conn.Session.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Mongo connected!")
	col := Conn.Collection(fmt.Sprintf("%s-%s", "608d2b1a97b7ea0d4aac30ea", "Messages"))
	for i := 1; i <= 100_000; i++ {
		println(i)
		m := &models.Message{
			From: "",
			Body: "Message " + string(rune(i)),
		}
		col.Save(m)
	}
}
