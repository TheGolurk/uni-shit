async function AssignPrivileges() {
    
    let typeUser = document.getElementById('floatingType').value;
    let table = document.getElementById('floatingInput').value;

    const url = 'http://localhost:8080/accesohora/create';
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