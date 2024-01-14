package main

import (
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func createFile(name string, file multipart.File) error {

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

func removeFile(name string) error {
	e := os.Remove(name)
	if e != nil {
		log.Fatal(e)
	}
	return nil
}
