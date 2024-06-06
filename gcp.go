package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/storage"

	"google.golang.org/api/option"
)

func (app *application) ConnectToGcp() *storage.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()
	options := option.WithCredentialsFile("credentials.json")
	client, err := storage.NewClient(ctx, options)

	if err != nil {
		fmt.Println(err)
	}

	return client

}

func TestBucket(bucketName string, client storage.Client) bool {
	bucket := client.Bucket(bucketName)
	if bucket != nil {
		return true
	}
	return false
}

func CreateBucket(bucketName string, client storage.Client, ctx context.Context) error {
	bucket := client.Bucket(bucketName)
	projectID := "celestial-tract-421819"
	// Creates the new bucket.
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	if err := bucket.Create(ctx, projectID, nil); err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
		return err
	}

	fmt.Printf("Bucket %v created.\n", bucketName)
	return nil

}
