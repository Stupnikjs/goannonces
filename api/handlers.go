package api

import (
	"fmt"
	"html/template"
	"net/http"
	"path"

	"github.com/Stupnikjs/goannonces/database"
)

var pathToTemplates = "./static/templates/"

type TemplateData struct {
	Data map[string]any
}

func render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) error {
	_ = r.Method

	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t), path.Join(pathToTemplates, "/base.layout.gohtml"))
	if err != nil {
		return err
	}
	err = parsedTemplate.Execute(w, td)
	if err != nil {
		return err
	}
	return nil

}

// template rendering

func (app *Application) RenderAccueil(w http.ResponseWriter, r *http.Request) {

	td := TemplateData{}
	td.Data = make(map[string]any)
	td.Data["Dep"] = Departements
	_ = render(w, r, "/new.gohtml", &td)
}

/*



api handlers
get by city



*/

func (app *Application) GetHTMLAnnonces(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(r.Form)
	annonces, err := app.DB.SelectAnnoncesQuery(r.Form)
	fmt.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	str := ""

	fmt.Println(annonces)
	for _, a := range annonces {
		str += database.AnnonceToHtml(a)
	}

	w.Write([]byte(str))
}
