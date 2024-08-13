package database

import (
	"database/sql"

	"github.com/Stupnikjs/goannonces/api"
)

type PostgresRepo struct {
	DB *sql.DB
}

func (m *PostgresRepo) GetAnnonce(id string) api.Annonce {

	return nil
}
