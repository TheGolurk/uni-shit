async function UpdateTypeUser() {

    let id = document.getElementById('floatingID').value;
    let type = document.getElementById('floatingType').value;


    const url = 'http://localhost:8070/';
    const body = {

        "id":id,
        "tipo_usuario": type,
    };

    console.log(body);

    try {
        const resp = await axios.post(url, body)

        alert('Creado con exito');

    } catch (err) {
        console.log(err);
        alert('No se pudo crear, valida la informacion');
    }

}


var form = document.getElementById("CreateTypeUser-form");
function handleForm(event) {
    event.preventDefault();
}

form.addEventListener('submit', handleForm);