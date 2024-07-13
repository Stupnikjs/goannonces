package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Application) Routes() http.Handler {

	mux := chi.NewRouter()

	// register routes
	mux.Get("/", app.RenderAccueil)
	mux.Get("/loader", app.RenderDragDrop)
	mux.Get("/youtube", app.RenderYoutube)

	mux.Get("/api/allobjects", app.ListObjectHandler)

	mux.Post("/upload", app.UploadFile)
	mux.Get("/stream/sound/{id}", app.StreamTrackFromGCPHandler)
	mux.Post("/api/track/set/artist", app.SetArtistHandler)
	mux.Post("/api/track/get/artist/suggestion", app.GetArtistSuggestionHandler)
	mux.Post("/api/track/upload", app.UploadTrackListHandler)
	mux.Post("/api/track/tag", app.UpdateTrackTagHandler)
	mux.Post("/api/track/remove", app.DeleteTrackHandler)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
