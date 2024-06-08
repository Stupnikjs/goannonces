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

type application struct {
	BucketName string
	DB         *sql.DB
	Port       int
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := application{
		BucketName: "firstappbuckname",
		Port:       3322,
	}

	http.ListenAndServe(fmt.Sprintf(":%d", app.Port), app.routes())

}
