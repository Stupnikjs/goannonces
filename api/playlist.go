package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Stupnikjs/zik/util"
)

func (app *Application) CreatePlaylistHandler(w http.ResponseWriter, r *http.Request) {
	reqJson := JsonReq{}
	bytes, err := io.ReadAll(r.Body)
	r.Body.Close()

	json.Unmarshal(bytes, &reqJson)

	fmt.Println(reqJson)

	if err != nil {
		util.WriteErrorToResponse(w, err, http.StatusInternalServerError)
	}

}

func (app *Application) AppendToPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	// call to app

}

func (app *Application) RemoveToPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	// call to app

}
