package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "cloud.google.com/go/storage"
	"github.com/joho/godotenv"
	_ "google.golang.org/api/option"
)

type Application struct {
	DB   *sql.DB
	Port int
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := Application{
		Port: 3322,
	}

	http.ListenAndServe(fmt.Sprintf(":%d", app.Port), app.routes())

}
