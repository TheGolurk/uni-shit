async function DeletePrivileges() {
    
    let typeUser = document.getElementById('floatingType').value;
    let table = document.getElementById('floatingInput').value;

    const url = `http://localhost:8080/accesohora/delete?=${f};`;
    const body = {
        "id_tipo": typeUser,
        "tabla": table
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