package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type application struct {
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := application{}

	// app.InitTables()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, app.routes())

}
