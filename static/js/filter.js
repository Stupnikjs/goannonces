let annoncesRow = document.querySelector("#json").getAttribute("data-json")
let annonces = JSON.parse(annoncesRow)
console.log(annonces)



/* choice div */
let choiceDiv = document.querySelector("#choiceDiv")
let locationDiv = document.querySelector("#locationDiv")



let pharmacienBtn = document.createElement("button") 
pharmacienBtn.textContent = "Tu es Pharmacien"

let preparateurBtn = document.createElement("button") 
preparateurBtn.textContent = "Tu es Preparateur"

let rayonisteBtn = document.createElement("button") 
rayonisteBtn.textContent = "Tu es Rayoniste"




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






function processProfessionBtn(){
    choiceDiv.classList.add("d-none")
    locationDiv.classList.remove("d-none")
    locationDiv.classList.add("showlocationDiv")
    let select = ReturnSelectDep(profession, annonces)
    locationDiv.append(select)
}