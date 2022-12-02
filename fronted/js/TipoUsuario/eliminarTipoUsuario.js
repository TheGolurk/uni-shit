async function UpdateTypeUser() {

    let id = document.getElementById('floatingID').value;


    const url = 'http://localhost:8070/';
    const body = {

        "id":id,
    };

    console.log(body);

    try {
        const resp = await axios.post(url, body)

        showOKMessage('ok','Creado con exito');

    } catch (err) {
        switch (err.response.status) {
            // 409 Duplicado
            case 409:
                showErrMessage('Error', 'No se pudo eliminar ');
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


var form = document.getElementById("CreateTypeUser-form");
function handleForm(event) {
    event.preventDefault();
}

form.addEventListener('submit', handleForm);