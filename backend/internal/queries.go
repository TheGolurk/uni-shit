package internal

var (

	// TIPO_ACCESO
	Tipoacceso_select = `SELECT * FROM TIPO_ACCESO`

	Tipoacceso_update = `UPDATE TIPO_ACCESO SET 
	HORARIO_INICIO = ?,
	HORARIO_FINAL = ?,
	ID_TIPO = ?
	WHERE ID = ?;`

	Tipoacceso_delete = `DELETE FROM TIPO_ACCESO WHERE ID = ?;`

	Tipoacceso_create = `
	INSERT INTO TIPO_ACCESO(ID_TIPO, HORARIO_INICIO, HORARIO_FINAL)
    VALUES(?, ?, ?);
	`

	//USUARIO
	Usuario_select = `SELECT ID_TIPO, TABLA, HORARIO_INICIO, HORARIO_FINAL FROM TIPO_ACCESO;`

	// CLIENTE
	Cliente_select = `SELECT * FROM CLIENTE;`

	Cliente_create = `INSERT INTO CLIENTE( NOMBRE, APELLIDO, DIRECCION, ESTADO)
    VALUES(?, ?, ?,?);`

	Cliente_delete = `DELETE FROM CLIENTE WHERE ID_CLIENTE = ?;`

	Cliente_update = `UPDATE CLIENTE SET NOMBRE = ?,
	APELLIDO = ?,
	DIRECCION = ?,
	ESTADO= ?
	WHERE ID_CLIENTE = ?;`

	//VENTA
	Venta_select = `SELECT ID_VENTA,ID_UsuarioVenta, ID_PRO, TOTAL, IVA, FECHA_VENTA, ID_CLI FROM VENTA;`

	Venta_create = `INSERT INTO VENTA(ID_UsuarioVenta, ID_PRO, TOTAL, IVA, FECHA_VENTA, ID_CLI)
    VALUES(?, ?, ?, ?, ?, ?);`

	Venta_delete = `DELETE FROM VENTA WHERE ID_VENTA = ?;`

	Venta_update = `UPDATE VENTA SET
                 ID_UsuarioVenta = ?,
    ID_PRO = ?,
	TOTAL = ?,
	IVA = ?,
	FECHA_VENTA = ?,
	ID_CLI =?
	WHERE ID_VENTA = ?;`

	//PRODUCTO
	Producto_select = `SELECT PESOxUNIDAD, PRECIOxUNIDAD FROM PRODUCTO;`

	Producto_create = `INSERT INTO PRODUCTO( PESOxUNIDAD, PRECIOxUNIDAD)
    VALUES (?, ?);`

	Producto_delete = `DELETE FROM PRODUCTO WHERE ID_PRODUCTO = ?;`

	Producto_update = `UPDATE PRODUCTO SET PESOxUNIDAD = ?,
	PRECIOxUNIDAD = ?
	WHERE ID_PRODUCTO = ?;`

	// ACCESO
	Acceso_create = `INSERT INTO ACCESO(ID_TIPO, TABLAS) VALUES(?, ?);`
	Acceso_get    = `SELECT * FROM ACCESO;`
	Acceso_delete = `DELETE FROM ACCESO WHERE ID = ?;`
	Acceso_update = `UPDATE ACCESO SET ID_TIPO = ?, TABLAS = ? WHERE ID = ?;`

	// TIPO USUARIO
	TipoUser_Create = `INSERT INTO TIPO_USUARIO(TIPO_USUARIO) VALUES(?);`
	TipoUser_get    = `SELECT * FROM TIPO_USUARIO;`
	TipoUser_Update = `UPDATE TIPO_USUARIO SET TIPO_USUARIO = ? WHERE ID = ?;`
	TipoUser_Delete = `DELETE FROM TIPO_USUARIO WHERE ID = ?;`

	Report_Client     = `SELECT COUNT(*) FROM USUARIO;`
	Report_TipoAccess = `SELECT COUNT(*) FROM TIPO_ACCESO;`
	Report_Acess      = `SELECT COUNT(*) FROM ACCESO;`

	Report_Sell = `SELECT SUM(TOTAL) FROM VENTA WHERE ID_UsuarioVenta = ? AND FECHA_VENTA BETWEEN ? AND ?;
`

	Report_db = `SELECT COUNT(*) FROM BACKUP_HISTORY WHERE DATE_OF_BACKUP = ? AND WHO = ?;`
)
