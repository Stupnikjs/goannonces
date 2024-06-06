package main

import (
	"context"
	"fmt"
	"log"

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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println(err)
	}
	buck := client.Bucket("lastbucketnamethatsit")
	// err = CreateBucket(client, buck, ctx)
	defer client.Close()
	if err != nil {
		fmt.Println(err)
	}

	err = PushFileToBucket(buck)
	if err != nil {
		fmt.Println(err)
	}
}
