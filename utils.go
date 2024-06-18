package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Stupnikjs/zik/pkg/gstore"
	"github.com/Stupnikjs/zik/pkg/repo"
)

type ErrorResp struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func IsInSlice[T comparable](str T, arr []T) bool {
	for _, s := range arr {
		if str == s {
			return true
		}

	}
	return false
}

func ByteFromMegaFile(file io.Reader) ([]byte, error) {

	reader := bufio.NewReader(file)

	finalByteArr := make([]byte, 0, 2048*1000)

	for {
		soloByte, err := reader.ReadByte()
		if err != nil {
			log.Println(err)
			break
		}

		finalByteArr = append(finalByteArr, soloByte)
	}

	return finalByteArr, nil

}

func WriteErrorToResponse(w http.ResponseWriter, err error, status int) {
	errResp := ErrorResp{}
	errResp.Error = err.Error()
	bytes, jsonErr := json.Marshal(errResp)
	if jsonErr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bytes)
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
