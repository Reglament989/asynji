package main

import (
	"io/fs"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dgryski/go-metro"
)

// dotFileHidingFile is the http.File use in dotFileHidingFileSystem.
// It is used to wrap the Readdir method of http.File so that we can
// remove files and directories that start with a period from its output.
type dotFileHidingFile struct {
	http.File
}

// Readdir is a wrapper around the Readdir method of the embedded File
// that filters out all files that start with a period in their name.
func (f dotFileHidingFile) Readdir(n int) (fis []fs.FileInfo, err error) {
	files, err := f.File.Readdir(n)
	for _, file := range files { // Filters out the dot files
		if !strings.HasPrefix(file.Name(), ".") {
			fis = append(fis, file)
		}
	}
	return
}

// dotFileHidingFileSystem is an http.FileSystem that hides
// hidden "dot files" from being served.
type dotFileHidingFileSystem struct {
	http.FileSystem
}

// Open is a wrapper around the Open method of the embedded FileSystem
// that serves a 403 permission error when name has a file or directory
// with whose name starts with a period in its path.
func (fsys dotFileHidingFileSystem) Open(name string) (http.File, error) {
	file, err := fsys.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}
	return dotFileHidingFile{file}, err
}

func SaveToFile(data []byte, size int64, dst string) error {
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func ReadMulti(file *multipart.FileHeader) (string, []byte, error) {
	src, err := file.Open()
	if err != nil {
		return "", nil, err
	}
	defer src.Close()
	bytes := make([]byte, file.Size)
	_, err = src.Read(bytes)
	if err != nil {
		return "", nil, err
	}
	hash := metro.Hash64(bytes, 1)
	// log.Println(hash)
	return strconv.FormatUint(hash, 20), bytes, nil
}
