async function ViewSale() {
    const url = 'http://localhost:8070/venta/get';

    try {
        const resp = await axios.get(url)
        console.log(resp);

        const table = document.getElementById("tabla-datos");
        resp.data.forEach(item => {
            let row = table.insertRow();
            let idUsuario = row.insertCell(0);
            idUsuario.innerHTML = item.id_usuario_venta;

            let idPro = row.insertCell(1);
            idPro.innerHTML = item.id_pro;

            let total = row.insertCell(2);
            total.innerHTML = item.total;

            let iva = row.insertCell(3);
            iva.innerHTML = item.iva;

            let fecha = row.insertCell(4);
            fecha.innerHTML = item.fecha_venta;

            let cli = row.insertCell(5);
            cli.innerHTML = item.id_cli;
        });

    } catch (err) {
        console.log(err);
    }

}
ViewSale();
