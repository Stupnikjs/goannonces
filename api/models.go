package api

import (
	"fmt"

	"github.com/Stupnikjs/goannonces/database"
)

type Application struct {
	Port int
	DB   database.DBRepo
}

type Annonce struct {
	Id          string `json:"id"`
	Url         string `json:"url"`
	PubDate     string `json:"pubdate"`
	Ville       string `json:"ville"`
	Lieu        string `json:"lieu"`
	Departement int    `json:"departement"`
	Description string `json:"description"`
	Profession  string `json:"profession"`
	Contrat     string `json:"contrat"`
	Created_at  string `json:"created_at"`
}

func AnnonceToHtml(a Annonce) string {

	str := fmt.Sprintf(`
		<a class="annonceLink" href="%s">
	 		<span class="depSpan">%d</span>
			<span class="lieuSpan">%s</span>
		</a>

	`, a.Url, a.Departement, a.Lieu)

	return str
}
