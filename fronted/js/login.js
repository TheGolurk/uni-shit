const administrador = 1;
const personal = 2;

async function ExecuteLogin() {
    let username = document.getElementById('floatingInput').value;
    let password = document.getElementById('floatingPassword').value;

    const url = 'http://localhost:8070/login';
    const body = {
        "username": username,
        "password": password
    };

    try {
        const resp = await axios.post(url, body);

        console.log(resp.data.id_tipo);

        let urlr = "login.html";
        if (resp.data.id_tipo == 1) {
            urlr = 'administrador.html';
        }else if(resp.data.id_tipo == 2) {
            urlr = 'personal.html';
        }else {
            alert('Tipo de usuario desconocido');
            return;
        }

        setCookie("user-login", JSON.stringify(resp.data), 2);

        window.location = urlr;
        

    } catch (err) {
        console.log(err);
        alert('Usuario o Contraseña invalidos');
    } 

}

function Salir() {
    eraseCookie('user-login');
    window.location = "index.html";
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