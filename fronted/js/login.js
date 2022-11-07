async function ExecuteLogin() {
    let username = document.getElementById('floatingInput').value;
    let password = document.getElementById('floatingPassword').value;

    const url = 'http://localhost:8080/login';
    const body = {
        "username": username,
        "password": password
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