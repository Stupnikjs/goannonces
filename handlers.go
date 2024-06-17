package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/Stupnikjs/zik/pkg/gstore"
	"github.com/Stupnikjs/zik/pkg/repo"
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
	msg, err := app.LoadMultipartReqToBucket(r, BucketName)
	if err != nil {
		WriteErrorToResponse(w, err, 404)
	}

	w.Write([]byte(fmt.Sprintf("%s", msg)))
}

func ListObjectHandler(w http.ResponseWriter, r *http.Request) {
	names, err := gstore.ListObjectsBucket(BucketName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	byteNames, err := json.Marshal(names)

	if err != nil {
		WriteErrorToResponse(w, err, 404)
	}

	w.Write(byteNames)

}

func (app *Application) LoadMultipartReqToBucket(r *http.Request, bucketName string) (string, error) {
	objNames, err := gstore.ListObjectsBucket(BucketName)

	already := []string{}
	if err != nil {
		return "", err
	}

	for _, headers := range r.MultipartForm.File {

		for _, h := range headers {

			if IsInSlice(h.Filename, objNames) {
				// format a messgage with already present files
				already = append(already, h.Filename)

				break
			}

			file, err := h.Open()

			if err != nil {
				return "", err
			}

			defer file.Close()

			finalByteArr, err := ByteFromMegaFile(file)

			if err != nil {
				return "", err
			}

			err = gstore.LoadToBucket(bucketName, h.Filename, finalByteArr)

			if err != nil {
				return "", err
			}

			track := repo.Track{}
			url, err := gstore.GetObjectURL(bucketName, h.Filename)
			track.StoreURL = url
			track.Name = h.Filename
			track.Selected = false
			err = app.DB.PushTrackToSQL(track)
			if err != nil {
				return "", err
			}
		}

	}
	msg := ""
	for _, s := range already {
		msg += fmt.Sprintf(" %s were alreaddy in bucket ", s)
	}
	return msg, nil

}

func (app *Application) UploadTrackFromGCPHandler(w http.ResponseWriter, r *http.Request) {
	trackid := chi.URLParam(r, "id")
	track, err := app.DB.GetTrackFromId(trackid)
	if err != nil {
		WriteErrorToResponse(w, err, 404)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := storage.NewClient(ctx)
	if err != nil {
		WriteErrorToResponse(w, err, 404)
		return
	}
	bucket := client.Bucket(BucketName)
	obj := bucket.Object(track.Name)
	defer client.Close()
	reader, err := obj.NewReader(ctx)

	defer reader.Close()
	if err != nil {
		WriteErrorToResponse(w, err, 404)
		return
	}
	w.Header().Set("Content-Type", "audio/mpeg")
	w.WriteHeader(http.StatusOK)

	_, _ = io.Copy(w, reader)

}

func (app *Application) DeleteTrackHandler(w http.ResponseWriter, r *http.Request) {
	/*
		reqStruct := reqStruct{}

		bytes, err := io.ReadAll(r.Body)
		err = json.Unmarshal(bytes, reqStruct)
	*/
	trackid := ""
	trackidInt, err := strconv.Atoi(trackid)
	if err != nil {
		WriteErrorToResponse(w, err, 404)
		return
	}
	trackidInt32 := int32(trackidInt)
	err = app.DB.DeleteTrack(trackidInt32)
	if err != nil {
		WriteErrorToResponse(w, err, 404)
		return
	}
	w.Write([]byte("track deleted succefully"))
}

func (app *Application) DeleteGCPObjectHandler(w http.ResponseWriter, r *http.Request) {
	// call to app
	trackid := chi.URLParam(r, "id")
	track, err := app.DB.GetTrackFromId(trackid)
	if err != nil {
		WriteErrorToResponse(w, err, 404)
		return
	}
	err = gstore.DeleteObjectInBucket(BucketName, track.Name)
	if err != nil {
		WriteErrorToResponse(w, err, 404)
		return
	}
	w.Write([]byte(fmt.Sprintf("file %s deleted succesfully in bucker", track.Name)))

}

func (app *Application) IncrementPlayCountHandler(w http.ResponseWriter, r *http.Request) {
	// call to app

}

type tagRequest struct {
	Tag string `json:"tag"`
}

func (app *Application) UpdateTrackTagHandler(w http.ResponseWriter, r *http.Request) {
	// test request content type

	// read request body json
	body, err := io.ReadAll(r.Body)

	if err != nil {
		WriteErrorToResponse(w, err, 404)
		return
	}
	tr := tagRequest{}
	json.Unmarshal(body, &tr)

	if tr.Tag != "" {
		trackid := chi.URLParam(r, "id")
		trackidInt, err := strconv.Atoi(trackid)
		if err != nil {
			WriteErrorToResponse(w, err, 404)
			return
		}
		trackidInt32 := int32(trackidInt)
		err = app.DB.UpdateTrackTag(trackidInt32, tr.Tag)

		if err != nil {
			WriteErrorToResponse(w, err, 404)
			return

		}

	}

	defer r.Body.Close()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("tag updated succefuly"))
}
