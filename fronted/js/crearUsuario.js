async function ExecuteLogin() {
    let username = document.getElementById('floatingInput').value;
    let password = document.getElementById('floatingPassword').value;
    let name = document.getElementById('floatingName').value;
    let lastName = document.getElementById('floatingLastName').value;




    const url = 'http://localhost:8080/login';
    const body = {
        "username": username,
        "password": password,
        "nombre": name,
        "apellido": lastName
    };

    try {
        const resp = await axios.post(url, body)
        console.log(resp);
    } catch (err) {
        console.log(err);
    }

}


var form = document.getElementById("login-form");
function handleForm(event) { 
    event.preventDefault(); 
}

form.addEventListener('submit', handleForm);