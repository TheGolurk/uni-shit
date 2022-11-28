async function CreateSchedule() {
    let startTime = document.getElementById('StartTime').value;
    let finalTime = document.getElementById('FinalTime').value;
    let tables = document.getElementById('tablas').value;
    let typeUser = document.getElementById('UserType').value;


    const url = 'http://localhost:8070/accesohora/modify';
    const body = {
        "HoraInicio": startTime,
        "HoraFinal": finalTime,
        "Tablas": tables,
        "IDTipo": parseInt(typeUser),
    };

    try {
        const resp = await axios.put(url, body);
        
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