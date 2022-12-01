async function ViewSchedule() {
    const url = 'http://localhost:8070/accesohora/get';

    try {
        const resp = await axios.get(url)

        const table = document.getElementById("tabla-datos");
        resp.data.forEach(item => {
            let row = table.insertRow();
            let date = row.insertCell(0);
            date.innerHTML = item.id_tipo;


            let apellido = row.insertCell(1);
            apellido.innerHTML = item.hora_inicio;

            let tipo = row.insertCell(2);
            tipo.innerHTML = item.hora_final;
        });

    } catch (err) {
        console.log(err);
    }

}

ViewSchedule();