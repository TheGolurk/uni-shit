async function ViewSale() {
    const url = 'http://localhost:8070/venta/get';

    try {
        const resp = await axios.get(url)
        console.log(resp);

        const table = document.getElementById("tabla-datos");
        resp.data.forEach(item => {
            let row = table.insertRow();

            let idVenta = row.insertCell(0);
            idventa.innerHTML = item.id_venta;

            let idUsuario = row.insertCell(1);
            idUsuario.innerHTML = item.id_usuario_venta;

            let idPro = row.insertCell(2);
            idPro.innerHTML = item.id_pro;

            let total = row.insertCell(3);
            total.innerHTML = item.total;

            let iva = row.insertCell(4);
            iva.innerHTML = item.iva;

            let fecha = row.insertCell(5);
            fecha.innerHTML = item.fecha_venta;

            let cli = row.insertCell(6);
            cli.innerHTML = item.id_cli;
        });

    } catch (err) {
        console.log(err);
    }

}
ViewSale();
