package repo

import "context"

type Playlist struct {
	Name   string
	Tracks []int
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

func (rep *PostgresRepo) GetAllPlaylists() ([]Playlist, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rows, err := rep.DB.QueryContext(ctx, GetAllPlaylistsQuery)
	if err != nil {
		return nil, err
	}
	playlists := []Playlist{}

	for rows.Next() {
		playlist := Playlist{}
		rows.Scan(
			&playlist.Name,
		)
		playlists = append(playlists, playlist)
	}

	return playlists, nil
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
