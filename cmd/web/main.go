package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

type config struct {
	addr      string
	staticUIDir string
	staticDir string
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr:       getEnv("ADDR", ":3500"),
		staticUIDir: getEnv("STATIC_UI_DIR", "./ui/static"),
		staticDir:  getEnv("STATIC_DIR", "/static"),
	}

    mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(cfg.staticUIDir))
	mux.Handle("GET /static/", http.StripPrefix(cfg.staticDir, fileServer))

    mux.HandleFunc("GET /{$}", home)
    mux.HandleFunc("GET /snippet/view/{id}", snippetView)
    mux.HandleFunc("GET /snippet/create", snippetCreate)
    mux.HandleFunc("POST /snippet/create", snippetCreatePost)

    log.Printf("starting server on %s", cfg.addr)
    
    err = http.ListenAndServe(cfg.addr, mux)
    log.Fatal(err)
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}