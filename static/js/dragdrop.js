
let files = []
let submitdrag = document.querySelector("#submit_drag")
let msgDiv = document.querySelector("#msg")


let resp = await fetch("/api/allobjects", {
method: "GET"
})

let trackObjects = await resp.json()
console.log(trackObjects)



// #drop_zone ondrop=dropHandler(ev)
function dropHandler(ev) {
    
    let alreadyInBucket = []
    // Prevent default behavior (Prevent file from being opened)
    ev.preventDefault();
  
    if (ev.dataTransfer.items) {
      
        [...ev.dataTransfer.items].forEach((item, i) => {
        // If dropped items aren't files, reject them
          if (item.kind === "file")  {
          const file = item.getAsFile();
          if (trackObjects.includes(file.name)){
            console.log(`${file.name} already in bucket`)
            alreadyInBucket.push(file.name)
          } else {
            files.push(file)
            displayFileName(file.name)
          }
          
        }
      });
    } else {
      
      [...ev.dataTransfer.files].forEach((file, i) => {
        console.log(`… file[${i}].name = ${file.name}`);
        if (trackObjects.includes(file.name)){
          alreadyInBucket.push(file.name)
        } else {
          displayFileName(file.name)
          files.push(file)  
        }
        
      });
    }
    let dropZone = document.querySelector("#drop_zone")
    let alreadyInBucketUl = document.createElement("ul")
    for (let i; i < alreadyInBucket.length; i++ ){
      let li = document.createElement("li")
      li.textContent = `file ${alreadyInBucket[i]} were already in bucket`
      alreadyInBucketUl.appendChild(li)
    }
  }

  // #drag_zone ondrag=dragOverHandler(ev)
  function dragOverHandler(ev) {
    console.log("File(s) in drop zone");
    // Prevent default behavior (Prevent file from being opened)
    ev.preventDefault();
    // work with chrome 
  }
  
  
  // 
 function removeFromDropZone(file){
  
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
      msgDiv.textContent = "success uploading"
      console.log('Success:', resp.body);
      // window.location.assign("/")
  } else {
      console.error('Error:', error);
      msgDiv.textContent = error 
  }
}


submitdrag.addEventListener("click", async (e) => {
  e.preventDefault()
  await submitHandler()

})

function displayFileName(name) {
    let dropZone = document.querySelector("#drop_zone")
    let div = document.createElement("div");
 
    div.classList.add("file_item")
    div.textContent = name;
  
    dropZone.appendChild(div);
}
