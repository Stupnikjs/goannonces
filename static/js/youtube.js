let ytBtn = document.querySelector(".ytBtn")




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
    if (resp.ok) {
        window.location.assign("/")
    } else {
        console.log(await resp.json())
    }
   
})

