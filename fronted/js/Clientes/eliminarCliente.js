async function DeleteClient() {
    let name = document.getElementById('firstName').value;
    let lastName = document.getElementById('lastName').value;
    let actualAddress = document.getElementById('address').value;
    let  state = document.getElementById('state').value;

    const url = 'http://localhost:8070/cliente/delete';
    const body = {
        "Nombre": name,
        "Apellido": lastName,
        "Direccion": actualAddress,
        "Estado": state,
    };

    try {
        const resp = await axios.delete(url, body)
        
        alert('Creado con exito');

    } catch (err) {
        console.log(err);
        alert('No se pudo crear, valida la informacion');
    }

}

var form = document.getElementById("client-form");
function handleForm(event) {
    event.preventDefault();
}

form.addEventListener('submit', handleForm);