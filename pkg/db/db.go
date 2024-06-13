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
