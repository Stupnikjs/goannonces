package repo

var InitTableQuery string = `
CREATE TABLE IF NOT EXISTS tracks (
	id serial PRIMARY KEY,
	name VARCHAR,
	storage_url VARCHAR,
	selected BOOLEAN, 
	listen_count INTEGER,
	size INTEGER,
	tag VARCHAR 
	); 

CREATE TABLE IF NOT EXISTS playlists (
	id SERIAL PRIMARY KEY,
 	name VARCHAR
	
);

CREATE TABLE IF NOT EXISTS playlist_items (
	trackid INTEGER REFERENCES tracks(id)
 ON DELETE CASCADE,
	playlistid INTEGER REFERENCES playlists(id) 
 ON DELETE CASCADE
	
);


`

var InsertTrackQuery string = `
INSERT INTO tracks ( 
	name,
	storage_url, 
	selected,
	size, 
	tag
) VALUES (
	$1, $2, $3, $4, $5, $6,$7
)
`

var GetAllTrackQuery string = `
SELECT 
	id, 
	name,
	storage_url, 
	selected, 
	size,
	tag
FROM tracks 
`

var GetTrackFromIdQuery string = `
SELECT  
	name,
	storage_url,
	selected,   
	size, 
	tag
FROM tracks 
WHERE id = $1; 
`

var DeleteTrackQuery string = `
DELETE FROM tracks WHERE id = $1; 
`

var UpdateTrackTagQuery string = `
UPDATE tracks
	SET tag = $2
 	WHERE id = $1; 
`

// Playlist

var CreatePlaylistQuery string = `
INSERT INTO playlist (name)  
VALUES ($1) ; 
`

var GetAllPlaylistsQuery string = `
SELECT p.name  
FROM playlists p; 
`

var InsertPlaylistTrackQuery string = `
INSERT INTO playlist_items (playlistid, trackid)  
VALUES ( $1, $2 ) ; 
`

var DeletePlaylistQuery string = `
DELETE FROM playlist 
WHERE name = $1; 
`

var GetPlaylistWithTrackByPlaylistIdQuery string = `
SELECT pi.trackid 
FROM playlist_items pi
WHERE pi.playlistid = $1; 
`

var GetPlaylistWithTrackByPlaylistNameQuery string = `
SELECT pi.trackid 
FROM playlist_items pi
JOIN playlist p ON pi.playlistid = p.id 
WHERE p.name = $1; 
`
