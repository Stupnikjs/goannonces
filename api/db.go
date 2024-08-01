package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// fly postgres connect -a pharmadb
func openDB() (*sql.DB, error) {
	uri := os.Getenv("DATABASE_URL")

	print("hell", uri)
	db, err := sql.Open(
		"postgres",
		uri,
	)
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return db, nil
}

func (app *Application) ConnectToDB() (*sql.DB, error) {

	connection, err := openDB()

	if err != nil {
		return nil, err
	}
	log.Println("Connected to Postgres!")
	return connection, nil
}
