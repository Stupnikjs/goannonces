package api

type Application struct {
	Port int
}

type Annonce struct {
	Url         string `json:"url"`
	PubDate     string `json:"pubdate"`
	Lieu        string `json:"lieu"`
	Departement int    `json:"departement"`
	Profession  string `json:"profession"`
	Contrat     string `json:"contrat"`
}
