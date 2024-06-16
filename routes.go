package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Application) routes() http.Handler {

	mux := chi.NewRouter()

	// register routes
	mux.Get("/", app.RenderAccueil)
	mux.Get("/loader", app.RenderLoader)
	mux.Get("/urls", ListObjectHandler)
	mux.Post("/upload", app.UploadFile)
	mux.Get("/stream/sound/{id}", app.UploadTrackFromGCPHandler)

	// static assets

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
