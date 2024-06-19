package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type JsonReq struct {
	Action string    `json:"action"`
	Object ApiObject `json:"object"`
}

type ApiObject struct {
	Type  string `json:"type"`
	Id    string `json:"id"`
	Body  string `json:"body,omitempty"`
	Field string `json:"field,omitempty"`
}

func (app *Application) routes() http.Handler {

	mux := chi.NewRouter()

	// register routes
	mux.Get("/", app.RenderAccueil)
	mux.Get("/loader", app.RenderDragDrop)

	mux.Get("/urls", ListObjectHandler)

	mux.Post("/upload", app.UploadFile)
	mux.Get("/stream/sound/{id}", app.UploadTrackFromGCPHandler)

	mux.Post("/api/track/tag", app.UpdateTrackTagHandler)
	mux.Post("/api/track/remove", app.DeleteTrackHandler)

	mux.Post("/api/playlist/create", app.CreatePlaylistHandler)
	mux.Post("/api/playlist/append", app.AppendToPlaylistHandler)
	mux.Post("/api/playlist/remove", app.RemoveToPlaylistHandler)

	// static assets

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
