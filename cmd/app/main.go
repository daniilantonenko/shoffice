package main

import (
	"log"
	"net/http"
	"shoffice/internal/app"
)

func main() {
	// Initializing the Web Server
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", app.Index)
	http.HandleFunc("/send", app.SendForm)
	http.HandleFunc("/confirmation", app.Confirmation)
	http.HandleFunc("/generate", app.Generate)
	http.HandleFunc("/qrcode", app.AjaxHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(err)
}
