package repo

var InitTableQuery string = `
CREATE TABLE IF NOT EXISTS tracks (
	id serial PRIMARY KEY,
	name VARCHAR,
	artist VARCHAR,
	storage_url VARCHAR,
	selected BOOLEAN, 
	size INTEGER,
	tag VARCHAR 
	); 


`

var InsertTrackQuery string = `
INSERT INTO tracks ( 
	name,
	artist,
	storage_url, 
	selected,
	size, 
	tag
) VALUES (
	$1, $2, $3, $4, $5, $6
)
`

var GetAllTrackQuery string = `
SELECT 
	id, 
	name,
	artist,
	storage_url, 
	selected, 
	size,
	tag
FROM tracks 
`

var GetTrackFromIdQuery string = `
SELECT  
	name,
	artist,
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

var UpdateTrackArtistQuery string = `
UPDATE tracks
	SET artist = $2
 	WHERE id = $1; 
`
