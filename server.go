package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net"
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

type PageData struct {
	Images []string
}

func generate(w http.ResponseWriter, r *http.Request) {

	pageData := PageData{
		Images: getIp(),
	}

	render(w, "generate", pageData)
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

func getIp() []string {
	// get list of available addresses
	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var arr []string

	for _, addr := range addr {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// check if IPv4 or IPv6 is not nil
			if ipnet.IP.To4() != nil || ipnet.IP.To16() == nil {
				// print available addresses
				fmt.Println(ipnet.IP.String())
				//arr[i] = "ipnet.IP.String()"
				arr = append(arr, ipnet.IP.String())
			}
		}
	}
	// TODO: return []string
	return arr
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// AJAX Request Handler
func ajaxHandler(w http.ResponseWriter, r *http.Request) {
	// Checkin POST method
	/*if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}*/

	//parse request to struct
	/*var d PageData
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		log.Println("NewDecoder")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// create json response from struct
	a, err := json.Marshal(d)
	if err != nil {
		log.Println("Marshal")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println(a)*/

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "image/png")

	var base64Encoding string
	base64Encoding += "data:image/png;base64,"
	base64Encoding += toBase64(generateQrByte("http://localhost:8080/"))

	w.Write([]byte(base64Encoding))
}

func main() {
	// Initializing the Web Server
	http.HandleFunc("/", index)
	http.HandleFunc("/send", sendForm)
	http.HandleFunc("/confirmation", confirmation)
	http.HandleFunc("/generate", generate)
	http.HandleFunc("/ajax", ajaxHandler)
	http.ListenAndServe(":8080", nil)
}
