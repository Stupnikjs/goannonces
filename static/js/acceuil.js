let cards = document.querySelectorAll(".audiocard")

for ( let i = 0; i < cards.length; i++ ){
    let button = cards[i].querySelector(".nameBtn")
    let tagBtn = cards[i].querySelector(".tagBtn")
    let selectedBtn = cards[i].querySelector(".selectedBtn")

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


}