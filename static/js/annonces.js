let annoncesRow = document.querySelector("#json").getAttribute("data-json")

let annonces = JSON.parse(annoncesRow)
console.log(annonces)


function loadAnnonces(profession, departement, annonces) {
    let annoncesDiv = document.querySelector("#annoncesDiv")

    let filtered = annonces.filter( el => el["profession"] == profession)

    filtered = filtered.filter(el => el["departement"] == departement) 

    for (let annonce of filtered){
        
        let el = createAnnoncesElement(annonce)
        annoncesDiv.appendChild(el) 
    }
    
}



function createAnnoncesElement(annonce){
    let a = document.createElement("a")
    a.classList.add("annonceDiv")
    let lieuSpan = document.createElement("span")
    lieuSpan.classList.add("lieuSpan")
    let dateSpan = document.createElement("span")
    dateSpan.classList.add("dateSpan")
    let depSpan = document.createElement("span")
    depSpan.classList.add("depSpan")
    
    let contratSpan = document.createElement("span")
    contratSpan.classList.add("jobSpan")

    if (annonce["lieu"]) {

        lieuSpan.textContent = annonce["lieu"]
        
    }
    if (annonce["pubdate"]) {

        dateSpan.textContent = annonce["pubdate"]
        
    }
    
    if (annonce["contrat"]) {
        contratSpan.textContent = annonce["contrat"]
        
    }

    if (annonce["departement"]){
        depSpan.textContent = (annonce["departement"])
    }
    if (annonce["url"]){
        a.href = annonce["url"]
    }



    a.appendChild(lieuSpan)
    a.appendChild(dateSpan)
    a.appendChild(depSpan)
    a.appendChild(contratSpan)
    return a
}