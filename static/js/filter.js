
/* choice div */
let choiceDiv = document.querySelector("#choiceDiv")
let locationDiv = document.querySelector("#locationDiv")

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
    processProfessionBtn()
})
preparateurBtn.addEventListener("click", (e) => {
    e.preventDefault()
    profession = "Préparateur"
    processProfessionBtn()
})
rayonisteBtn.addEventListener("click", (e) => {
    e.preventDefault()
    profession = "Rayoniste"
    processProfessionBtn()

})




submitDepBtn.addEventListener("click", (e) => {
    e.preventDefault()
    let input = document.querySelector("input[type='number']")
    departement = parseInt(input.value)
    loadAnnonces(profession, departement, annonces)
})





function processProfessionBtn(){
    let span = document.createElement("span")
    span.textContent = profession
    choiceDiv.removeChild(pharmacienBtn)
    choiceDiv.removeChild(preparateurBtn)
    choiceDiv.removeChild(rayonisteBtn)
    choiceDiv.appendChild(span)
    locationDiv.classList.remove("d-none")
    locationDiv.classList.add("showlocationDiv")
}