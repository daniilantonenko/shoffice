package app

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func CreateFile(name string, file multipart.File) error {

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

func ReadConfig(fileName string) (config Config) {
	// TODO: return a generic structure

	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)
	if err != nil {
		log.Println("error:", err)
	}

	return config
}
