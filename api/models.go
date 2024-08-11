package api

import (
	"database/sql"
	"fmt"
	"strconv"
)

type Application struct {
	Port int
	DB   *sql.DB
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

func MapContains(m map[string][]string, key string) bool {
	v, exists := m[key]

	// handle case with int == 0
	if Contains(v, "") {
		return false
	} else if !exists {
		return false
	} else {
		return true
	}

}

func GetFilteredAnnonces(form map[string][]string) []Annonce {

	filteredAnnonces := []Annonce{}
	annonces, _ := LoadAnnonces()

	// Check if required keys are present in the form
	job := MapContains(form, "job")
	dep := MapContains(form, "dep")
	contract := MapContains(form, "contract")

	fmt.Println(job, dep, contract)
	for _, an := range annonces {
		// Assume that all criteria are required to match if they are provided
		matches := true

		// Apply filters based on presence in the form
		if job {
			if an.Profession != form["job"][0] {
				matches = false
			}
		}
		if dep {
			// map over
			fmt.Println(form["dep"][0])
			depNum, err := strconv.Atoi(form["dep"][0])
			if err != nil || an.Departement != depNum {
				matches = false
			}
		}
		if contract {
			if an.Contrat != form["contract"][0] {
				matches = false
			}
		}

		// If all criteria are satisfied, append the annonce
		if matches {
			filteredAnnonces = append(filteredAnnonces, an)
		}
	}

	return filteredAnnonces
}
