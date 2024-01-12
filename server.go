package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"time"
)

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Print(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Print(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
}

func index(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		log.Println(err)
	}
	configMail := readConfig("conf.json")

	tmpl.ExecuteTemplate(w, "index", struct {
		Name string
	}{
		Name: configMail.CompanyName,
	})
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

	//Checking a valid file extension
	fileExt := filepath.Ext(fileHeader.Filename)
	if !slices.Contains(configMail.FileFormats, fileExt) {
		log.Println(fileExt)
		http.Error(w, "A file with an invalid extension.", http.StatusBadRequest)
		return
	}

	fileName := fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), fileExt)

	// Create file
	dst, err := os.Create(fileName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, fileForm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		log.Println("Successfully Uploaded File")
	}

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
	}

	render(w, "templates/confirmation.html", nil)
}

func main() {
	// Initializing the Web Server
	http.HandleFunc("/", index)
	http.HandleFunc("/send", sendForm)
	http.ListenAndServe(":8080", nil)
}
