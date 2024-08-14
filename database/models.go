package database

import (
	"fmt"
)

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
			<span class="contratSpan">%s</span>
			<span class="dateSpan">%s</span>
			<span class="professionSpan">%s</span>
		</a>

	`, a.Url, a.Departement, a.Lieu, a.Contrat, a.PubDate, a.Profession)

	return str
}

func (a Annonce) IsInArray(arr []Annonce) bool {

	for _, an := range arr {
		if a.Id == an.Id {
			return true
		}
	}
	return false

}

func IntersectionAnnonces(ann1 []Annonce, ann2 []Annonce) []Annonce {
	retArr := []Annonce{}
	if len(ann1) >= len(ann2) {
		for _, ann := range ann1 {
			if ann.IsInArray(ann2) {
				retArr = append(retArr, ann)
			}
		}
	} else {
		for _, an := range ann2 {
			if an.IsInArray(ann1) {
				retArr = append(retArr, an)
			}
		}

	}
	return retArr

}
