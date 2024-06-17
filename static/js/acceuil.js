let cards = document.querySelectorAll(".audiocard")

for ( let i = 0; i < cards.length; i++ ){
    let button = cards[i].querySelector(".nameBtn")
    let tagBtn = cards[i].querySelector(".tagBtn")
    let selectedBtn = cards[i].querySelector(".selectedBtn")
    let sumbitTagBtn = cards[i].querySelector(".submitTagBtn")

    // create toggle button (todo) 
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
    tagBtn.addEventListener("click", (e) => {
        e.preventDefault()
        let selected = tagBtn.getAttribute("selected")
        if (selected == "false") {
            selected = null
        }
        if (!selected){
            tagBtn.setAttribute("selected", "true")
            let hiddenTagDiv = cards[i].querySelector(".hiddenTagDiv")
            hiddenTagDiv.classList.remove("hiddenTagDiv")
            // hiddenTagDiv.style.display = "block" 
        } else {
            tagBtn.setAttribute("selected", "false")
            // trouver autre chose 
            let hiddenTagDiv = cards[i].querySelector(".hiddenTagDiv")
            hiddenTagDiv.classList.add("hiddenTagDiv")
        }
       
    })

    selectedBtn.addEventListener("click", (e) => {
        e.preventDefault()
        let selected = selectedBtn.getAttribute("selected")
        if (selected == "false") {
            selected = null
        }
        if (!selected){
            selectedBtn.setAttribute("selected", "true")
            selectedBtn.classList.add("selectedHeart")
        } else {
            selectedBtn.setAttribute("selected", "false")
            selectedBtn.classList.remove("selectedHeart")
        }
       
    })
    sumbitTagBtn.addEventListener("click", async (e) => {
        e.preventDefault()
        let input = cards[i].querySelector(".tagInput")
        let trackid = sumbitTagBtn.id
        let resp = await fetch(`/track/tag/${trackid}`, {
            method: "POST", 
            headers: {
                'Content-Type': 'application/json'
            }, 
            body: JSON.stringify({
                "tag": input.value 
            })

        })
        if (resp.ok) {
            window.location.assign("/")
        } else {
            console.log(resp.body)
        }
       
    })

}