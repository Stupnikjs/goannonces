package database

type DBRepo interface {
	CheckId(id string) bool
	SelectAnnoncesQuery(map[string][]string) ([]Annonce, error)
	FilterAnnonces(string, string) ([]Annonce, error)
}
