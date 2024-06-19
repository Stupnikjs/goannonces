package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/Stupnikjs/zik/pkg/gstore"
	"github.com/go-chi/chi/v5"
)

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

// reqModel

func (app *Application) DeleteTrackHandler(w http.ResponseWriter, r *http.Request) {

	req := JsonReq{}
	bytes, err := io.ReadAll(r.Body)
	err = json.Unmarshal(bytes, &req)
	if err != nil {
		fmt.Println(err)
		WriteErrorToResponse(w, err, 404)
		return
	}
	trackid := req.Object.Id
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("track deleted succefully"))
}

// reqModel
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
		trackidInt32 := int32(trackidInt)
		err = app.DB.UpdateTrackTag(trackidInt32, req.Object.Body)

		if err != nil {
			WriteErrorToResponse(w, err, 404)
			return

		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("tag updated succefuly"))
		return

	}
	WriteErrorToResponse(w, fmt.Errorf("Wrong request format %s", ""), http.StatusBadRequest)
	return

}
