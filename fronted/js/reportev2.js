const ctx = document.getElementById('myChart');

let myChart = new Chart(ctx, {
    type: 'bar',
    data: {
        labels: ['Total de ventas'],
        datasets: [{
            label: 'Calculo de ventas totales',
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

async function SellReport() {
    let who = document.getElementById('who').value;
    let fecha_inicio = document.getElementById('when').value;
    let fecha_fin = document.getElementById('when2').value;

    const url = `http://localhost:8070/reportes/totasell?username=${ parseInt(who)}&f_inicio=${fecha_inicio}&f_final=${fecha_fin}`;

    try {
        const resp = await axios.get(url)
        
        console.log(resp);

        myChart.destroy();
        myChart = new Chart(ctx, {
            type: 'bar',
            data: {
                labels: ['Total de ventas'],
                datasets: [{
                    label: 'Calculo de ventas totales',
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
    }

}


var form = document.getElementById("CreateTypeUser-form");
function handleForm(event) {
    event.preventDefault();
}

form.addEventListener('submit', handleForm);