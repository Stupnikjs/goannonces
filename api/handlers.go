package api

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/Stupnikjs/zik/gstore"
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

func (app *Application) RenderDragDrop(w http.ResponseWriter, r *http.Request) {
	var TracksNames = []string{}
	td := TemplateData{}
	TrackList, err := app.DB.GetAllTracks()
	if err != nil {
		WriteErrorToResponse(w, err, http.StatusInternalServerError)
	}
	for _, tr := range TrackList {
		TracksNames = append(TracksNames, tr.Name)
	}
	bytes, err := json.Marshal(TracksNames)

	if err != nil {
		WriteErrorToResponse(w, err, http.StatusInternalServerError)
	}
	td.Data = map[string]any{}
	td.Data["Tracks"] = string(bytes)
	_ = render(w, r, "/dragdrop.gohtml", &td)
}
func (app *Application) RenderYoutube(w http.ResponseWriter, r *http.Request) {

	td := TemplateData{}

	_ = render(w, r, "/youtube.gohtml", &td)
}

/* Api calls */

func (app *Application) ListObjectHandler(w http.ResponseWriter, r *http.Request) {
	names, err := gstore.ListObjectsBucket(app.BucketName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	byteNames, err := json.Marshal(names)

	if err != nil {
		WriteErrorToResponse(w, err, 404)
	}

	w.Write(byteNames)

}

func (app *Application) StreamTrackFromGCPHandler(w http.ResponseWriter, r *http.Request) {

	trackid := chi.URLParam(r, "id")
	trackidInt, err := strconv.Atoi(trackid)

	if err != nil {
		WriteErrorToResponse(w, err, 404)
	}

	track, err := app.DB.GetTrackFromId(trackidInt)
	if err != nil {
		WriteErrorToResponse(w, err, 404)
		return
	}

	// to refactor in gstorage ReaderFromObjName

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := storage.NewClient(ctx)
	if err != nil {
		WriteErrorToResponse(w, err, 500)
		return
	}

	bucket := client.Bucket(app.BucketName)
	obj := bucket.Object(track.Name)
	defer client.Close()

	reader, err := obj.NewReader(ctx)

	if err != nil {
		WriteErrorToResponse(w, err, 404)
		return
	}
	defer reader.Close()

	w.Header().Set("Content-Type", "audio/mpeg")
	w.WriteHeader(http.StatusOK)

	_, _ = io.Copy(w, reader)

}

/* Tracks */

func (app *Application) DeleteTrackHandler(w http.ResponseWriter, r *http.Request) {
	jsonReq, err := ParseJsonReq(r)
	if err != nil {
		WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}
	trackid := jsonReq.Object.Id
	trackidInt, err := strconv.Atoi(trackid)
	if err != nil {
		WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}

	track, err := app.DB.GetTrackFromId(trackidInt)

	if err != nil {
		WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}

	err = app.DB.DeleteTrack(trackidInt)
	if err != nil {
		WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}

	err = gstore.DeleteObjectInBucket(app.BucketName, track.Name)
	if err != nil {
		WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("track deleted succefully"))
}

func (app *Application) UpdateTrackTagHandler(w http.ResponseWriter, r *http.Request) {
	jsonReq, err := ParseJsonReq(r)

	if err != nil {
		WriteErrorToResponse(w, err, 404)
		return
	}

	if jsonReq.Object.Field == "tag" && jsonReq.Object.Type == "track" && jsonReq.Action == "update" {

		trackidInt, err := strconv.Atoi(jsonReq.Object.Id)
		if err != nil {
			WriteErrorToResponse(w, err, 404)
			return
		}
		tag, ok := jsonReq.Object.Body.(string)
		if !ok {
			WriteErrorToResponse(w, fmt.Errorf("body malformed"), 404)
			return
		}
		err = app.DB.UpdateTrackTag(trackidInt, tag)

		if err != nil {
			WriteErrorToResponse(w, err, 404)
			return

		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("tag updated succefuly"))
		return

	}
	WriteErrorToResponse(w, fmt.Errorf("wrong request format %s", ""), http.StatusBadRequest)

}

func (app *Application) UploadTrackListHandler(w http.ResponseWriter, r *http.Request) {
	jsonReq, err := ParseJsonReq(r)

	if err != nil {
		WriteErrorToResponse(w, err, 404)
		return
	}

	if jsonReq.Action == "upload" && jsonReq.Object.Type == "tracklist" {
		tracklist, ok := jsonReq.Object.Body.([]string)
		if ok {
			for _, trackName := range tracklist {
				fmt.Println(trackName)
			}
		}
	}

}
func (app *Application) SetArtistHandler(w http.ResponseWriter, r *http.Request) {
	jsonReq, err := ParseJsonReq(r)

	if err != nil {
		WriteErrorToResponse(w, err, 404)
		return
	}

	if jsonReq.Action == "update" && jsonReq.Object.Type == "track" {
		artist, ok := jsonReq.Object.Body.(string)
		id, err := strconv.Atoi(jsonReq.Object.Id)
		if err != nil {
			WriteErrorToResponse(w, err, 404)
			return

		}
		if ok {
			err = app.DB.UpdateTrackArtist(id, artist)
			if err != nil {
				WriteErrorToResponse(w, err, 404)
				return

			}
			w.Write([]byte("artist updated"))
		}
	}

}

func (app *Application) GetArtistSuggestionHandler(w http.ResponseWriter, r *http.Request) {
	jsonReq, err := ParseJsonReq(r)

	if err != nil {
		WriteErrorToResponse(w, err, 404)
		return
	}
	if jsonReq.Action == "artist_suggestion" {
		title, ok := jsonReq.Object.Body.(string)
		if ok {
			suggestions, err := getArtistSuggestion(title)
			if err != nil {
				WriteErrorToResponse(w, err, 404)
				return
			}
			w.Write(suggestions)
			return
		}

	}

}

func getArtistSuggestion(title string) ([]byte, error) {
	strArr := strings.Split(title, "-")
	bytes, err := json.Marshal(strArr)
	if err != nil {
		return nil, err
	}
	return bytes, nil

}
