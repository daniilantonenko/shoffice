// Package app configures and runs application.
package app

import (
	"log"
	"net/http"
)

// Run creates objects via constructors.
func Run(cfg *Configuration) {

	serverPort := cfg.ServerPort

	// EMAIL Server
	//es := cfg.EmailServer

	// HTTP Server
	fs := http.FileServer(http.Dir("./web/static"))

	//mux := http.NewServeMux()

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", Index(*cfg))
	http.Handle("/send", SendForm(*cfg))
	http.HandleFunc("/confirmation", Confirmation)
	http.HandleFunc("/generate", Generate)
	http.HandleFunc("/qrcode", AjaxHandler)

	log.Printf("Server is running at 127.0.0.1%s", serverPort)
	log.Printf("QR generate page 127.0.0.1%s/generate", serverPort)

	err := http.ListenAndServe(serverPort, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(err)
}
