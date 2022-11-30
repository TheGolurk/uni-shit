async function DeletePrivileges() {
    
    let idAccess = document.getElementById('floatingType1').value;


    const url = `http://localhost:8080/accesohora/delete?=${idAccess};`;
    const body = {
        "id": idAccess,
    };

    try {
        const resp = await axios.delete(url)
        
        alert('Se eliminó correctamente');

    } catch (err) {
        console.log(err);
        alert('No se pudo eliminar');
    }

}

var form = document.getElementById("UserPrivileges-form");
function handleForm(event) {
    event.preventDefault();
}


form.addEventListener('submit', handleForm);