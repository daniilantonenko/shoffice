package app

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"slices"
	"time"
)

func render(w http.ResponseWriter, filename string, data interface{}) {
	//tmpl, err := template.ParseFiles("templates/"+filename+".html", "templates/header.html", "templates/footer.html")
	tmpl, err := template.ParseFiles("./web/templates/"+filename+".html", "./web/templates/header.html", "./web/templates/footer.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Sorry, Parse went wrong", http.StatusInternalServerError)
	}

	if err := tmpl.ExecuteTemplate(w, filename, data); err != nil {
		log.Print(err)
		http.Error(w, "Sorry, Execute went wrong", http.StatusInternalServerError)
	}
}

func Confirmation(w http.ResponseWriter, r *http.Request) {
	render(w, "confirmation", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	configMail := ReadConfig("conf.json")
	// configuration.
	render(w, "index", configMail)
}

func Generate(w http.ResponseWriter, r *http.Request) {

	render(w, "generate", map[string][]string{"Addresses": GetIp()})
}

func SendForm(w http.ResponseWriter, r *http.Request) {
	// Checkin POST method
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read config
	configMail := ReadConfig("conf.json")

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
	if err := CreateFile(fileName, fileForm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch configMail.Mode {
	case "email":
		// Checking data and sending a message
		if commentForm != "" && fileForm != nil {
			simpleMail := Mail{
				Name:     "Anonym",
				Message:  commentForm,
				Subject:  "Shared a file with you",
				ToEmail:  configMail.FromEmail,
				FileName: fileName}

			err := simpleMail.SendMail(configMail)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Remove File after sending email
			if err := RemoveFile(fileName); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			log.Println("Not all fields were filled in!")
		}
	default:
		// The files are not sent and remain in the directory
	}

	http.Redirect(w, r, "/confirmation", http.StatusSeeOther)
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// AJAX Request Handler
func AjaxHandler(w http.ResponseWriter, r *http.Request) {
	// Checkin POST method
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Receiving POST data
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	stringBody := string(body)

	if !IsIPv4(stringBody) {
		http.Error(w, "IP is not an IPv4 address.", http.StatusBadRequest)
		return
	}

	ipString := "http://" + stringBody + ":8080/"

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "image/png")

	var base64Encoding string
	base64Encoding += "data:image/png;base64,"
	base64Encoding += toBase64(GenerateQrByte(ipString))

	w.Write([]byte(base64Encoding))
}
