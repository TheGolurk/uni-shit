
  (async function DeleteUser() {

    const url = `http://localhost:8070/reportes/totasell`;

    try {
        const resp = await axios.get(url)
        
        console.log(resp);
        const ctx = document.getElementById('myChart');

        new Chart(ctx, {
          type: 'bar',
          data: {
            labels: ['Total de ventas'],
            datasets: [{
              label: '# de ventas totales',
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

})()