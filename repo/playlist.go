package repo


import (

)


func (rep *PostgresRepo) CreatePlaylist(playlistName string) error {
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()

                _, err := rep.DB.ExecContext(ctx, InsertPlaylistQuery, playlistName)
                if err != nil {
                        return err
                }
        }
        return nil
}
func (rep *PostgresRepo) PushPlaylistItem(playlistName string, trackId int32) error {
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()

                _, err := rep.DB.ExecContext(ctx, InsertPlaylistTrackQuery, playlistName, trackid)
                if err != nil {
                        return err
                }
        }
        return nil
}