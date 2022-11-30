async function ModifyPrivileges() {
    
    let idAccess = document.getElementById('floatingType1').value;
    let typeUser = document.getElementById('floatingType2').value;
    let table = document.getElementById('floatingInput').value;

    const url = 'http://localhost:8080/accesohora/modify';
    const body = {
        "id": idAccess,
        "id_tipo": typeUser,
        "tabla": table
    };

    try {
        const resp = await axios.put(url, body)
        console.log(resp);
    } catch (err) {
        console.log(err);
    }

}

var form = document.getElementById("UserPrivileges-form");
function handleForm(event) {
    event.preventDefault();
}


form.addEventListener('submit', handleForm);