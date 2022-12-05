async function CreateSchedule() {
    let startTime = document.getElementById('StartTime').value;
    let finalTime = document.getElementById('FinalTime').value;
    let typeUser = document.getElementById('tipo').value;

    const url = 'http://localhost:8070/accesohora/create';
    const body = {
        "hora_inicio": startTime,
        "hora_final": finalTime,
        "id_tipo": parseInt(typeUser),
    };

    try {
        const resp = await axios.post(url, body);
        
        showOKMessage('ok','Creado correctamente');
        
    } catch (err) {
        switch (err.response.status) {
            // 409 Duplicado
            case 409:
                showErrMessage('Error', 'Duplicado ');
                break;
            // 400 el request esta mal
            case 400:
                showErrMessage('Error', 'Error en la información, compruebe de nuevo');
                break;
            // 500 error del backend
            case 500:
                showErrMessage('Error', 'Ocurrió un error inesperado, intente más tarde');
                break;
            default:
                showErrMessage('Error', 'Ocurrió un error inesperado, intente más tarde');
                break;
        }
        console.log(err);
    }

}


var form = document.getElementById("Schedule-form");
function handleForm(event) {
    event.preventDefault();
}

form.addEventListener('submit', handleForm);