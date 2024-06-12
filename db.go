package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
)

type Track struct {
	ID             int32
	Name           string
	StoreURL       string
	SelectionCount int
	PlayCount      int
	Size           int32
	UploadDate     string
}

var InitTableQuery string = `
CREATE TABLE IF NOT EXISTS tracks (
	id serial PRIMARY KEY,
	name VARCHAR,
	storage_url VARCHAR,
	selected_count INTEGER, 
	listen_count INTEGER, 
	upload_date DATE,
	size INTEGER
	)
`

var InsertTrackQuery string = `
INSERT INTO tracks ( 
	name,
	storage_url, 
	selected_count, 
	listen_count, 
	upload_date, 
	size
) VALUES (
	$1, $2, $3, $4, $5, $6, $7
)
`

func openDB() (*sql.DB, error) {

	cleanup, err := pgxv4.RegisterDriver("cloudsql-postgres", cloudsqlconn.WithCredentialsFile("credentials.json"))
	if err != nil {
		fmt.Println(err)
	}
	// call cleanup when you're done with the database connection
	defer cleanup()

	if err != nil {
		fmt.Println(err)

	}

	// Call cleanup when you're done with the database connection

	db, err := sql.Open(
		"cloudsql-postgres",
		fmt.Sprintf("host=%s user=postgres password=%s dbname=postgres sslmode=disable", os.Getenv("SQL_HOST"), os.Getenv("SQL_PASSWORD")))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return db, nil
}

func (app *Application) connectToDB() (*sql.DB, error) {

	connection, err := openDB()

	if err != nil {
		return nil, err
	}
	log.Println("Connected to Postgres!")
	return connection, nil
}

func (app *Application) PushTrackToSQL(track Track) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := app.DB.ExecContext(
		ctx,
		InsertTrackQuery,
		track.Name,
		track.StoreURL,
		track.SelectionCount,
		track.PlayCount,
		time.Now().Format("11-11-2023"),
		track.Size,
	)
	if err != nil {
		return err
	}
	return nil
}

func (app *Application) InitTable() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := app.DB.ExecContext(ctx, InitTableQuery)
	if err != nil {
		log.Fatalf("error initing table %v", err)

	}
}
