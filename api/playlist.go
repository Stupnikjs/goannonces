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

	err = json.Unmarshal(bytes, &reqJson)

	body, ok := reqJson.Object.Body.(map[string]string)
 if reqJson.Action == "create" && reqJson.Object.Type == "playlist && ok  {
   tracksIds := body["ids"] 
   for _, id := range tracksIds {
   app.DB.PushPlaylistItem(body["name"],id)
}
}

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
