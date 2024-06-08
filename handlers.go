package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path"
)

var pathToTemplates = "./static/templates/"

type TemplateData struct {
	Data map[string]any
}

func render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) error {

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

func (app *application) RenderAccueil(w http.ResponseWriter, r *http.Request) {

	mocksTracks := []Track{

		{Name: "Jinjo - Always there for u.mp3",
			StoreURL: "erpgiregip",
		},
		{
			Name:     "some track name",
			StoreURL: "urlejfefzkn.com",
		},
	}
	td := TemplateData{}
	td.Data["tracks"] = mocksTracks
	_ = render(w, r, "/acceuil.gohtml", &td)
}

func (app *application) RenderLoader(w http.ResponseWriter, r *http.Request) {

	_ = render(w, r, "/loader.gohtml", &TemplateData{})
}

func (app *application) UploadFile(w http.ResponseWriter, r *http.Request) {
	// load file to gcp bucket

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Retrieve the file from the form data
	file, header, err := r.FormFile("uploadfile")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	buf, err := io.ReadAll(file)

	fmt.Fprintf(w, "File uploaded successfully: %v", header.Filename)

	err = app.LoadToBucket(header.Filename, buf)
	if err != nil {
		http.Error(w, "Error Loading file to BUCKET", http.StatusInternalServerError)
		return
	}
}

/*
func (app *application) RenderSoloTrack(w http.ResponseWriter, r *http.Request) {
  TrackId := request.GetParams()

  trackStream := DownloadFromGcp(TrackId)
  td := TemplateData{}
  td.Data["track"] = trackStream
  _ = render(w, "/trackplayer.gohtml", &TemplateData{})

}
*/
