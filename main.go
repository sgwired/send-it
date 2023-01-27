package main

import (
	"net/http"
)

const portNumber = ":8080"




func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(portNumber, nil)

}


