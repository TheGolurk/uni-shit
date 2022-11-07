async function Schedule() {
    let startTime = document.getElementById('floatingStartTime').value;
    let finalTime = document.getElementById('floatingFinalTime').value;
    let typeUser = document.getElementById('floatingType').value;


    const url = 'http://localhost:8080/user/accesohora';
    const body = {
        "horarioInicio": startTime,
        "horarioFinal": finalTime,
        "id_tipo": typeUser

    };

    try {
        const resp = await axios.post(url, body)
        console.log(resp);
    } catch (err) {
        console.log(err);
    }

}


var form = document.getElementById("Schedule-form");
function handleForm(event) {
    event.preventDefault();
}

form.addEventListener('submit', handleForm);