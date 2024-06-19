package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"

	"github.com/Stupnikjs/zik/internal/gstore"
	"github.com/Stupnikjs/zik/pkg/util"
	"github.com/go-chi/chi/v5"
)

var pathToTemplates = "./static/templates/"

type TemplateData struct {
	Data map[string]any
}

func render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) error {
	_ = r.Method

	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t), path.Join(pathToTemplates, "/base.layout.gohtml"))
	if err != nil {
		return err
	}
	err = parsedTemplate.Execute(w, td)
	if err != nil {
		return err
	}
	return nil

}

// template rendering

func (app *Application) RenderAccueil(w http.ResponseWriter, r *http.Request) {

	td := TemplateData{}
	tracks, err := app.DB.GetAllTracks()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	td.Data = make(map[string]any)
	td.Data["Tracks"] = tracks
	_ = render(w, r, "/acceuil.gohtml", &td)
}

func (app *Application) RenderLoader(w http.ResponseWriter, r *http.Request) {

	_ = render(w, r, "/loader.gohtml", &TemplateData{})
}

func (app *Application) RenderSingleTrack(w http.ResponseWriter, r *http.Request) {
	trackid := chi.URLParam(r, "id")
	td := TemplateData{}
	track, err := app.DB.GetTrackFromId(trackid)
	td.Data = map[string]any{}
	td.Data["Track"] = track
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = render(w, r, "/singletrack.gohtml", &td)
}

func (app *Application) RenderDragDrop(w http.ResponseWriter, r *http.Request) {

	td := TemplateData{}

	_ = render(w, r, "/dragdrop.gohtml", &td)
}

func (app *Application) UploadFile(w http.ResponseWriter, r *http.Request) {
	// load file to gcp bucket

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Retrieve the file from the form data
	msg, err := app.LoadMultipartReqToBucket(r, app.BucketName)
	if err != nil {
		util.WriteErrorToResponse(w, err, 404)
	}

	w.Write([]byte(fmt.Sprintf("%s", msg)))
}

func (app *Application) ListObjectHandler(w http.ResponseWriter, r *http.Request) {
	names, err := gstore.ListObjectsBucket(app.BucketName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	byteNames, err := json.Marshal(names)

	if err != nil {
		util.WriteErrorToResponse(w, err, 404)
	}

	w.Write(byteNames)

}

func (app *Application) IncrementPlayCountHandler(w http.ResponseWriter, r *http.Request) {
	// call to app

}

type tagRequest struct {
	Tag string `json:"tag"`
}
