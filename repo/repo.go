package repo

type DBrepo interface {
	InitTable()
	GetAllTracks() ([]Track, error)
	PushTrackToSQL(Track) error
	GetTrackFromId(int32) (*Track, error)
	DeleteTrack(int32) error
	UpdateTrackTag(int32, string) error
}
