
  (async function DeleteUser() {

    const url = `http://localhost:8070/reportes/userinfo`;

    try {
        const resp = await axios.get(url)
        
        console.log(resp);
        const ctx = document.getElementById('myChart');

        new Chart(ctx, {
          type: 'bar',
          data: {
            labels: ['Usuarios', 'Tipos de Accesos', 'Accesos a Tablas'],
            datasets: [{
              label: '# de personas o accesos',
              data: [resp.data.cliente, resp.data.tipo_accesso, resp.data.acceso],
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