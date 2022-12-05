async function AssignPrivileges() {

    let typeUser = document.getElementById('tipo').value;
    let table = document.getElementById('floatingInput').value;

    const url = 'http://localhost:8070/acces/create';
    const body = {
        "id_tipo": parseInt(typeUser),
        "tablas": table
    };

    try {
        const resp = await axios.post(url, body)
        showOKMessage('ok','Creado con exito')
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
            // 500 error del backend
            case 500:
                showErrMessage('Error', 'Ocurrió un error inesperado, intente más tarde');
                break;
            default:
                showErrMessage('Error', 'Ocurrió un error inesperado, intente más tarde');
                break;
        }
    }

}

var form = document.getElementById("UserPrivileges-form");
function handleForm(event) {
    event.preventDefault();
}


form.addEventListener('submit', handleForm);