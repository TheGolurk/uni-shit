INSERT INTO TIPO_USUARIO(TIPO_USUARIO) VALUES ('ADMINISTRADOR');
INSERT INTO TIPO_USUARIO(TIPO_USUARIO)  VALUES ('PERSONAL');
SELECT *FROM TIPO_USUARIO;



SELECT * FROM USUARIO;
INSERT INTO USUARIO( USERNAME, PASS, NOMBRE, APELLIDO, TIPO_ID)
    VALUES('HOLA2', '1234', 'JUANITO', 'SOLAR', 2);
UPDATE USUARIO SET USERNAME= 'JUJU',
                   PASS = '1111',
                   NOMBRE = 'KIKO',
                   APELLIDO = 'POO',
                   TIPO_ID = 1
                   WHERE ID_USUARIO = 3;
DELETE FROM USUARIO WHERE ID_USUARIO = 3;
SELECT *FROM USUARIO;



SELECT * FROM TIPO_ACCESO;
INSERT INTO TIPO_ACCESO(ID_TIPO, TABLA, HORARIO_INICIO, HORARIO_FINAL)
    VALUES(1, 'USUARIO', '7', '13');
INSERT INTO TIPO_ACCESO(ID_TIPO, TABLA, HORARIO_INICIO, HORARIO_FINAL)
    VALUES( 2,'PRODUCTO', '13:00', '12:00');
UPDATE TIPO_ACCESO SET ID_TIPO = 1,
                       TABLA= 'VENTA',
                       HORARIO_INICIO = '12:00',
                       HORARIO_FINAL ='15:00'
                       WHERE ID_TIPO = 2;
DELETE FROM TIPO_ACCESO WHERE ID_TIPO = 1;



SELECT * FROM CLIENTE;
INSERT INTO CLIENTE( NOMBRE, APELLIDO, DIRECCION, ESTADO)
    VALUES('JORGE', 'JUAREZ','AV. LOMAS DE CORTES','MORELOS');

INSERT INTO CLIENTE( NOMBRE, APELLIDO, DIRECCION, ESTADO)
    VALUES('KARLA', 'PEREZ','AV. COCOYOC','PUEBLA');

UPDATE CLIENTE SET NOMBRE = 'BDGD',
                   APELLIDO = 'GDGD',
                   DIRECCION ='FSGDG',
                   ESTADO='GUERRERO'
                   WHERE ID_CLIENTE = 3;
DELETE FROM CLIENTE WHERE ID_CLIENTE = 1;



SELECT * FROM VENTA;
INSERT INTO VENTA(ID_UsuarioVenta, ID_PRO, TOTAL, IVA, FECHA_VENTA, ID_CLI)
    VALUES(4, 1, 23.00, 1.8, '2022-10-02 12:00:09', 2);
UPDATE VENTA  SET ID_PRO = 1,
                  TOTAL = 29,
                  IVA = 1.89,
                  FECHA_VENTA = '2022-09-02 12:00:09',
                  ID_CLI =2
                  WHERE ID_VENTA = 2;
DELETE FROM VENTA WHERE ID_VENTA= 1;




SELECT * FROM PRODUCTO;
INSERT INTO PRODUCTO( PESOxUNIDAD, PRECIOxUNIDAD)
    VALUES (23.89, 30);
UPDATE PRODUCTO SET PESOxUNIDAD = 23.00,
                    PRECIOxUNIDAD = 40.50
                    WHERE ID_PRODUCTO = 1;
DELETE FROM PRODUCTO WHERE ID_PRODUCTO = 1;

