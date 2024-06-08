package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"path"

	"cloud.google.com/go/storage"
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

func (app *application) RenderLoader(w http.ResponseWriter, r *http.Request) {

	_ = render(w, r, "/loader.gohtml", &TemplateData{})
}

func (app *application) UploadFile(w http.ResponseWriter, r *http.Request) {
	// load file to gcp bucket
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println(err)
	}
	buck := client.Bucket("lastbucketnamethatsit")
	// err = CreateBucket(client, buck, ctx)
	defer client.Close()
	if err != nil {
		fmt.Println(err)
	}

	// err = PushFileToBucket(buck)
	if err != nil {
		fmt.Println(err)
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
