let cards = document.querySelectorAll(".audiocard")

for ( let i = 0; i < cards.length; i++ ){
    let button = cards[i].querySelector("button")
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
            console.log(button.getAttribute("data-url"))
            cards[i].appendChild(audio)
        }
       
    })
}