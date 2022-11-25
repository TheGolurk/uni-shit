package internal

var (

	// TIPO_ACCESO
	tipoacceso_select = `SELECT * FROM TIPO_ACCESO`

	tipoacceso_update = `UPDATE TIPO_ACCESO SET ID_TIPO = ?,
	TABLA= ?,
	HORARIO_INICIO = ?,
	HORARIO_FINAL = ?
	WHERE ID_TIPO = ?;`

	tipoacceso_delete = `DELETE FROM TIPO_ACCESO WHERE ID_TIPO = ?;`

	tipoacceso_create = `
	INSERT INTO TIPO_ACCESO(ID_TIPO, TABLA, HORARIO_INICIO, HORARIO_FINAL)
    VALUES(?, ?, ?, ?);
	`

	//USUARIO
	usuario_select = `SELECT ID_TIPO, TABLA, HORARIO_INICIO, HORARIO_FINAL FROM TIPO_ACCESO;`

	// CLIENTE
	cliente_select = `SELECT NOMBRE, APELLIDO, DIRECCION, ESTADO FROM CLIENTE;`

	cliente_create = `INSERT INTO CLIENTE( NOMBRE, APELLIDO, DIRECCION, ESTADO)
    VALUES(?, ?, ?,?);`

	cliente_delete = `DELETE FROM CLIENTE WHERE ID_CLIENTE = ?;`

	cliente_update = `UPDATE CLIENTE SET NOMBRE = ?,
	APELLIDO = ?,
	DIRECCION = ?,
	ESTADO= ?
	WHERE ID_CLIENTE = ?;`

	//VENTA
	venta_select = `SELECT ID_UsuarioVenta, ID_PRO, TOTAL, IVA, FECHA_VENTA, ID_CLI FROM VENTA;`

	venta_create = `INSERT INTO VENTA(ID_UsuarioVenta, ID_PRO, TOTAL, IVA, FECHA_VENTA, ID_CLI)
    VALUES(?, ?, ?, ?, ?, ?);`

	venta_delete = `DELETE FROM VENTA WHERE ID_VENTA= ?;`

	venta_update = `UPDATE VENTA  SET ID_PRO = ?,
	TOTAL = ?,
	IVA = ?,
	FECHA_VENTA = ?,
	ID_CLI =?
	WHERE ID_VENTA = 2;`

	//PRODUCTO
	producto_select = `SELECT PESOxUNIDAD, PRECIOxUNIDAD FROM PRODUCTO;`

	producto_create = `INSERT INTO PRODUCTO( PESOxUNIDAD, PRECIOxUNIDAD)
    VALUES (?, ?);`

	producto_delete = `DELETE FROM PRODUCTO WHERE ID_PRODUCTO = ?;`

	producto_update = `UPDATE PRODUCTO SET PESOxUNIDAD = ?,
	PRECIOxUNIDAD = ?
	WHERE ID_PRODUCTO = ?;`
)
