async function DeleteSale() {
    let venta = document.getElementById('IdVenta').value;

    const url = `http://localhost:8070/venta/delete?id=${venta}`;

    try {
        const resp = await axios.delete(url)
        
        showOKMessage('ok','Se eliminó correctamente');

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


var form = document.getElementById("sale-form");
function handleForm(event) {
    event.preventDefault();
}

form.addEventListener('submit', handleForm);