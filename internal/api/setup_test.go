package api

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var app Application

func TestMain(m *testing.M) {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := Application{}
	app.Port = 2222
	code := m.Run()

	// Exit with the received code
	os.Exit(code)
}
