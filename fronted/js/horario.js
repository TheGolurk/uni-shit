async function Schedule() {
    let startTime = document.getElementById('floatingStartTime').value;
    let finalTime = document.getElementById('floatingFinalTime').value;
    let tables = document.getElementById('tablas').value;
    let typeUser = document.getElementById('floatingType').value;


    const url = 'http://localhost:8080/user/accesohora';
    const body = {
        "hora_inicio": startTime,
        "hora_final": finalTime,
        "id_tipo": parseInt(typeUser),
        "tablas": tables
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