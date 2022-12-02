async function UpdateTypeUser() {

    let id = document.getElementById('floatingID').value;


    const url = 'http://localhost:8070/';
    const body = {

        "id":id,
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