async function ViewUser() {
    const url = 'http://localhost:8070/user/get';

    try {
        const resp = await axios.get(url)

        const table = document.getElementById("tabla-datos");
        resp.data.forEach(item => {
            let row = table.insertRow();
            let date = row.insertCell(0);
            date.innerHTML = item.username;

            let name = row.insertCell(1);
            name.innerHTML = item.nombre;

            let apellido = row.insertCell(2);
            apellido.innerHTML = item.apellido;

            let tipo = row.insertCell(3);
            tipo.innerHTML = item.id_tipo;
        });

    } catch (err) {
        console.log(err);
    }

}

ViewUser();
