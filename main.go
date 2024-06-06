package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/storage"
	_ "cloud.google.com/go/storage"
	"github.com/joho/godotenv"
	_ "google.golang.org/api/option"
)

type application struct {
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// app.InitTables()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println(err)
	}
	b := client.Bucket("michel")
	attr, err := b.Attrs(ctx)

	err = b.Create(ctx, "celestial-tract-421819", attr)
	fmt.Println(os.Getenv("GOOGLE_APPPLICATION_CREDENTIALS"))
	if err != nil {
		print("here error")
		fmt.Println(err)
	}

}
