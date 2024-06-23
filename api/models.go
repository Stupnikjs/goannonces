package api

import "github.com/Stupnikjs/zik/repo"

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
