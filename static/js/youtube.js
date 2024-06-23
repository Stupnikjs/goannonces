let ytBtn = document.querySelector(".ytBtn")
let respDiv = document.querySelector(".respDiv")



ytBtn.addEventListener("click", async(e) => {
    e.preventDefault()
    let input = document.querySelector("input")
    let id = input.value.split("=")[1]
    let resp = await fetch(`/youtube/mp3`, {
        method: "POST", 
        headers: {
            'Content-Type': 'application/json'
        }, 
        body: JSON.stringify({
            "ytid": id
        })

    })
    let responseText = await resp.json()
            
    if (resp.ok) {
        respDiv.textContent = responseText
    } else {
        respDiv.textContent = `Error: ${responseText}`
    }
})

