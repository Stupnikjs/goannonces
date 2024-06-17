
let files = []

let submitdrag = document.querySelector("#submit_drag")

function dropHandler(ev) {
    console.log("File(s) dropped");
    let dragZone = document.querySelector("#drag_zone")
    // Prevent default behavior (Prevent file from being opened)
    ev.preventDefault();
  
    if (ev.dataTransfer.items) {
      // Use DataTransferItemList interface to access the file(s)
      [...ev.dataTransfer.items].forEach((item, i) => {
        // If dropped items aren't files, reject them
        if (item.kind === "file") {
          const file = item.getAsFile();
          console.log(`… file[${i}].name = ${file.name}`);
          files.push(file)
          
        }
      });
    } else {
      // Use DataTransfer interface to access the file(s)
      [...ev.dataTransfer.files].forEach((file, i) => {
        console.log(`… file[${i}].name = ${file.name}`);
        let p = document.createElement("p")
        p.textContent = file.name
        dragZone.appendChild(p)
        files.push(file)  
      });
    }
  }

  function dragOverHandler(ev) {
    console.log("File(s) in drop zone");
  
    // Prevent default behavior (Prevent file from being opened)
    ev.preventDefault();
  }

  function submitHandler() {

    console.log("submit handler")
  const formData = new FormData();

  files.forEach((file, index) => {
      formData.append('files[]', file, file.name);
  });

  fetch('/upload', {
      method: 'POST',
      body: formData
  })
  .then(response => response.json())
  .then(data => {
      console.log('Success:', data);
      window.location.assign("/dragdrop")
  })
  .catch(error => {
      console.error('Error:', error);
  });
}


submitdrag.addEventListener("click", (e) => {
e.preventDefault()

submitHandler()

})