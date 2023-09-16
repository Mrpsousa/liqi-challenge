package main

import (
	"api/internal/handler"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()
	r.Get("/keys", handler.GetHandler)
	r.Post("/address", handler.PostHandler)

	log.Println("INFO: Server is running on port 8000")
	http.ListenAndServe(":8000", r)
}
