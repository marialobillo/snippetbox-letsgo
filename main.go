package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("starting server on :3500")

	err := http.ListenAndServe(":3500", mux)
	log.Fatal(err)
}