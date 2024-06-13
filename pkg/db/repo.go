package repo

type DBrepo interface {
	InitTable()
	GetAllTracks() []Track
	PushTrackToSQL(Track)
}
