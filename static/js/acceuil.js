let cards = document.querySelectorAll(".audiocard")


// loping over each sound card 
for ( let i = 0; i < cards.length; i++ ){
    let button = cards[i].querySelector(".nameBtn")
    let tagBtn = cards[i].querySelector(".tagBtn")
    let selectedBtn = cards[i].querySelector(".selectedBtn")
    let sumbitTagBtn = cards[i].querySelector(".submitTagBtn")
    let deleteBtn = cards[i].querySelector(".deleteBtn")

    // Button to display the audio player  
    button.addEventListener("click", (e) => {
        e.preventDefault()
        let selected = button.getAttribute("selected")
        if (!selected){
            button.setAttribute("selected", "true")
            let audio = document.createElement("audio")
            audio.src = button.getAttribute("data-url") 
            audio.controls = true
            audio.preload = "auto"
            cards[i].appendChild(audio)
        }
       
    })

    
    // button to display tag input 
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

    

    // heart selected 
    selectedBtn.addEventListener("click", (e) => {
        e.preventDefault()
        let selected = selectedBtn.getAttribute("selected")
        
        if (!selected || selected == "false"){
            selectedBtn.setAttribute("selected", "true")
            selectedBtn.classList.add("selectedHeart")
        } else {
            selectedBtn.setAttribute("selected", "false")
            selectedBtn.classList.remove("selectedHeart")
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
    sumbitTagBtn.addEventListener("click", async (e) => {
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
