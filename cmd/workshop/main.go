package main

import (
	"log"
	"net/http"

	"github.com/Artemchikus/internal/api/jokes"
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

	apiClient := jokes.NewJokeClient(cfg.JokeURL)

	h := handler.NewHandler(apiClient)	

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	path := cfg.Host+":"+cfg.Port
	log.Printf("Satrting server at %s", path)
	err = http.ListenAndServe(path, r)
	log.Fatal(err)

	log.Print("Shutting server down:")
}