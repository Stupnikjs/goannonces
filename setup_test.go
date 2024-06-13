package main

import (
	"log"
	"os"
	"testing"

	"github.com/Stupnikjs/zik/pkg/gstore"
	"github.com/joho/godotenv"
)

var TestBucket string = "mysuperstronktestbuck"

func TestMain(m *testing.M) {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	gstore.CreateBucket(TestBucket)
	// Run all the tests
	defer gstore.DeleteBucket(TestBucket)
	code := m.Run()

	// Exit with the received code
	os.Exit(code)
}
