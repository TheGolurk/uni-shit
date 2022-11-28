async function DeleteClient() {
    let name = document.getElementById('firstName').value;
    let lastName = document.getElementById('lastName').value;
    let actualAddress = document.getElementById('address').value;
    let  state = document.getElementById('state').value;

    const url = `http://localhost:8070/cliente/delete?=${name}`;


    try {
        const resp = await axios.delete(url)
        
        alert('Se eliminó correctamente');

    } catch (err) {
        console.log(err);
        alert('No se pudo eliminar');
    }

}

var form = document.getElementById("client-form");
function handleForm(event) {
    event.preventDefault();
}

form.addEventListener('submit', handleForm);