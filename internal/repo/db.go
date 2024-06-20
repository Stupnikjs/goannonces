package repo

import (
	"context"
	"database/sql"
	"log"
	"strconv"
)

type Track struct {
	ID             int32
	Name           string
	StoreURL       string
	Selected       bool
	SelectionCount int
	PlayCount      int
	Size           int32
	Tag            string
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
		track.Selected,
		track.SelectionCount,
		track.PlayCount,
		track.Size,
		track.Tag,
	)
	if err != nil {
		return err
	}
	return nil
}

func (rep *PostgresRepo) GetTrackFromId(id string) (*Track, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	row := rep.DB.QueryRowContext(ctx, GetTrackFromIdQuery, id)
	track := &Track{}
	intID, err := strconv.Atoi(id)

	if err != nil {
		return nil, err

	}
	track.ID = int32(intID)
	row.Scan(
		&track.Name,
		&track.StoreURL,
		&track.Selected,
		&track.PlayCount,
		&track.SelectionCount,
		&track.Size,
		&track.Tag,
	)
	return track, nil
}

func (rep *PostgresRepo) InitTable() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := rep.DB.ExecContext(ctx, InitTableQuery)
	if err != nil {
		log.Fatalf("error initing table %v", err)

	}
}

func (rep *PostgresRepo) GetAllTracks() ([]Track, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rows, err := rep.DB.QueryContext(ctx, GetAllTrackQuery)
	if err != nil {
		return nil, err
	}
	tracks := []Track{}
	track := Track{}
	for rows.Next() {
		err := rows.Scan(
			&track.ID,
			&track.Name,
			&track.StoreURL,
			&track.Selected,
			&track.PlayCount,
			&track.SelectionCount,
			&track.Size,
			&track.Tag,
		)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, track)

	}
	return tracks, nil
}

func (rep *PostgresRepo) DeleteTrack(trackId int32) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := rep.DB.ExecContext(ctx, DeleteTrackQuery, trackId)
	if err != nil {
		return err
	}
	return nil
}

func (rep *PostgresRepo) UpdateTrackTag(trackId int32, tag string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := rep.DB.ExecContext(ctx, UpdateTrackTagQuery, trackId, tag)
	if err != nil {
		return err
	}
	return nil
}

// get most played Track with num arg
// create route to increment PlayCount
// selectcnt
