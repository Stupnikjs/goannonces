package api

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/Stupnikjs/zik/gstore"
	"github.com/Stupnikjs/zik/repo"
	"github.com/Stupnikjs/zik/util"
	"github.com/Stupnikjs/zik/ytb"
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

	td := TemplateData{}

	_ = render(w, r, "/dragdrop.gohtml", &td)
}
func (app *Application) RenderYoutube(w http.ResponseWriter, r *http.Request) {

	td := TemplateData{}

	_ = render(w, r, "/youtube.gohtml", &td)
}

func (app *Application) RenderPlaylist(w http.ResponseWriter, r *http.Request) {
	tracks, err := app.DB.GetAllTracks()
	if err != nil {
		WriteErrorToResponse(w, err, http.StatusInternalServerError)
	}
	playlists, err := app.DB.GetAllPlaylists()

	if err != nil {
		WriteErrorToResponse(w, err, http.StatusInternalServerError)
	}
	fmt.Println(playlists)
	bytes, err := json.Marshal(tracks)
	if err != nil {
		WriteErrorToResponse(w, err, http.StatusInternalServerError)
	}

	td := TemplateData{}
	td.Data = make(map[string]any)
	td.Data["Tracks"] = string(bytes)
	td.Data["Playlists"] = playlists

	_ = render(w, r, "/playlist.gohtml", &td)
}

/* Api calls */

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
		WriteErrorToResponse(w, err, 404)
	}

	w.Write([]byte(msg))
}

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

func (app *Application) UploadTrackFromGCPHandler(w http.ResponseWriter, r *http.Request) {

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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := storage.NewClient(ctx)
	if err != nil {
		WriteErrorToResponse(w, err, 404)
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

type ytRequest struct {
	YtID string `json:"ytid"`
}

func (app *Application) YoutubeToGCPHandler(w http.ResponseWriter, r *http.Request) {

	ytReq := ytRequest{}

	bytes, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}
	json.Unmarshal(bytes, &ytReq)

	mp3file, err := ytb.Download(ytReq.YtID)

	if err != nil {
		fmt.Println(err)
		WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}
	// load MP3 filed to bicket

	file, err := os.Open(mp3file)
	if err != nil {
		WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}
	defer file.Close()

	bytes, err = ByteFromMegaFile(file)
	if err != nil {
		WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}

	mp3file = util.ProcessMp3Name(mp3file)
	err = gstore.LoadToBucket(app.BucketName, mp3file, bytes)
	if err != nil {
		WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}
	track := repo.Track{
		Name: mp3file,
	}
	err = app.DB.PushTrackToSQL(track)

	if err != nil {
		WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}

	util.CleanAllTempDir()
	w.Write([]byte("youtube music uploaded on gcp"))

}

/* Playlist */

func (app *Application) CreatePlaylistHandler(w http.ResponseWriter, r *http.Request) {
	reqJson, err := ParseJsonReq(r)

	if err != nil {
		WriteErrorToResponse(w, err, http.StatusBadRequest)
		return
	}
	body, ok := reqJson.Object.Body.(map[string]string)
	if reqJson.Action == "create" && reqJson.Object.Type == "playlist" && ok {
		name := body["name"]
		err = app.DB.CreatePlaylist(name)

		if err != nil {
			WriteErrorToResponse(w, err, http.StatusBadRequest)
			return
		}
	}

}

func (app *Application) AppendToPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	// call to app

	// get playlist id and track id from req

}

func (app *Application) RemoveToPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	// call to app

}
