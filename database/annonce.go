package database

import (
	"context"
	"database/sql"

	"github.com/Stupnikjs/goannonces/api"
)

type PostgresRepo struct {
	DB *sql.DB
}

func (m *PostgresRepo) CheckId(id string) bool {
	var res int
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	selectStmt := `
	SELECT 1
		FROM annonces
		WHERE id = $1
		LIMIT 1;
	`
	row := m.DB.QueryRowContext(ctx, selectStmt, id)

	err := row.Scan(&res)
	if res == 1 && err != nil {
		return true
	}
	return false
}

func (m *PostgresRepo) SelectAnnoncesQuery(form map[string][]string) ([]api.Annonce, error) {
	annonces := []api.Annonce{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	job := MapContains(form, "job")
	dep := MapContains(form, "dep")
	contract := MapContains(form, "contract")

	selectFiltered := `
SELECT id, url, pubdate, ville, lieu, departement, description, profession, contrat, created_at
FROM annonces
WHERE 
;
`

	rows, err := m.DB.QueryContext(ctx, selectFiltered)

	return annonces, nil
}
