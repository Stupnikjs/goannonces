let cards = document.querySelectorAll(".audiocard"); 

console.log(cards)
let filterDiv = document.querySelector("#filterDiv"); 
let inputFilter = document.querySelector("#inputFilter");  
let paddingGrids = document.querySelectorAll(".padding-grid")
let selectedTrack = []



    /* 
    *  
    *  Filter 
    * 
    */

inputFilter.addEventListener("input", (e) => {
    filter = e.target.value
    
    for ( let i = 0; i < cards.length; i++ ){
        let tag = cards[i].querySelector(".tagp") 
        let name = cards[i].querySelector(".name")
       if (!name.textContent.includes(filter)){ 
           cards[i].style.display = "none"  
        } else {
            cards[i].style.display = "block" 
        }
    }
}) 




// loping over each sound card 
for ( let i = 0; i < cards.length; i++ ){
    let button = cards[i].querySelector(".nameBtn")
    let tagBtn = cards[i].querySelector(".tagBtn")
    let selectedBtn = cards[i].querySelector(".selectedBtn")
    let sumbitTagBtn = cards[i].querySelector(".submitTagBtn")
    let deleteBtn = cards[i].querySelector(".deleteBtn")

    /*  
    *  
    *  Audio Player
    * 
    */

    button.addEventListener("click", (e) => {
        e.preventDefault()
        let selected = button.getAttribute("selected")
        if (!selected || button.getAttribute("selected") == "false"){
            console.log(button.getAttribute("data-url") )
            button.setAttribute("selected", "true")
            let audio = cards[i].querySelector("audio") ? cards[i].querySelector("audio") : document.createElement("audio");  
            audio.src = button.getAttribute("data-url") 
            audio.controls = true
            audio.preload = "auto"
            audio.style.display = "block"
            cards[i].appendChild(audio)
        } else {
            button.setAttribute("selected", "false")
            let audio = cards[i].querySelector("audio")
            audio.style.display = "none"
        }
       
    })

    
    /*  
    *  
    *  Tag toggle 
    * 
    */ 

    tagBtn.addEventListener("click", (e) => {
        e.preventDefault()
        let selected = tagBtn.getAttribute("selected")
        if (selected == "false") {
            selected = null
        }
        if (!selected){
            tagBtn.setAttribute("selected", "true")
            let tagDiv = cards[i].querySelector(".tagDiv")
            tagDiv.classList.remove("display-none")

        } else {
            tagBtn.setAttribute("selected", "false")
            // trouver autre chose 
            let tagDiv = cards[i].querySelector(".tagDiv")
            tagDiv.classList.add("display-none")
        }
       
    })

    

    /*  
    *  
    *  Selected 
    * 
    */

    selectedBtn.addEventListener("click", (e) => {

        e.preventDefault()

        let selected = selectedBtn.getAttribute("selected")
        let gridSelected = paddingGrids[1]
        let name = cards[i].querySelector(".name")

        if (!selected || selected == "false"){
           
            selectedBtn.setAttribute("selected", "true")
            selectedBtn.classList.add("selectedHeart")

            
            if (!selectedTrack.includes(name.textContent)){
                let nameGrid = document.createElement("p")
                nameGrid.textContent = name.textContent
                let id = simpleHash(name.textContent) 
                nameGrid.id = id 
                gridSelected.appendChild(nameGrid)
                selectedTrack.push(name.textContent)
            }
           
            
        } else {
            selectedBtn.setAttribute("selected", "false")
            selectedBtn.classList.remove("selectedHeart")
            let id = simpleHash(name.textContent)
            let toremove = document.getElementById(id)
            // selectedTrack = selectedTrack.filer( x => x != )
            // remove trackName 
            toremove.remove()
        }
       
    })


    
    
    // delete the track then refresh 
    deleteBtn.addEventListener("click", async (e) => {
        e.preventDefault()
        
        let trackid = deleteBtn.id
        let resp = await fetch(`/api/track/remove`, {
            method: "POST", 
            headers: {
                'Content-Type': 'application/json'
            }, 
            body: JSON.stringify({
                "action": "delete",
                "object": {
                   "type" : "track" ,
                    "id": trackid
                }
            })

        })
        if (resp.ok) {
            window.location.assign("/")
        } else {
            console.log(resp.json())
        }
       
    })



    // post the tag then refresh 
    sumbitTagBtn.addEventListener("click", async(e) => {
        e.preventDefault()
        let input = cards[i].querySelector(".tagInput")
        let trackid = sumbitTagBtn.id
        let resp = await fetch(`/api/track/tag`, {
            method: "POST", 
            headers: {
                'Content-Type': 'application/json'
            }, 
            body: JSON.stringify({
                                "action": "update",
                "object": {
                    "type" : "track",
                    "id": trackid,
                    "field" : "tag",
                    "body": input.value
                }
            })

        })
        if (resp.ok) {
            window.location.assign("/")
        } else {
            console.log(resp.json())
        }
       
    })

}

function simpleHash(input){
    let code = 0
    for ( let i=0; i < input.length-1; i++){
        code += input[i].charCodeAt(0)
    }
    let hex = btoa(input)
    return code.toString() + hex.slice(0, 5) 

}


