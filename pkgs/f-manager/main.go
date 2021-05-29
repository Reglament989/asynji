package main

import (
	"log"
	"net/http"

	"github.com/Reglament989/asynji/pkgs/asynji/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

const fileDir = "./file-data/"

var _fs http.FileSystem = dotFileHidingFileSystem{http.Dir(fileDir)}

func main() {
	Init()
	r := gin.Default()
	r.Use(middlewares.Auth())
	r.MaxMultipartMemory = 512 << 20 // 512MB max size to upload
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "Cannot get file from form",
				"error":  err.Error(),
			})
			return
		}
		fileHash, bytes, err := ReadMulti(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  err.Error(),
				"status": "Error on readMulti",
			})
			return
		}
		fileId, err := CheckByHashIsExists(fileHash)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  err.Error(),
				"status": "Error on checkHash",
			})
			return
		}
		if fileId != "" {
			c.JSON(200, gin.H{
				"id":     fileId,
				"status": "Success",
			})
			return
		}
		fileId = xid.New().String()
		filename := fileDir + fileId
		err = SaveFileIndexToDb(fileId, fileHash, c.GetString("userId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  err.Error(),
				"status": "Error on SaveFileIndex",
			})
			return
		}
		SaveToFile(bytes, file.Size, filename)
		c.JSON(201, gin.H{
			"id":     fileId,
			"status": "Success",
		})
	})
	r.GET("/_/:fileid", func(c *gin.Context) {
		fileid := c.Param("fileid")
		c.FileFromFS(fileid, _fs)
		err := UpdateAcessTimes(fileid)
		if err != nil {
			log.Println(err)
		}
	})
	r.Run(":51500")
}
