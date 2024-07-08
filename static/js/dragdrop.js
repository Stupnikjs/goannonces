
let files = []
let submitdrag = document.querySelector("#submit_drag")
let msgDiv = document.querySelector("#msg")

// #drop_zone ondrop=dropHandler(ev)
function dropHandler(ev) {
    console.log("File(s) dropped");
    // Prevent default behavior (Prevent file from being opened)
    ev.preventDefault();
  
    if (ev.dataTransfer.items) {
      
        [...ev.dataTransfer.items].forEach((item, i) => {
        // If dropped items aren't files, reject them
          if (item.kind === "file") {
          const file = item.getAsFile();
          console.log(Object.keys(item));
          files.push(file)
          displayFileName(file.name)
        }
      });
    } else {
      
      [...ev.dataTransfer.files].forEach((file, i) => {
        console.log(`… file[${i}].name = ${file.name}`);
        console.log(Object.keys(file))  
        displayFileName(file.name)
        files.push(file)  
      });
    }
  }

  // #drag_zone ondrag=dragOverHandler(ev)
  function dragOverHandler(ev) {
    console.log("File(s) in drop zone");
    
    // Prevent default behavior (Prevent file from being opened)
    ev.preventDefault();
  }

 async function submitHandler() {

  console.log("submit handler")
  const formData = new FormData();

  files.forEach((file, index) => {
      formData.append('files[]', file, file.name);
  });

  let resp = await fetch('/upload', {
      method: 'POST',
      body: formData
  })
  
  // do something while waiting resp 
  // loader animation 
  msgDiv.textContent = "uploading files"

  if ( resp.ok ) {
      console.log('Success:', resp.body);
      // window.location.assign("/")
  } else {
      console.error('Error:', error);
  }
}


submitdrag.addEventListener("click", (e) => {
e.preventDefault()
await submitHandler()

})

function displayFileName(name) {
    let dropZone = document.querySelector("#drop_zone")
    let div = document.createElement("div");
    div.classList.add("file_item")
    div.innerHTML = `<i class="fa-solid fa-music"></i>`
    div.textContent = name;
    dropZone.appendChild(div);
}
