package main

import (
	"net/http"

	"trinity.jakks.net/Administrator/send-it/cmd/web/pkg/handlers"
)

const portNumber = ":8080"




func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(portNumber, nil)

}


