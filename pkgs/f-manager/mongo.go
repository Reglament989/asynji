package main

import (
	"fmt"
	"os"

	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

var Conn *bongo.Connection

type File struct {
	bongo.DocumentBase `bson:",inline"`
	FileId             string
	Owner              string
	AccessTimes        int
	Hash               string
}

func Init() {
	config := &bongo.Config{
		ConnectionString: os.Getenv("MONGO_URI"),
		Database:         "files",
	}
	var err error
	Conn, _ = bongo.Connect(config)
	err = Conn.Session.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Mongo connected!")
}

func SaveFileIndexToDb(fileid string, hash string, owner string) error {
	col := Conn.Collection("Files")
	file := &File{
		FileId:      fileid,
		Owner:       owner,
		AccessTimes: 0,
		Hash:        hash,
	}
	err := col.Save(file)
	return err
}

func CheckByHashIsExists(hash string) (string, error) {
	col := Conn.Collection("Files")
	file := File{}
	err := col.FindOne(bson.M{"hash": hash}, &file)
	if _, ok := err.(*bongo.DocumentNotFoundError); ok {
		return "", nil
	}
	return file.FileId, err
}

func UpdateAcessTimes(fileid string) error {
	col := Conn.Collection("Files")
	return col.Collection().Update(bson.M{"fileid": fileid}, bson.M{"$inc": bson.M{"accesstimes": 1}})
}
