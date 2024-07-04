package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Stupnikjs/zik/repo"
)

type JsonReq struct {
	Action string    `json:"action"`
	Object ApiObject `json:"object"`
}

type ApiObject struct {
	Type  string `json:"type"`
	Id    string `json:"id,omitempty"`
	Body  any    `json:"body,omitempty"`
	Field string `json:"field,omitempty"`
}

type Application struct {
	DB         repo.DBrepo
	Port       int
	BucketName string
}

func ParseJsonReq(r *http.Request) (*JsonReq, error) {

	reqJson := JsonReq{}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body.Close()

	err = json.Unmarshal(bytes, &reqJson)

	if err != nil {
		return nil, err
	}
	return &reqJson, nil

}
