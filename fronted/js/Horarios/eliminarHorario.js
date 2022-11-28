async function DeleteSchedule() {
    let startTime = document.getElementById('StartTime').value;
    let finalTime = document.getElementById('FinalTime').value;
    let tables = document.getElementById('tablas').value;
    let typeUser = document.getElementById('UserType').value;


    const url = `http://localhost:8070/accesohora/delete?=${f};`;
    try {
        const resp = await axios.delete(url)
        
        alert('Se eliminó correctamente');

    } catch (err) {
        console.log(err);
        alert('No se pudo eliminar');
    }
}


var form = document.getElementById("Schedule-form");
function handleForm(event) {
    event.preventDefault();
}

form.addEventListener('submit', handleForm);