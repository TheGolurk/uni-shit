async function DeleteSale() {
    let venta = document.getElementById('IdVenta').value;

    const url = `http://localhost:8070/venta/delete?id=${venta}`;

    try {
        const resp = await axios.delete(url)
        
        alert('Se eliminó correctamente');

    } catch (err) {
        console.log(err);
        alert('No se pudo eliminar');
    }

}


var form = document.getElementById("sale-form");
function handleForm(event) {
    event.preventDefault();
}

form.addEventListener('submit', handleForm);