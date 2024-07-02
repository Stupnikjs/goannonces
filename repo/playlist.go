package repo

import "context"



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

func (rep *PostgresRepo) InsertPlaylistTrack(playlistId int, trackId int) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := rep.DB.ExecContext(ctx, InsertPlaylistTrackQuery, playlistId, trackId)
	if err != nil {
		return err
	}
	return nil
}
