package util

import (
	"encoding/json"
	"net/http"
)

type ErrorResp struct {
	Error   string `json:"error"`
	Message string `json:"message"`
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
