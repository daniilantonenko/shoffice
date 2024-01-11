package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func readConfig(fileName string) (config ConfigMail) {
	// TODO: return a generic structure

	file, _ := os.Open(fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}

	return config
}

func index(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "index", nil)
}

func send(w http.ResponseWriter, r *http.Request) {
	emailForm := r.FormValue("inputEmail")
	commentForm := r.FormValue("inputComment")
	fileForm := r.FormValue("inputFile")

	if emailForm == "" || commentForm == "" || fileForm == "" {
		fmt.Fprintf(w, "Не все данные заполнены!")
	} else {
		fmt.Println(emailForm, commentForm, fileForm)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func handleRequest() {
	http.HandleFunc("/", index)
	http.HandleFunc("/send", send)
	http.ListenAndServe(":8080", nil)
}

func main() {
	// Initializing the Web Server
	handleRequest()

	configMail := readConfig("conf.json")
	fmt.Println(configMail)
	//sendMail(simpleMail, configMail)

}
