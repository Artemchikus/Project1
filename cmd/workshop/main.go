package main

import (
	"log"
	"net/http"

	"github.com/Artemchikus/internal/handler"
	"github.com/go-chi/chi"
)

func main() {	

	h := handler.NewHandler()	

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	log.Print("Satrting server:")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)

	log.Print("Shutting server down:")
}