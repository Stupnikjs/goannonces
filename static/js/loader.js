const inputSubmit = document.querySelector("#submitFileNum")

let fileNumber = 1

inputSubmit.addEventListener("change", (e) => {
    fileNumber = document.querySelector("selectFileNum").value
    
    let form = document.querySelector("form")

    for (let i=0; i < fileNumber; i++){
        let fileInput = document.createElement("input")
        fileInput.type = "file"
        fileInput.name = "file"+parseInt(i)
        form.appendChild(fileInput)
    }

})

