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

func (app *application) LoadToBucket(fileName string, []byte data) error {


	
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		client, err := storage.NewClient(ctx)


		buck := client.Bucket(app.BucketName)

		// Check if bucket already created
		err = CreateBucket(client, buck, ctx)

		if err != nil {
			return err 
		}
		defer client.Close()
		if err != nil {
			return err 
		}


			obj := buck.Object(fileName)

			writer := obj.NewWriter(ctx)
   writer.Write(buf)
			defer writer.Close()
			if err != nil {
				fmt.Println(err)
			}


return nil 
} 

