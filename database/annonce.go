package database

import (
	"context"
	"database/sql"
	"fmt"
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

func (m *PostgresRepo) SelectAnnoncesQuery(form map[string][]string) ([]Annonce, error) {
	var err error
	annonces := []Annonce{}
	arr := FormToArray(form)
	var count int

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
		if arr[i][1] == "" || arr[i][1] == "0" {
			fmt.Println("null val")
			continue
		}

		if count == 0 {
			annonces, err = m.FilterAnnonces(arr[i][0], arr[i][1])
			if err != nil {
				return nil, err
			}
		}

		newannonces, err := m.FilterAnnonces(arr[i][0], arr[i][1])
		if err != nil {
			return nil, err
		}
		annonces = IntersectionAnnonces(annonces, newannonces)
		count += 1
	}
	return annonces, nil

}

func (m *PostgresRepo) FilterAnnonces(sqlfield string, value string) ([]Annonce, error) {
	annonces := []Annonce{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	selectFiltered := fmt.Sprintf(`
SELECT id, url, pubdate, ville, lieu, departement, description, profession, contrat, created_at
FROM annonces
WHERE %s=$1
;
`, sqlfield)
	fmt.Println(selectFiltered)
	rows, err := m.DB.QueryContext(ctx, selectFiltered, value)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var annonce Annonce
		// Scan each row into the annonce struct
		err := rows.Scan(
			&annonce.Id,
			&annonce.Url,
			&annonce.PubDate,
			&annonce.Ville,
			&annonce.Lieu,
			&annonce.Departement,
			&annonce.Description,
			&annonce.Profession,
			&annonce.Contrat,
			&annonce.Created_at,
		)
		if err != nil {
			return nil, err
		}
		annonces = append(annonces, annonce)
	}
	return annonces, nil
}
