async function ModifyClient() {
    let id = document.getElementById('idClient').value;
    let name = document.getElementById('firstName').value;
    let lastName = document.getElementById('lastName').value;
    let actualAddress = document.getElementById('address').value;
    let  state = document.getElementById('state').value;

    let cookie = getCookie('user-login');
    let userinfo = JSON.parse(cookie);

    const url = `http://localhost:8070/cliente/modify?idtipo=${parseInt(userinfo.id_tipo)}`;
    const body = {
        "id_cliente": parseInt(id),
        "nombre": name,
        "apellido": lastName,
        "direccion": actualAddress,
        "estado": state,
    };

    try {
        const resp = await axios.put(url, body)
        
        showOKMessage('ok','Actualizado con exito');

    } catch (err) {
        switch (err.response.status) {
            // 409 Duplicado
            case 409:
                showErrMessage('Error', 'Duplicado ');
                break;
            // 400 el request esta mal
            case 400:
                showErrMessage('Error', 'Error en la información, compruebe de nuevo');
                break;
            case 401:
                    showErrMessage('Error', 'No tienes permisos sobre esta tabla');
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


var form = document.getElementById("client-form");
function handleForm(event) {
    event.preventDefault();
}

form.addEventListener('submit', handleForm);