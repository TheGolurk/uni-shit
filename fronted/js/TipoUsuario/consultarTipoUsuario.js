async function ViewTYpe() {
    const url = 'http://localhost:8070/typeuser/get';

    try {
        const resp = await axios.get(url)
        console.log(resp);

        const table = document.getElementById("tabla-datos");
        resp.data.forEach(item => {
            let row = table.insertRow();

            let id = row.insertCell(0);
            id.innerHTML = item.id;

            let tipo = row.insertCell(1);
            tipo.innerHTML = item.tipo_usuario;
        });

    } catch (err) {
        console.log(err);
    }

}


ViewTYpe();