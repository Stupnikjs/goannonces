package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Stupnikjs/goannonces/api"
	"github.com/Stupnikjs/goannonces/database"
	"github.com/joho/godotenv"
	_ "google.golang.org/api/option"
)

func main() {

	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	app := api.Application{
		Port: 8080,
	}

	conn, err := api.ConnectToDB()

	app.DB = &database.PostgresRepo{
		DB: conn,
	}

	if err != nil {
		fmt.Println(err)
	}

	http.ListenAndServe(fmt.Sprintf(":%d", app.Port), app.Routes())

}
