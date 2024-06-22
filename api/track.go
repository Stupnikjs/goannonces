package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/Stupnikjs/zik/gstore"
	"github.com/Stupnikjs/zik/repo"
	"github.com/Stupnikjs/zik/util"
	"github.com/Stupnikjs/zik/ytb"
	"github.com/go-chi/chi/v5"
)

func (app *Application) UploadTrackFromGCPHandler(w http.ResponseWriter, r *http.Request) {
	trackid := chi.URLParam(r, "id")

	trackidInt, err := strconv.Atoi(trackid)
	if err != nil {
		util.WriteErrorToResponse(w, err, 404)
	}

	track, err := app.DB.GetTrackFromId(int32(trackidInt))
	if err != nil {
		util.WriteErrorToResponse(w, err, 404)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := storage.NewClient(ctx)
	if err != nil {
		util.WriteErrorToResponse(w, err, 404)
		return
	}
	bucket := client.Bucket(app.BucketName)
	obj := bucket.Object(track.Name)
	defer client.Close()
	reader, err := obj.NewReader(ctx)

	defer reader.Close()
	if err != nil {
		util.WriteErrorToResponse(w, err, 404)
		return
	}
	w.Header().Set("Content-Type", "audio/mpeg")
	w.WriteHeader(http.StatusOK)

	_, _ = io.Copy(w, reader)

}

// Not working

func (app *Application) DeleteTrackHandler(w http.ResponseWriter, r *http.Request) {
	req := JsonReq{}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		util.WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(bytes, &req)
	if err != nil {
		fmt.Println(err)
		util.WriteErrorToResponse(w, err, 404)
		return
	}
	fmt.Println(req)
	trackid := req.Object.Id
	trackidInt, err := strconv.Atoi(trackid)
	if err != nil {
		util.WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}
	trackidInt32 := int32(trackidInt)

	track, err := app.DB.GetTrackFromId(trackidInt32)

	if err != nil {
		util.WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}

	err = app.DB.DeleteTrack(trackidInt32)
	if err != nil {
		util.WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}

	err = gstore.DeleteObjectInBucket(app.BucketName, track.Name)
	if err != nil {
		util.WriteErrorToResponse(w, err, http.StatusInternalServerError)
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
		util.WriteErrorToResponse(w, err, 404)
		return
	}
	req := JsonReq{}
	err = json.Unmarshal(body, &req)

	if err != nil {
		fmt.Println(err)
		util.WriteErrorToResponse(w, err, 404)
		return
	}

	if req.Object.Field == "tag" && req.Object.Type == "track" && req.Action == "update" {

		trackidInt, err := strconv.Atoi(req.Object.Id)
		if err != nil {
			util.WriteErrorToResponse(w, err, 404)
			return
		}
		trackidInt32 := int32(trackidInt)
		err = app.DB.UpdateTrackTag(trackidInt32, req.Object.Body)

		if err != nil {
			util.WriteErrorToResponse(w, err, 404)
			return

		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("tag updated succefuly"))
		return

	}
	util.WriteErrorToResponse(w, fmt.Errorf("Wrong request format %s", ""), http.StatusBadRequest)
	return

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
	json.Unmarshal(bytes, &ytReq)

	mp3file, err := ytb.Download(ytReq.YtID)

	if err != nil {
		fmt.Println(err)
		util.WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}
	// load MP3 filed to bicket

	file, err := os.Open(mp3file)
	defer file.Close()
	if err != nil {
		util.WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}
	bytes, err = util.ByteFromMegaFile(file)
	if err != nil {
		util.WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}

	mp3file = util.ProcessMp3Name(mp3file)
	err = gstore.LoadToBucket(app.BucketName, mp3file, bytes)
	if err != nil {
		util.WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}
	track := repo.Track{
		Name: mp3file,
	}
	err = app.DB.PushTrackToSQL(track)

	if err != nil {
		util.WriteErrorToResponse(w, err, http.StatusInternalServerError)
		return
	}

	util.CleanAllTempDir()
	w.Write([]byte("youtube music uploaded on gcp"))

}
