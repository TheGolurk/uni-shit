const ctx = document.getElementById('myChart');

let myChart = new Chart(ctx, {
    type: 'bar',
    data: {
        labels: ['Total de respaldos'],
        datasets: [{
            label: '# de respaldos hechos',
            data: [0],
            borderWidth: 1
        }]
    },
    options: {
        scales: {
            y: {
                beginAtZero: true
            }
        }
    }
});

async function Reportv3() {

    let who = document.getElementById('who').value;
    let when = document.getElementById('when').value;
    const url = `http://localhost:8070/reportes/backup?who=${who}&where=${when}`;

    try {
        const resp = await axios.get(url)

        myChart.destroy();
        myChart = new Chart(ctx, {
            type: 'bar',
            data: {
                labels: ['Total de respaldos'],
                datasets: [{
                    label: '# de respaldos hechos',
                    data: [resp.data.total],
                    borderWidth: 1
                }]
            },
            options: {
                scales: {
                    y: {
                        beginAtZero: true
                    }
                }
            }
        });

    } catch (err) {
        console.log(err);
        switch (err.response.status) {
            // 409 Duplicado
            case 409:
                showErrMessage('Error', 'No se pudo eliminar ');
                break;
            // 400 el request esta mal
            case 400:
                showErrMessage('Error', 'Error en la información, compruebe de nuevo');
                break;
            // 500 error del backend
            case 500:
                showErrMessage('Error', 'Ocurrió un error inesperado, intente más tarde');
                break;
            default:
                showErrMessage('Error', 'Ocurrió un error inesperado, intente más tarde');
                break;
        }
        
    }

}

function removeData(chart) {
    chart.data.labels.pop();
    chart.data.datasets.forEach((dataset) => {
        dataset.data.pop();
    });
    chart.update();
}

var form = document.getElementById("CreateTypeUser-form");
function handleForm(event) { 
    event.preventDefault(); 
}

form.addEventListener('submit', handleForm);