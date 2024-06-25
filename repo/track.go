package repo

import (
	"context"
	"database/sql"
	"log"
)

type Track struct {
	ID       int
	Name     string
	StoreURL string
	Selected bool
	UserId   int
	Size     int32
	Tag      string
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
		track.Size,
		track.Tag,
	)
	if err != nil {
		return err
	}
	return nil
}

func (rep *PostgresRepo) GetTrackFromId(id int) (*Track, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	row := rep.DB.QueryRowContext(ctx, GetTrackFromIdQuery, id)
	track := &Track{}

	track.ID = id
	row.Scan(
		&track.Name,
		&track.StoreURL,
		&track.Selected,
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

func (rep *PostgresRepo) DeleteTrack(trackId int) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := rep.DB.ExecContext(ctx, DeleteTrackQuery, trackId)
	if err != nil {
		return err
	}
	return nil
}

func (rep *PostgresRepo) UpdateTrackTag(trackId int, tag string) error {
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
