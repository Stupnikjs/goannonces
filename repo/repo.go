package repo

type DBrepo interface {
	InitTable()
	GetAllTracks() ([]Track, error)
	PushTrackToSQL(Track) error
	GetTrackFromId(int) (*Track, error)
	DeleteTrack(int) error
	UpdateTrackTag(int, string) error

	// Playlist
	InsertPlaylistTrack(string, int)
}
