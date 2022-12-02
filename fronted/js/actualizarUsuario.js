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

        showOKMessage('Ok', 'Se actualizo con exito')

    } catch (err) {
        switch (err.response.status) {
            // 409 Duplicado
            case 409:
                showErrMessage('Error', 'No se pudo actualizar ');
                break;
            // 400 el request esta mal
            case 400:
                showErrMessage('Error', 'Error en la información, compruebe de nuevo');
                break;
            // 500 error del backend
            case 500:
                showErrMessage('Error', 'Ocurrió un error inesperado, intente más tarde');
                break;
            default:
                showErrMessage('Error', 'Ocurrió un error inesperado, intente más tarde');
                break;
        }

        console.log(err);
    }

}


var form = document.getElementById("UpdateUser-form");
function handleForm(event) { 
    event.preventDefault(); 
}

form.addEventListener('submit', handleForm);