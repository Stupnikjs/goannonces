package repo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

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

type PostgresRepo struct {
	DB *sql.DB
}

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

func (rep *PostgresRepo) connectToDB() (*sql.DB, error) {

	connection, err := openDB()

	if err != nil {
		return nil, err
	}
	log.Println("Connected to Postgres!")
	return connection, nil
}

func (rep *PostgresRepo) PushTrackToSQL(track Track) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := rep.DB.ExecContext(
		ctx,
		InsertTrackQuery,
		track.Name,
		track.StoreURL,
		track.SelectionCount,
		track.PlayCount,
		track.Size,
	)
	if err != nil {
		return err
	}
	return nil
}

func (rep *PostgresRepo) InitTable() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := rep.DB.ExecContext(ctx, InitTableQuery)
	if err != nil {
		log.Fatalf("error initing table %v", err)

	}
}

func (rep *PostgresRepo) GetAllTracks() []Track {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rows, err := rep.DB.QueryContext(ctx, GetAllTrackQuery)
	if err != nil {
		fmt.Println(err)
	}
	tracks := []Track{}
	track := Track{}
	for rows.Next() {
		err := rows.Scan(
			&track.ID,
			&track.Name,
			&track.StoreURL,
			&track.PlayCount,
			&track.SelectionCount,
			&track.Size,
		)
		if err != nil {
			fmt.Println(err)
			break
		}
		tracks = append(tracks, track)

	}
	return tracks
}
