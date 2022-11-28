async function GetClient() {
    const url = 'http://localhost:8070/cliente/get';

    try {
        const resp = await axios.get(url)
        console.log(resp);

        const table = document.getElementById("tabla-datos");
        resp.data.forEach(item => {
            let row = table.insertRow();
            let nombre = row.insertCell(0);
            nombre.innerHTML = item.nombre;

            let apellido = row.insertCell(1);
            apellido.innerHTML = item.apellido;

            let dir = row.insertCell(2);
            dir.innerHTML = item.direccion;

            let estado = row.insertCell(3);
            estado.innerHTML = item.estado;

        });

    } catch (err) {
        console.log(err);
    }

}
GetClient();
