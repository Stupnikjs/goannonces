package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ReqModel struct {
	Action string `json:"action"`
	Object ApiObject `json:"object"`
 Field string  `json:"field"`
	Body   string `json:"body"`
}

type ApiObject {
  Name `json:"name"`
  Id `json:"id"`

}

func (app *Application) routes() http.Handler {

	mux := chi.NewRouter()

	// register routes
	mux.Get("/", app.RenderAccueil)
	mux.Get("/urls", ListObjectHandler)
	mux.Post("/upload", app.UploadFile)
	mux.Get("/stream/sound/{id}", app.UploadTrackFromGCPHandler)
	mux.Post("/track/tag/{id}", app.UpdateTrackTagHandler)
	mux.Post("/track/delete", app.DeleteTrackHandler)
	mux.Get("/loader", app.RenderDragDrop)

	mux.Post("/playlist/create", app.CreatePlaylistHandler)
	mux.Post("/playlist/append", app.AppendToPlaylistHandler)
	mux.Post("/playlist/remove", app.RemoveToPlaylistHandler)

	// static assets

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
