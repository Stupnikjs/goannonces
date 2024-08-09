const departements = {
    "Ain":                     1,
	"Aisne":                   2,
	"Allier":                  3,
	"Alpes-de-Haute-Provence": 4,
	"Hautes-Alpes":            5,
	"Alpes-Maritimes":         6,
	"Ardèche":                 7,
	"Ardennes":                8,
	"Ariège":                  9,
	"Aube":                    10,
	"Aude":                    11,
	"Aveyron":                 12,
	"Bouches-du-Rhône":        13,
	"Calvados":                14,
	"Cantal":                  15,
	"Charente":                16,
	"Charente-Maritime":       17,
	"Cher":                    18,
	"Corrèze":                 19,
	"Corse-du-Sud":            2, // Special case for Corsica
	"Haute-Corse":             2, // Special case for Corsica
	"Côte-d'Or":               21,
	"Côtes-d'Armor":           22,
	"Creuse":                  23,
	"Doubs":                   25,
	"Drôme":                   26,
	"Eure":                    27,
	"Eure-et-Loir":            28,
	"Finistère":               29,
	"Gard":                    30,
	"Haute-Garonne":           31,
	"Gers":                    32,
	"Gironde":                 33,
	"Hérault":                 34,
	"Ille-et-Vilaine":         35,
	"Indre":                   36,
	"Indre-et-Loire":          37,
	"Isère":                   38,
	"Jura":                    39,
	"Landes":                  40,
	"Loir-et-Cher":            41,
	"Loire":                   42,
	"Haute-Loire":             43,
	"Loire-Atlantique":        44,
	"Loiret":                  45,
	"Lot":                     46,
	"Lot-et-Garonne":          47,
	"Lozère":                  48,
	"Maine-et-Loire":          49,
	"Manche":                  50,
	"Marne":                   51,
	"Haute-Marne":             52,
	"Mayenne":                 53,
	"Meurthe-et-Moselle":      54,
	"Meuse":                   55,
	"Morbihan":                56,
	"Moselle":                 57,
	"Nièvre":                  58,
	"Nord":                    59,
	"Oise":                    60,
	"Orne":                    61,
	"Pas-de-Calais":           62,
	"Puy-de-Dôme":             63,
	"Pyrénées-Atlantiques":    64,
	"Hautes-Pyrénées":         65,
	"Pyrénées-Orientales":     66,
	"Bas-Rhin":                67,
	"Haut-Rhin":               68,
	"Rhône":                   69,
	"Haute-Saône":             70,
	"Saône-et-Loire":          71,
	"Sarthe":                  72,
	"Savoie":                  73,
	"Haute-Savoie":            74,
	"Paris":                   75,
	"Seine-Maritime":          76,
	"Seine-et-Marne":          77,
	"Yvelines":                78,
	"Deux-Sèvres":             79,
	"Somme":                   80,
	"Tarn":                    81,
	"Tarn-et-Garonne":         82,
	"Var":                     83,
	"Vaucluse":                84,
	"Vendée":                  85,
	"Vienne":                  86,
	"Haute-Vienne":            87,
	"Vosges":                  88,
	"Yonne":                   89,
	"Territoire de Belfort":   90,
	"Essonne":                 91,
	"Hauts-de-Seine":          92,
	"Seine-Saint-Denis":       93,
	"Val-de-Marne":            94,
	"Val-d'Oise":              95,
	"Guadeloupe":              971,
	"Martinique":              972,
	"Guyane":                  973,
	"La Réunion":              974,
	"Mayotte":                 976,
}



function ReturnSelectDep(profession, annonces) {

   let select = document.createElement("select")
   let locationDiv = document.querySelector("#locationDiv")
   select.classList.add("selectDep")
   for ( let dep of Object.entries(departements)) {
    let option = document.createElement("option")
    option.textContent = `${dep[1]}  ${dep[0]}` 
    select.appendChild(option)
   }
   let deps = []
   select.addEventListener('change', (e) => {
    e.preventDefault()
    let dep = e.target.value.split(" ") 
	if (!deps.includes(parseInt(dep))){
		deps.push(dep[0])
		
		let depSelected = document.createElement("div")
		let span = document.createElement("span")
		span.innerHTML = `
		<i class="fa-solid fa-x></i>
		`
		depSelected.appendChild(span)
		
		depSelected.textContent = dep[1]
		depSelected.addEventListener("click", (e) => {
			console.log("remove")
			deps = deps.filter( x => x != dep[0])
		})
		locationDiv.appendChild(depSelected)
		resetAnnonces()
		loadAnnonces(deps, profession, annonces)
	}

	// ajouter un submit btn 
    
    
   })
   return select
   



}