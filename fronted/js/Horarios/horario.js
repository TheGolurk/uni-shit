async function CreateSchedule() {
    let startTime = document.getElementById('StartTime').value;
    let finalTime = document.getElementById('FinalTime').value;
    let tables = document.getElementById('tablas').value;
    let typeUser = document.getElementById('tipo').value;

    const url = 'http://localhost:8070/accesohora/create';
    const body = {
        "hora_inicio": startTime,
        "hora_final": finalTime,
        "tablas": tables,
        "id_tipo": parseInt(typeUser),
    };

    try {
        const resp = await axios.post(url, body);
        
        alert('Creado correctamente');
        
    } catch (err) {
        console.log(err);
        alert("No se pudo crear");
    }

}


var form = document.getElementById("Schedule-form");
function handleForm(event) {
    event.preventDefault();
}

form.addEventListener('submit', handleForm);