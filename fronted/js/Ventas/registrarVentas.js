async function RegisterSale() {
    let idUser = document.getElementById('Idusuario').value;
    let idProduct = document.getElementById('IdProducto').value;
    let total = document.getElementById('Total').value;
    let iva = document.getElementById('Iva').value;
    let date = document.getElementById('fechaVenta').value;
    let idClient = document.getElementById('IdCliente').value;

    const url = 'http://localhost:8070/venta/create';
    const body = {
        "IdUusarioVenta": idUser,
        "IdPro": idProduct,
        "Total": total,
        "Iva": iva,
        "FechaVenta": date,
        "IdCli": idClient,
    };

    try {
        const resp = await axios.post(url, body)

        alert('Creado con exito');

    } catch (err) {
        console.log(err);
        alert('No se pudo crear, valida la informacion');
    }

}


var form = document.getElementById("sale-form");
function handleForm(event) {
    event.preventDefault();
}

form.addEventListener('submit', handleForm);