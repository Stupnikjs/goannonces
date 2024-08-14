package database

import (
	"context"
	"database/sql"
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


selectFiltered := `
SELECT id, url, pubdate, ville, lieu, departement, description, profession, contrat, created_at
FROM annonces
WHERE 
;
`
