package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Stupnikjs/goannonces/api"
	_ "google.golang.org/api/option"
)

func main() {

	app := api.Application{
		Port: 8080,
	}

	db, err := app.ConnectToDB()

	app.DB = db

	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(fmt.Sprintf(":%d", app.Port), app.Routes())

}
