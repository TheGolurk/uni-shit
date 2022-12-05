async function ModifySale() {
    let id  = document.getElementById('IdVenta').value;
    let idUser = document.getElementById('Idusuario').value;
    let idProduct = document.getElementById('IdProducto').value;
    let total = document.getElementById('Total').value;
    let iva = document.getElementById('Iva').value;
    let date = document.getElementById('fechaVenta').value;
    let idClient = document.getElementById('IdCliente').value;

    const url = 'http://localhost:8070/venta/modify';
    const body = {
        "id_venta": parseInt(id),
        "id_usuario_venta": parseInt(idUser),
        "id_pro": parseInt(idProduct),
        "total": parseFloat(total),
        "iva": parseFloat(iva),
        "fecha_venta": date,
        "id_cli": parseInt(idClient),
    };

    try {
        const resp = await axios.put(url, body)
        showOKMessage('ok','Actualizado con exito');

    } catch (err) {

        switch (err.response.status) {
            // 409 Duplicado
            case 409:
                showErrMessage('Error', 'No se pudo Modificar ');
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