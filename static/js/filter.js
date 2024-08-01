
/* choice div */
let choiceDiv = document.querySelector("#choiceDiv")
let locationDiv = document.querySelector("#locationDiv")
let annoncesRow = document.querySelector("#json").getAttribute("data-json")

let annonces = JSON.parse(annoncesRow)

console.log(annonces)

let submitDepBtn = document.querySelector("#submitDepBtn")


let pharmacienBtn = document.createElement("button") 
pharmacienBtn.textContent = "Je suis Pharmacien"

let preparateurBtn = document.createElement("button") 
preparateurBtn.textContent = "Je suis Preparateur"

let rayonisteBtn = document.createElement("button") 
rayonisteBtn.textContent = "Je suis Rayoniste"




let profession = ""
let departement = 0


choiceDiv.appendChild(pharmacienBtn)
choiceDiv.appendChild(preparateurBtn)
choiceDiv.appendChild(rayonisteBtn)


pharmacienBtn.addEventListener("click", (e) => {
    e.preventDefault()
    profession = "Pharmacien"
    choiceDiv.style.display = "none" 
    locationDiv.classList.remove("d-none")
    locationDiv.classList.add("showlocationDiv")
})
preparateurBtn.addEventListener("click", (e) => {
    e.preventDefault()
    profession = "Préparateur"
    choiceDiv.style.display = "none" 
    locationDiv.classList.remove("d-none")
    locationDiv.classList.add("showlocationDiv")
})
rayonisteBtn.addEventListener("click", (e) => {
    e.preventDefault()
    profession = "Rayoniste"
    choiceDiv.style.display = "none" 
    locationDiv.classList.remove("d-none")
    locationDiv.classList.add("showlocationDiv")

})




submitDepBtn.addEventListener("click", (e) => {
    e.preventDefault()
    let input = document.querySelector("input[type='number']")
    departement = parseInt(input.value)
    loadAnnonces(profession, departement, annonces)
})



function loadAnnonces(profession, departement, annonces) {
    let annoncesDiv = document.querySelector("#annoncesDiv")


    for (let annonce of annonces){
        console.log(annonces["profession"])
        let el = createAnnoncesElement(annonce)
        annoncesDiv.appendChild(el) 
    }
    
}



function createAnnoncesElement(annonce){
    
    let div = document.createElement("div")
    div.classList.add("annonceDiv")
    let span = document.createElement("span")
    if (annonce["profession"]) {
        span.textContent = annonce["profession"]
    }
    div.appendChild(span)
    return div
}