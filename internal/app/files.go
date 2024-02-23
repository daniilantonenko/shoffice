package app

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

func CreateFile(name string, file io.Reader) error {

	if name != "" {
		// Create file
		dst, err := os.Create(name)

		if err != nil {
			return errors.New("cannot create file")
		}

		defer dst.Close()

		// Copy the uploaded file to the created file on the filesystem
		if _, err := io.Copy(dst, file); err != nil {
			return errors.New("unable to save file")
		} else {
			log.Println("Successfully saved file")
		}
	} else {
		return errors.New("no name provided")
	}
	return nil
}

func RemoveFile(name string) error {
	e := os.Remove(name)
	if e != nil {
		log.Fatal(e)
	}
	return nil
}

func RootDir() string {
	f, _ := os.Getwd()
	f = filepath.Dir(filepath.Dir(f))
	return f
}
