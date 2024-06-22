package gstore

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func CreateBucket(bucketName string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := storage.NewClient(ctx)

	if err != nil {
		fmt.Println(err)
		return err
	}
	bucket := client.Bucket(bucketName)

	bucketAttrs := &storage.BucketAttrs{
		Location: "EU",
	}

	projectID := "zikstudio"
	// Creates the new bucket.

	if err := bucket.Create(ctx, projectID, bucketAttrs); err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
		return err
	}

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

func DeleteBucket(bucketName string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := storage.NewClient(ctx)

	if err != nil {
		return err
	}
	bucket := client.Bucket(bucketName)
	err = bucket.Delete(ctx)
	if err != nil {
		return err
	}
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
			return nil, err
		}
		names = append(names, attrs.Name)
	}
	return names, nil
}

func DeleteObjectInBucket(bucketName string, objectName string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := storage.NewClient(ctx)

	_ = client.Bucket(bucketName)
	if err != nil {
		return err
	}

	return nil
}
