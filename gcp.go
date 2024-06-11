package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

var BucketName string = "firstappbucknamezikapp"

func CreateBucket(bucketName string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := storage.NewClient(ctx)

	if err != nil {
		fmt.Println(err)
		return err
	}
	bucket := client.Bucket(bucketName)

	projectID := "zikstudio"
	// Creates the new bucket.

	// CHOSE EUROPE WEST

	if err := bucket.Create(ctx, projectID, nil); err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
		return err
	}

	fmt.Print("Bucket created.\n")
	return nil

}

func LoadToBucket(bucketName string, fileName string, data []byte) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := storage.NewClient(ctx)

	if err != nil {
		fmt.Println(err)
		return err
	}
	bucket := client.Bucket(bucketName)

	defer client.Close()

	obj := bucket.Object(fileName)

	writer := obj.NewWriter(ctx)
	writer.Write(data)
	defer writer.Close()

	// get object url to store in sql
	// get object id

	return nil
}

func GetObjectURL(bucketName string, objectName string) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := storage.NewClient(ctx)

	if err != nil {
		return "", err
	}
	bucket := client.Bucket(bucketName)
	obj := bucket.Object(objectName)

	attr, err := obj.Attrs(ctx)

	return attr.MediaLink, nil
}

func ListObjectsBucket(bucketName string) ([]string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := storage.NewClient(ctx)

	bucket := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	query := &storage.Query{Prefix: ""}

	var names []string
	it := bucket.Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		names = append(names, attrs.Name)
	}
	return names, nil
}

func (app *Application) LoadMultipartReqToBucket(r *http.Request, bucketName string) error {
	objNames, err := ListObjectsBucket(BucketName)

	if err != nil {
		return err
	}

	for _, headers := range r.MultipartForm.File {

		for _, h := range headers {
			file, err := h.Open()

			if err != nil {
				return err
			}

			defer file.Close()

			finalByteArr, err := ByteFromMegaFile(file)

			if err != nil {
				return err
			}

			err = LoadToBucket(bucketName, h.Filename, finalByteArr)

			if err != nil {
				return err
			}

		}

	}
	return nil

}

func ByteFromMegaFile(file io.Reader) ([]byte, error) {

	reader := bufio.NewReader(file)

	finalByteArr := make([]byte, 0, 2048*1000)

	for {
		soloByte, err := reader.ReadByte()
		if err != nil {
			log.Println(err)
			break
		}

		finalByteArr = append(finalByteArr, soloByte)
	}

	return finalByteArr, nil

}
