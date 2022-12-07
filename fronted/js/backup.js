async function CreateBackup() {

    let cookie = getCookie('user-login');
    let userinfo = JSON.parse(cookie);

    const url = `http://localhost:8070/db/backup?username=${userinfo.username}`;

    try {
        const resp = await axios.post(url)
        
        showOKMessage('Exito', 'Base de datos respaldada');

    } catch (err) {

        switch (err.response.status) {
            // 409 Duplicado
            case 409:
                    showErrMessage('Error', 'Usuario duplicado');
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
        
        console.log(err);
    }
}