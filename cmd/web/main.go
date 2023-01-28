package main

import (
	"log"
	"net/http"

	"trinity.jakks.net/Administrator/send-it/cmd/web/config"
	"trinity.jakks.net/Administrator/send-it/cmd/web/pkg/handlers"
	"trinity.jakks.net/Administrator/send-it/cmd/web/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
