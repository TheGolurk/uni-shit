async function DeleteUser() {
    let username = document.getElementById('floatingInput').value;

    const url = `http://localhost:8080/user/delete?username=${username}`;

    try {
        const resp = await axios.delete(url)
        
        alert('Se eliminó correctamente');

    } catch (err) {
        console.log(err);
        alert('No se pudo eliminar');
    }

}

var form = document.getElementById("DeleteUser-form");
function handleForm(event) { 
    event.preventDefault(); 
}

form.addEventListener('submit', handleForm);