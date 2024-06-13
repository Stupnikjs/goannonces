package repo

type Dbrepo interface {
	connectToDB()
	InitTable()
	GetAllTracks() []Track
	PushTrackToSQL(Track)
}
