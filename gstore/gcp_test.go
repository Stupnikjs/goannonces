package gstore

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/Stupnikjs/zik/util"
)

var TestDeleteBucketName = "somedeletebuckettestname"

func TestCreateBucket(t *testing.T) {
	randint := rand.IntN(2000)

	randBuckName := fmt.Sprintf("randombuck%dname", randint)

	err := CreateBucket(randBuckName)

	if err != nil {
		t.Errorf("Unexpected error creating bucket %s", err.Error())

	}

	err = LoadToBucket(randBuckName, "rand.txt", []byte("random text"))

	if err != nil {
		t.Errorf("Unexpected error loading file to newbucket %s", err.Error())

	}

	defer DeleteBucket(randBuckName)
}

func TestGetBucketObject(t *testing.T) {

	data := []byte("this is test files content")

	err := LoadToBucket(TestBucketName, "test.txt", data)

	// Call get bucket method
	if err != nil {
		t.Errorf(" error in getter object %s", err)
	}

}

func TestDeleteBucket(t *testing.T) {
	fmt.Println("there")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := storage.NewClient(ctx)

	if err != nil {
		t.Error("Error creating client in test delete bucket")
	}

	bucket := client.Bucket(TestDeleteBucketName)

	bucketAttrs := &storage.BucketAttrs{
		Location: "EU",
	}

	projectID := "zikstudio"
	// Creates the new bucket.

	if err := bucket.Create(ctx, projectID, bucketAttrs); err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	}

	client.Close()

	err = DeleteBucket(TestDeleteBucketName)

	if err != nil {
		t.Errorf("error deleting bucket %s", err)
	}

	_, err = ListObjectsBucket(TestDeleteBucketName)

	if err == nil {
		t.Errorf("error should trigger since bucket is deleted instead got %s", err)
	}

}
func TestDeleteBucketObject(t *testing.T) {
	fmt.Println("here")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, _ := storage.NewClient(ctx)
	bucket := client.Bucket(TestBucketName)
	obj := bucket.Object("test.txt")
	writer := obj.NewWriter(ctx)
	writer.Write([]byte("this is test content"))
	writer.Close()
	client.Close()

	err := DeleteObjectInBucket(TestBucketName, "test.txt")
	if err != nil {
		t.Errorf("error deleting object %v", err)
	}

	list, err := ListObjectsBucket(TestBucketName)

	if err == nil {
		t.Error("expected error since object got deleted")
	}
	if util.IsInSlice[string]("test.txt", list) {
		t.Error("test.txt should have been removed")
	}

	if err == nil {
		t.Error("Should return error since object is suppoed to be deleted")
	}

}
