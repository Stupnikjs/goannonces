package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
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

 mocksTracks = []Tracks{

    {Name : "Jinjo - Always there for u.mp3",
    StorageURL: "",
    },
   {
},
}



	_ = render(w, r, "/acceuil.gohtml", &TemplateData{})
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
/*
	// Create a file in the server to save the uploaded file
	dst, err := os.Create(header.Filename)
	if err != nil {
		http.Error(w, "Unable to create the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the created file on the server
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	// Return a success message
	fmt.Fprintf(w, "File uploaded successfully: %v", header.Filename)
*/
	 err = app.LoadToBucket(header.FileName, buf)

  
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
