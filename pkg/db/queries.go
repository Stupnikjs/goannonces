package repo

var InitTableQuery string = `
CREATE TABLE IF NOT EXISTS tracks (
	id serial PRIMARY KEY,
	name VARCHAR,
	storage_url VARCHAR,
	selected_count INTEGER, 
	listen_count INTEGER, 
	size INTEGER
	)
`

var InsertTrackQuery string = `
INSERT INTO tracks ( 
	name,
	storage_url, 
	selected_count, 
	listen_count, 
	size
) VALUES (
	$1, $2, $3, $4, $5
)
`

var GetAllTrackQuery string = `
SELECT 
	id, 
	name,
	storage_url, 
	selected_count, 
	listen_count, 
	size
FROM tracks 
`

var GetTrackFromIdQuery string = `
SELECT  
	name,
	storage_url, 
	selected_count, 
	listen_count, 
	size
FROM tracks 
WHERE id = $1; 
`
