async function ExecuteLogin() {
    let username = document.getElementById('floatingInput').value;
    let password = document.getElementById('floatingPassword').value;
    let name = document.getElementById('floatingName').value;
    let lastName = document.getElementById('floatingLastName').value;
    let typeUser = document.getElementById('floatingType').value;


    const url = 'http://localhost:8080/create/user';
    const body = {
        "username": username,
        "password": password,
        "nombre": name,
        "apellido": lastName,
        "id_tipo": typeUser
    };

    try {
        const resp = await axios.post(url, body)
        console.log(resp);
    } catch (err) {
        console.log(err);
    }

}


var form = document.getElementById("CreateUser-form");
function handleForm(event) { 
    event.preventDefault(); 
}

form.addEventListener('submit', handleForm);