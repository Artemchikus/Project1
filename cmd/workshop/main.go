package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	h := handler.NewHandler(apiClient, cfg.CustomJoke)	

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	path := cfg.Host+":"+cfg.Port

	srv := &http.Server{
		Addr: path,
		Handler: r,
	}

	quit := make(chan os.Signal, 1) 
	done := make(chan error, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func ()  {
		<- quit
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		done <- srv.Shutdown(ctx)
	}()

	log.Printf("Satrting server at %s", path)
	_ = srv.ListenAndServe()

	err = <- done

	log.Printf("Shutting server down with %v:", err)
}