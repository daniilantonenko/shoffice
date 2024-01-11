package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	/*simpleMail := Mail{
	Name:    "da",
	Message: "Simple first test message.",
	Subject: "Simple",
	ToEmail: "daniil454122922@gmail.com"}
	*/
	//configMail := readConfig("conf.json")

	tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	tmpl.ExecuteTemplate(w, "index", nil)

	//fmt.Println(simpleMail, configMail)

	//fmt.Fprintf(w, "Name: "+simpleMail.Name)
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
	//http.Handle("/static/", http.StripPrefix("/static/"), http.FileServer(http.Dir("./css/")))
	http.HandleFunc("/", index)
	http.HandleFunc("/send", send)
	http.ListenAndServe(":8080", nil)
}

func main() {
	// Initializing the Web Server
	handleRequest()

	/*simpleMail := Mail{
		Name:    "da",
		Message: "Simple first test message.",
		Subject: "Simple"}

	configMail := readConfig("conf.json")

	fmt.Println(simpleMail, configMail)*/
	//sendMail(simpleMail, configMail)
}
