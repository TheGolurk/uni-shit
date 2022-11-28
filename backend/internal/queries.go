package internal

var (

	// TIPO_ACCESO
	Tipoacceso_select = `SELECT * FROM TIPO_ACCESO`

	Tipoacceso_update = `UPDATE TIPO_ACCESO SET 
	TABLA= ?,
	HORARIO_INICIO = ?,
	HORARIO_FINAL = ?
	WHERE ID_TIPO = ?;`

	Tipoacceso_delete = `DELETE FROM TIPO_ACCESO WHERE ID_TIPO = ?;`

	Tipoacceso_create = `
	INSERT INTO TIPO_ACCESO(ID_TIPO, TABLA, HORARIO_INICIO, HORARIO_FINAL)
    VALUES(?, ?, ?, ?);
	`

	//USUARIO
	Usuario_select = `SELECT ID_TIPO, TABLA, HORARIO_INICIO, HORARIO_FINAL FROM TIPO_ACCESO;`

	// CLIENTE
	Cliente_select = `SELECT NOMBRE, APELLIDO, DIRECCION, ESTADO FROM CLIENTE;`

	Cliente_create = `INSERT INTO CLIENTE( NOMBRE, APELLIDO, DIRECCION, ESTADO)
    VALUES(?, ?, ?,?);`

	Cliente_delete = `DELETE FROM CLIENTE WHERE ID_CLIENTE = ?;`

	Cliente_update = `UPDATE CLIENTE SET NOMBRE = ?,
	APELLIDO = ?,
	DIRECCION = ?,
	ESTADO= ?
	WHERE ID_CLIENTE = ?;`

	//VENTA
	Venta_select = `SELECT ID_UsuarioVenta, ID_PRO, TOTAL, IVA, FECHA_VENTA, ID_CLI FROM VENTA;`

	Venta_create = `INSERT INTO VENTA(ID_UsuarioVenta, ID_PRO, TOTAL, IVA, FECHA_VENTA, ID_CLI)
    VALUES(?, ?, ?, ?, ?, ?);`

	Venta_delete = `DELETE FROM VENTA WHERE ID_VENTA = ?;`

	Venta_update = `UPDATE VENTA  SET ID_PRO = ?,
	TOTAL = ?,
	IVA = ?,
	FECHA_VENTA = ?,
	ID_CLI =?
	WHERE ID_VENTA = 2;`

	//PRODUCTO
	Producto_select = `SELECT PESOxUNIDAD, PRECIOxUNIDAD FROM PRODUCTO;`

	Producto_create = `INSERT INTO PRODUCTO( PESOxUNIDAD, PRECIOxUNIDAD)
    VALUES (?, ?);`

	Producto_delete = `DELETE FROM PRODUCTO WHERE ID_PRODUCTO = ?;`

	Producto_update = `UPDATE PRODUCTO SET PESOxUNIDAD = ?,
	PRECIOxUNIDAD = ?
	WHERE ID_PRODUCTO = ?;`
)
