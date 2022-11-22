async function ViewUser() {
    let username = document.getElementById('floatingInput').value;
    let password = document.getElementById('floatingPassword').value;
    let name = document.getElementById('floatingName').value;
    let lastName = document.getElementById('floatingLastName').value;
    let typeUser = document.getElementById('floatingType').value;

    const url = 'http://localhost:8080/user';
    const body = {
        "username": username,
        "password": password,
        "nombre": name,
        "apellido": lastName,
        "id_tipo": typeUser
    };

    try {
        const resp = await axios.post(url, body)
        console.log(resp);
    } catch (err) {
        console.log(err);
    }

}

const items1 = [
    { username: "chris", nombre: "asdads", apellido: "asdad", id_tipo: 1 },
  ];

  
    const table = document.getElementById("tabla-datos");
    items1.forEach( item => {
      let row = table.insertRow();
      let date = row.insertCell(0);
      date.innerHTML = item.username;
      let name = row.insertCell(1);
      name.innerHTML = item.nombre;
      let apellido = row.insertCell(2);
      apellido.innerHTML = item.apellido;
      let tipo = row.insertCell(3);
      tipo.innerHTML = item.id_tipo;
    });
  


var form = document.getElementById("ViewUser-form");
function handleForm(event) { 
    event.preventDefault(); 
}

form.addEventListener('submit', handleForm);