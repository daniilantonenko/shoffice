package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"time"
)

type Config struct {
	EmailServer   string
	EmailPort     int
	FromEmail     string
	FromPass      string
	CompanyName   string
	FileFormats   []string
	MaxUploadSize int64
	Mode          string
}

func readConfig(fileName string) (config Config) {
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

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles("templates/"+filename+".html", "templates/header.html", "templates/footer.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Sorry, Parse went wrong", http.StatusInternalServerError)
	}

	if err := tmpl.ExecuteTemplate(w, filename, data); err != nil {
		log.Print(err)
		http.Error(w, "Sorry, Execute went wrong", http.StatusInternalServerError)
	}
}

func confirmation(w http.ResponseWriter, r *http.Request) {
	render(w, "confirmation", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	configMail := readConfig("conf.json")
	render(w, "index", configMail)
}

func generate(w http.ResponseWriter, r *http.Request) {
	render(w, "generate", generateQr("http://localhost:8080/"))
}

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

func sendForm(w http.ResponseWriter, r *http.Request) {
	// Checkin POST method
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read config
	configMail := readConfig("conf.json")

	// Checking a max upload size
	maxUploadSize := configMail.MaxUploadSize // 10 MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxUploadSize))
	err := r.ParseMultipartForm(int64(maxUploadSize))
	if err != nil {
		log.Println(err)
		http.Error(w, "A file with an invalid size.", http.StatusBadRequest)
		return
	}

	// Receiving POST data
	commentForm := r.FormValue("inputComment")

	// Read file
	fileForm, fileHeader, err := r.FormFile("inputFile")
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to process form.", http.StatusInternalServerError)
		return
	}
	defer fileForm.Close()

	// Checking a valid file extension
	fileExt := filepath.Ext(fileHeader.Filename)
	if !slices.Contains(configMail.FileFormats, fileExt) {
		log.Println(fileExt)
		http.Error(w, "A file with an invalid extension.", http.StatusBadRequest)
		return
	}

	fileName := fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), fileExt)

	// Create File
	if err := createFile(fileName, fileForm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch configMail.Mode {
	case "email":
		// Checking data and sending a message
		if commentForm == "" || fileForm == nil {
			log.Println("Not all fields were filled in!")
		} else {
			simpleMail := Mail{
				Name:     "Anonym",
				Message:  commentForm,
				Subject:  "Shared a file with you",
				ToEmail:  configMail.FromEmail,
				FileName: fileName}

			sendMail(simpleMail, configMail)

			// Remove File after sending email
			if err := removeFile(fileName); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	default:
		// The files are not sent and remain in the directory
	}

	http.Redirect(w, r, "/confirmation", http.StatusSeeOther)
}

func main() {
	// Initializing the Web Server
	http.HandleFunc("/", index)
	http.HandleFunc("/send", sendForm)
	http.HandleFunc("/confirmation", confirmation)
	http.HandleFunc("/generate", generate)
	http.ListenAndServe(":8080", nil)
}
