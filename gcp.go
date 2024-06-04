package main

import (
	_ "cloud.google.com/go/storage"
	"google.golang.org/api/option"
	_ "google.golang.org/api/option"
)

func (app *application) ConnectToGcp() {
	// read json credentials file
	creds := make([]byte, 1000)
	options := option.WithCredentialsJSON(creds)
	print(options)
}
