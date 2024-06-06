package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/storage"
)

func CreateBucket(client *storage.Client, bucket *storage.BucketHandle, ctx context.Context) error {

	projectID := "celestial-tract-421819"
	// Creates the new bucket.
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	if err := bucket.Create(ctx, projectID, nil); err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
		return err
	}

	fmt.Print("Bucket created.\n")
	return nil

}

func PushToBucket(bucket *storage.BucketHandle, filename string, data []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	obj := bucket.Object(filename)
	w := obj.NewWriter(ctx)
	_, err := w.Write(data)
	defer w.Close()

	if err != nil {
		return err
	}
	return nil
}
