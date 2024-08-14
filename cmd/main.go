package main

import (
	"fmt"
	"net/http"

	"github.com/Stupnikjs/goannonces/api"
	"github.com/Stupnikjs/goannonces/database"
	_ "google.golang.org/api/option"
)

func main() {

	app := api.Application{
		Port: 8080,
	}

	conn, err := app.ConnectToDB()

	app.DB = database.PostgresRepo{
		DB: conn,
	}

	if err != nil {
		fmt.Println(err)
	}

	http.ListenAndServe(fmt.Sprintf(":%d", app.Port), app.Routes())

}
