package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"
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
	annonces, err := LoadAnnonces()
	if err != nil {
		fmt.Println(err)
	}
	bytes, err := json.Marshal(annonces)

	if err != nil {
		fmt.Println(err)
	}
	td := TemplateData{}
	td.Data = make(map[string]any)
	td.Data["Annonces"] = string(bytes)
	_ = render(w, r, "/acceuil.gohtml", &td)
}


/*



api handlers 
get by city



*/



func (app *Application) GetAnnonces(w http.ResponseWriter, r *http.Request){
  

  
   

}