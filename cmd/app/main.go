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

	log.Println("Server is running at 127.0.0.1:8080")
	log.Println("QR generate page 127.0.0.1:8080/generate")
	log.Println("Server mode: default")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(err)
}
