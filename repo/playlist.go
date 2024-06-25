package repo

import "context"

type Playlist struct {
	Name      string
	ID        int32
	TracksIDs []int32
}

func (rep *PostgresRepo) CreatePlaylist(playlistName string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := rep.DB.ExecContext(ctx, CreatePlaylistQuery, playlistName)
	if err != nil {
		return err
	}

	return nil
}

func (rep *PostgresRepo) DeletePlaylist(playlistName string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := rep.DB.ExecContext(ctx, DeletePlaylistQuery, playlistName)
	if err != nil {
		return err
	}

	return nil
}

func (rep *PostgresRepo) InsertPlaylistTrack(playlistName string, trackId int) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := rep.DB.ExecContext(ctx, InsertPlaylistTrackQuery, playlistName, trackId)
	if err != nil {
		return err
	}
	return nil
}
