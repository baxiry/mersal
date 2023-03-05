const url = "http://localhost:8080"

// putData put form data to server
function putData(id) {
    let elmInput = $("#"+id);

    var data = new FormData();
    data.append(id, elmInput.value);
    fetch(url+"/upacount", { method: "POST", body: data })
    // .then(e) => console.log("response is : ", e.data); 
}

// createInput creates new input element
function creatInput(id) {
    console.log("id is : ", id)

    let elmInput = $("#"+id);
    elmInput.innerHTML = inputForm(id);//'<input type="text" placeholder='+id+' id='+id+' autofocus>'
    //elmInput.innerHTML += '<button class="btn btn-outline-primary"  name=update >update</button>'
}

// $ my awsome javascript framework
function $(element) {
    return document.querySelector(element)
}


function inputForm(id) {
    if (id == "photos") {
        return `
    <form class="row g-3">
        <div class="col-auto">
            <input class="form-control form-control-sm" id="formFileSm" type="file">
        </div>
        <div class="col-auto">
            <button type="" class="btn btn-primary mb-3">update</button>
        </div>
   </form>
    `
    }
    return `
    <div class="row g-3">
        <div class="col-auto">
            <input type="text" class="form-control" id="${id}" placeholder="${id}">
        </div>
        <div class="col-auto">
            <button type="" class="btn btn-primary mb-3" oncklic="console.log('hello frinds')">update</button>
        </div>
    </div>
    `
/*    return `
 *putData(${id})
 <form>
  <input type="text" id="${id}" name="${id}" placeholder="${id}" autofocus >
  <input type="submit" name="update" >
</form>
`*/
}
