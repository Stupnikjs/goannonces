package main

import (
	"html/template"
	"net/http"
	"path"
)

var pathToTemplates = "/static/templates/"

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

	_ = render(w, r, "/acceuil.gohtml", &TemplateData{})
}


func (app *application) RenderSoloTrack(w http.ResponseWriter, r *http.Request) {
  TrackId := request.GetParams()
  
  trackStream = :DownloadFromGcp(TrackId)
  td := TemplateData{}
  td.Data["track"] = trackStream
  _ = render(w, "/trackplayer.gohtml", &TemplateData{}

}
