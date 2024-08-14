package database

import (
	"context"
	"database/sql"
	"fmt"

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

	fmt.Println(len(form))

	for k, v := range form {
		annonces = m.FilterAnnonces(k, v)
	}

}

func (m *PostgresRepo) FilterAnnonces(sqlfield string, value string) ([]api.Annonce, error) {
	annonces := []api.Annonce{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	selectFiltered := fmt.Sprintf(`
SELECT id, url, pubdate, ville, lieu, departement, description, profession, contrat, created_at
FROM annonces
WHERE %s=%s
;
`, sqlfield, value)
	rows, err := m.DB.QueryContext(ctx, selectFiltered)

	return annonces, nil
}
