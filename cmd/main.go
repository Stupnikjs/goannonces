package main

import (
	"fmt"
	"net/http"

	"github.com/Stupnikjs/goannonces/api"
	_ "google.golang.org/api/option"
)

func main() {

	app := api.Application{
		Port: 8080,
		DB: 
	}

	/*
		db, err := app.ConnectToDB()

		app.DB = db

		if err != nil {
			fmt.Println(err)
		}
	*/

	http.ListenAndServe(fmt.Sprintf(":%d", app.Port), app.Routes())

}
