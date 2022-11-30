async function GetPrivileges() {
    const url = 'http://localhost:8070/accesohora/get';

    try {
        const resp = await axios.get(url)

        const table = document.getElementById("tabla-datos");
        resp.data.forEach(item => {

            let row = table.insertRow();
            let id = row.insertCell(0);
            id.innerHTML = item.id;


            let tipo = row.insertCell(1);
            tipo.innerHTML = item.id_tipo;


            let tables = row.insertCell(2);
            tables.innerHTML = item.tablas;
        });

    } catch (err) {
        console.log(err);
    }

}

GetPrivileges();