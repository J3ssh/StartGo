package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	//First basic GoLang webserver

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, World!") //displays on the page /hello

	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
