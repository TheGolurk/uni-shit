async function UpdateUser() {
    let username = document.getElementById('floatingInput').value;
    let name = document.getElementById('floatingName').value;
    let lastName = document.getElementById('floatingLastName').value;
    let typeUser = document.getElementById('floatingType').value;

    
    const url = 'http://localhost:8070/user/modify';
    const body = {
        "username": username,
        "nombre": name,
        "apellido": lastName,
        "id_tipo": parseInt(typeUser)
    };
    
    try {
        const resp = await axios.put(url, body)

        alert('Se actualizo con exito');

    } catch (err) {
        console.log(err);
        alert('No se pudo actualizar');
    }

}


var form = document.getElementById("UpdateUser-form");
function handleForm(event) { 
    event.preventDefault(); 
}

form.addEventListener('submit', handleForm);