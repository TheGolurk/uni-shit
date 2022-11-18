async function AssignPrivileges() {
    let table = document.getElementById('floatingInput').value;
    let typeUser = document.getElementById('floatingType').value;

    const url = 'http://localhost:8080/user';
    const body = {
        "id_tipo": typeUser,
        "tabla": table
    };

    try {
        const resp = await axios.post(url, body)
        console.log(resp);
    } catch (err) {
        console.log(err);
    }

}

var form = document.getElementById("AssignPrivileges-form");
function handleForm(event) {
    event.preventDefault();
}


form.addEventListener('submit', handleForm);