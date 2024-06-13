package repo

type Dbrepo interface {
	InitTable()
	GetAllTracks() []Track
	PushTrackToSQL(Track)
}
