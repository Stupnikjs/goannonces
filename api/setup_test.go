package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/joho/godotenv"
)

var app Application
var TestBucketName = "TestHandlersBucket"

func TestMain(m *testing.M) {

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := Application{}
	app.Port = 2222
	setup()
	code := m.Run()
	cleanup()
	// Exit with the received code
	os.Exit(code)
}

func setup() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatal(err)
	}

	bucket := client.Bucket(TestBucketName)

	bucketAttrs := &storage.BucketAttrs{
		Location: "EU",
	}

	projectID := "zikstudio"
	// Creates the new bucket.

	if err := bucket.Create(ctx, projectID, bucketAttrs); err != nil {
		fmt.Println("here")
		log.Fatalf("Failed to create bucket: %v", err)
	}

	obj := bucket.Object("test.mp3")
	writer := obj.NewWriter(ctx)

	testFilePath := `C:\Users\nboud\OneDrive\Bureau\Go_Projects\zik\static\download\test.mp3`

	file, _ := os.Open(testFilePath)
	b, _ := io.ReadAll(file)

	writer.Write(b)

}

func cleanup() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatal(err)
	}

	bucket := client.Bucket(TestBucketName)

	err = bucket.Delete(ctx)
	fmt.Println("here", err)

}
