package repo

var InitTableQuery string = `
CREATE TABLE IF NOT EXISTS tracks (
	id serial PRIMARY KEY,
	name VARCHAR,
	storage_url VARCHAR,
	selected BOOLEAN,
	selected_count INTEGER, 
	listen_count INTEGER, 
	size INTEGER,
	tag VARCHAR 
	)
`

var InsertTrackQuery string = `
INSERT INTO tracks ( 
	name,
	storage_url, 
	selected,
	selected_count, 
	listen_count, 
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
	selected_count, 
	listen_count, 
	size,
	tag
FROM tracks 
`

var GetTrackFromIdQuery string = `
SELECT  
	name,
	storage_url,
	selected,  
	selected_count, 
	listen_count, 
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
