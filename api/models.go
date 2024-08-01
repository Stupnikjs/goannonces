package api

import "database/sql"

type Application struct {
	Port int
	DB   *sql.DB
}

type Annonce struct {
	Url         string `json:"url"`
	PubDate     string `json:"pubdate"`
	Lieu        string `json:"lieu"`
	Region      string `json:"region"`
	Departement int    `json:"departement"`
	Description string `json:"description"`
	Profession  string `json:"profession"`
	Contrat     string `json:"contrat"`
}
