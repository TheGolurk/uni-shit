async function ViewSchedule() {
    const url = 'http://localhost:8070/accesohora/get';

    try {
        const resp = await axios.get(url)

        const table = document.getElementById("tabla-datos");
        resp.data.forEach(item => {
            let row = table.insertRow();
            let date = row.insertCell(0);
            date.innerHTML = item.id;

            let a = row.insertCell(1);
            a.innerHTML = item.id_tipo;

            let apellido = row.insertCell(2);
            apellido.innerHTML = item.hora_inicio;

            let tipo = row.insertCell(3);
            tipo.innerHTML = item.hora_final;
        });

    } catch (err) {
        showErrMessage('Error', 'Ocurrió un error inesperado, intente más tarde');
    }

}

ViewSchedule();