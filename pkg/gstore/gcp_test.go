package gstore

import (
	"context"
	"fmt"
	"math/rand/v2"
	"testing"

	"cloud.google.com/go/storage"
)

var TestBucket string = "mysuperstronktestbuck"

func TestCreateBucket(t *testing.T) {
	randint := rand.IntN(2000)

	randBuckName := fmt.Sprintf("randombuck%dname", randint)

	err := CreateBucket(randBuckName)

	if err != nil {
		t.Errorf("Unexpected error creating bucket %s", err.Error())

	}
	defer DeleteBucket(randBuckName)
}

func TestGetBucketObject(t *testing.T) {

	data := []byte("this is test files content")

	err := LoadToBucket(TestBucket, "test.txt", data)

	// Call get bucket method
	if err != nil {
		t.Errorf(" expected no error but go %s", err)
	}

}

func TestDeleteBucket(t *testing.T) {
	CreateBucket(TestBucket)
	err := DeleteBucket(TestBucket)

	if err != nil {

		t.Errorf("error deleting bucket %s", err)

	}
}
func TestDeleteBucketObject(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, _ := storage.NewClient(ctx)
	bucket := client.Bucket(TestBucket)
	obj := bucket.Object("test.txt")
	writer := obj.NewWriter(ctx)
	writer.Write([]byte("this is test content"))
	writer.Close()

	err := DeleteObjectInBucket(TestBucket, "test.txt")
	if err != nil {
		t.Errorf("error deleting object %v", err)
	}

}
