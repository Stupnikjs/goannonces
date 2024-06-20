package api

import "github.com/Stupnikjs/zik/internal/repo"

type JsonReq struct {
	Action string    `json:"action"`
	Object ApiObject `json:"object"`
}

type ApiObject struct {
	Type  string `json:"type"`
	Id    string `json:"id"`
	Body  string `json:"body,omitempty"`
	Field string `json:"field,omitempty"`
}

type Application struct {
	DB         repo.DBrepo
	Port       int
	BucketName string
}