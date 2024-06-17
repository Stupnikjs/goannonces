package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net/http"
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
