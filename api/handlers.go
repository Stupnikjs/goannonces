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
	// playlists, err := app.DB.GetAllPlaylists()
	var playlists []int

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
	req := JsonReq{}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(bytes, &req)
	if err != nil {
		fmt.Println(err)
		WriteErrorToResponse(w, err, 404)
		return
	}
	fmt.Println(req)
	trackid := req.Object.Id
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
	// test request content type

	// read request body json
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		WriteErrorToResponse(w, err, 404)
		return
	}
	req := JsonReq{}
	err = json.Unmarshal(body, &req)

	if err != nil {
		fmt.Println(err)
		WriteErrorToResponse(w, err, 404)
		return
	}

	if req.Object.Field == "tag" && req.Object.Type == "track" && req.Action == "update" {

		trackidInt, err := strconv.Atoi(req.Object.Id)
		if err != nil {
			WriteErrorToResponse(w, err, 404)
			return
		}
		tag, ok := req.Object.Body.(string)
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

func (app *Application) LoadMultipartReqToBucket(r *http.Request, bucketName string) (string, error) {
	objNames, err := gstore.ListObjectsBucket(app.BucketName)

	already := []string{}
	if err != nil {
		return "", err
	}

	for _, headers := range r.MultipartForm.File {

		for _, h := range headers {

			if util.IsInSlice(h.Filename, objNames) {
				// format a messgage with already present files
				already = append(already, h.Filename)

				break
			}

			file, err := h.Open()

			if err != nil {
				return "", err
			}

			defer file.Close()

			finalByteArr, err := util.ByteFromMegaFile(file)

			if err != nil {
				return "", err
			}

			err = gstore.LoadToBucket(bucketName, h.Filename, finalByteArr)

			if err != nil {
				return "", err
			}

			track := repo.Track{}
			url, err := gstore.GetObjectURL(bucketName, h.Filename)

			if err != nil {
				return "", err
			}

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

type ytRequest struct {
	YtID string `json:"ytid"`
}

func (app *Application) YoutubeToGCPHandler(w http.ResponseWriter, r *http.Request) {

	ytReq := ytRequest{}

	bytes, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		riteErrorToResponse(w, err, http.StatusInternalServerError)
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

	bytes, err = util.ByteFromMegaFile(file)
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
	reqJson := JsonReq{}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		WriteErrorToResponse(w, err, http.StatusBadRequest)
	}

	r.Body.Close()

	err = json.Unmarshal(bytes, &reqJson)

	if err != nil {
		WriteErrorToResponse(w, err, http.StatusBadRequest)
	}

	body, ok := reqJson.Object.Body.(map[string]string)
	if reqJson.Action == "create" && reqJson.Object.Type == "playlist" && ok {
		name := body["name"]
		app.DB.CreatePlaylist(name)

	}

}

func (app *Application) AppendToPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	// call to app

	// get playlist id and track id from req

}

func (app *Application) RemoveToPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	// call to app

}
