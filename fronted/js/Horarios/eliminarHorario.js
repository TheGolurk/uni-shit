async function DeleteSchedule() {

    
    let typeUser = document.getElementById('UserType').value;


    const url = `http://localhost:8070/accesohora/delete?=${typeUser};`;
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