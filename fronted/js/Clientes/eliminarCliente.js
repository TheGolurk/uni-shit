async function DeleteClient() {
    let id = document.getElementById('Idventa').value;

    const url = `http://localhost:8070/cliente/delete?id=${id}`;


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