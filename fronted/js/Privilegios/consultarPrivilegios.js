async function GetPrivileges() {
    const url = 'http://localhost:8070/accesohora/get';

    try {
        const resp = await axios.get(url)

        const table = document.getElementById("tabla-datos");
        resp.data.forEach(item => {

            let row = table.insertRow();
            let idtipo = row.insertCell(0);
            idtipo.innerHTML = item.id_tipo;

            let tables = row.insertCell(1);
            tables.innerHTML = item.tablas;

            let apellido = row.insertCell(2);
            apellido.innerHTML = item.apellido;

            let tipo = row.insertCell(3);
            tipo.innerHTML = item.id_tipo;
        });

    } catch (err) {
        console.log(err);
    }

}

GetPrivileges();