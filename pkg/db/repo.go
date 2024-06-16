package repo

type DBrepo interface {
	InitTable()
	GetAllTracks() ([]Track, error)
	PushTrackToSQL(Track) error
	GetTrackFromId(string) (*Track, error)
	DeleteTrack(int32) error
}
