package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"trinity.jakks.net/Administrator/send-it/cmd/web/config"
	"trinity.jakks.net/Administrator/send-it/cmd/web/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))

	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}