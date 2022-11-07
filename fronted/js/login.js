const administrador = 1;
const personal = 2;

async function ExecuteLogin() {
    let username = document.getElementById('floatingInput').value;
    let password = document.getElementById('floatingPassword').value;

    const url = 'http://localhost:8080/login';
    const body = {
        "username": username,
        "password": password
    };

    try {
        const resp = await axios.post(url, body);

        console.log('set cookie', resp.data);
        setCookie("user-login", resp.data, 2);

        //window.location.href = 'administrador.html'
        console.log(resp);

    } catch (err) {
        console.log(err);
        alert('Usuario o Contraseña invalidos');
    } 

}

function setCookie(name,value,days) {
    var expires = "";
    if (days) {
        var date = new Date();
        date.setTime(date.getTime() + (days*24*60*60*1000));
        expires = "; expires=" + date.toUTCString();
    }
    document.cookie = name + "=" + (value || "")  + expires + "; path=/";
}


var form = document.getElementById("login-form");
function handleForm(event) { 
    event.preventDefault(); 
}

form.addEventListener('submit', handleForm);