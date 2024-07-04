let createPlaylistBtn = document.querySelector(".createPlaylistBtn")
let dataDiv = document.querySelector("#trackDiv")
let tracks = dataDiv.getAttribute("data-json")
let errDiv = document.querySelector("#errDiv")

tracks = JSON.parse(tracks)



createPlaylistBtn.addEventListener("click", (e) => {

    e.preventDefault()
    createPlaylistBtn.style.display = "none"
    let inputName = document.createElement("input")
    let labelName = document.createElement("label")
    labelName.style.display = "block"
    labelName.textContent = "Name your Playlist"
    let submitBtn = document.createElement("button")
    submitBtn.textContent = "Submit Playlist"

    submitBtn.addEventListener("click", async(e) => {
        
        let apiObject = {
            "action": "create",
            "object": {
               "type" : "playlist",
                "body": {
                    "name": inputName.value,
                }
            }
        }
        let resp = await fetch("/api/playlist/create", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            }, 
            body: JSON.stringify(apiObject)
        })
        if (resp.ok) {
            
            errDiv.textContent = await resp.text()
        }


    })
    dataDiv.appendChild(labelName)
    dataDiv.appendChild(inputName)
    dataDiv.appendChild(submitBtn)

})


function loadTrackChoice(){
    
    let trackChoiceUl = document.createElement("ul")
    
    for (let i= 0; i < tracks.length; i++){

        let trackChoiceLi = document.createElement("li")

        trackChoiceLi.textContent = tracks[i]["Name"]
        trackChoiceLi.id = tracks[i]["ID"]
        let buttonSelect = document.createElement("button")
        buttonSelect.innerHTML = '<i class="fa-solid fa-plus"></i>'
        buttonSelect.addEventListener('click', (e) => {
            if (trackChoiceLi.ariaSelected == "false" || !trackChoiceLi.ariaSelected) {
            trackChoiceLi.ariaSelected = "true"
            buttonSelect.style.backgroundColor = "black"
            buttonSelect.style.color = "white"
            } else {
                trackChoiceLi.ariaSelected = "false"
                buttonSelect.style.backgroundColor = ""
                buttonSelect.style.color = ""

            }
        })
        trackChoiceLi.appendChild(buttonSelect)
        trackChoiceUl.appendChild(trackChoiceLi)
    }
    dataDiv.appendChild(trackChoiceUl)
}