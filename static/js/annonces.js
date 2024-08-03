

function loadAnnonces( departement, profession, annonces) {
    let annoncesDiv = document.querySelector("#annoncesDiv")
    console.log(profession)
    let filtered = annonces.filter( el => el["profession"] == profession)

    filtered = filtered.filter(el => el["departement"] == departement) 
    console.log(filtered)

    for (let annonce of filtered){
        console.log(annonce)
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

    if (annonce["region"]) {
        lieuSpan.textContent = annonce["region"]
    }
    if (annonce["pubdate"]) {
        dateSpan.textContent = `mise en ligne le ${annonce["pubdate"]}` 
    } 
    if (annonce["contrat"]) {
        contratSpan.textContent = annonce["contrat"]
        
    }
    if (annonce["departement"]){
        depSpan.textContent = (annonce["departement"])
    }
    if (annonce["description"]){
        descriptionSpan.textContent = (annonce["description"])
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