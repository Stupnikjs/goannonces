package main

import (
	_ "cloud.google.com/go/storage"
	"fmt"
	"github.com/Stupnikjs/zik/internal/api"
	"github.com/Stupnikjs/zik/internal/repo"
	"github.com/Stupnikjs/zik/pkg/util"
	"github.com/joho/godotenv"
	_ "google.golang.org/api/option"
	"log"
	"net/http"
	"os"
)

var BucketName string = "firstappbucknamezikapp"

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := api.Application{
		Port: 3322,
	}
	app.BucketName = BucketName

	conn, err := app.ConnectToDB()

	app.DB = &repo.PostgresRepo{
		DB: conn,
	}
	if err != nil {
		log.Fatalf("Error loading db conn: %v", err)
	}

	app.DB.InitTable()

	http.ListenAndServe(fmt.Sprintf(":%d", app.Port), app.Routes())

	curr, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}
	defer util.CleanAllTempDir(curr)

}
