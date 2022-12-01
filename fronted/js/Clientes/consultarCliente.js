async function GetClient() {
    const url = 'http://localhost:8070/cliente/get';

    try {
        const resp = await axios.get(url)
        console.log(resp);

        const table = document.getElementById("tabla-datos");
        resp.data.forEach(item => {

            let row = table.insertRow();

            let id = row.insertCell(0);
            id.innerHTML = item.id_cliente;

            let nombre = row.insertCell(1);
            nombre.innerHTML = item.nombre;

            let apellido = row.insertCell(2);
            apellido.innerHTML = item.apellido;

            let dir = row.insertCell(3);
            dir.innerHTML = item.direccion;

            let estado = row.insertCell(4);
            estado.innerHTML = item.estado;

        });

    } catch (err) {
        console.log(err);
    }

}
GetClient();
