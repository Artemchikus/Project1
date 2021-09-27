package main

import (
	"log"
	"net/http"

	"github.com/Artemchikus/internal/config"
	"github.com/Artemchikus/internal/handler"

	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"
)

func main() {	

	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler()	

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	log.Print("Satrting server:")
	err = http.ListenAndServe(":8080", r)
	log.Fatal(err)

	log.Print("Shutting server down:")
}