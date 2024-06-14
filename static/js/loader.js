const inputSubmit = document.querySelector("#submitFileNum")

let fileNumber = 1
let form = document.querySelector("form")



inputSubmit.addEventListener("click", (e) => {
    fileNumber = document.querySelector("#selectFileNum").value
    

    for (let i=0; i < fileNumber; i++){
        let fileInput = document.createElement("input")
        fileInput.type = "file"
        fileInput.name = "file"+parseInt(i)
        form.appendChild(fileInput)
    }

})

