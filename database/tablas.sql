CREATE DATABASE ESTANCIAII;
USE ESTANCIAII;

CREATE TABLE USUARIO(
	ID_USUARIO INT auto_increment PRIMARY KEY,
    USERNAME VARCHAR(1000),
    PASS VARCHAR(1000),
    NOMBRE VARCHAR(100),
    APELLIDO VARCHAR(100),
    TIPO_ID INT,
    FOREIGN KEY(TIPO_ID) REFERENCES TIPO_USUARIO(ID)
);

DROP TABLE USUARIO;

CREATE TABLE TIPO_USUARIO(
	ID INT auto_increment PRIMARY KEY,
    TIPO_USUARIO VARCHAR(100)
);

DROP TABLE TIPO_USUARIO;

CREATE TABLE TIPO_ACCESO(
	ID_TIPO INT,
	TABLA VARCHAR(100),
    HORARIO_INICIO timestamp,
    HORARIO_FINAL timestamp,
    FOREIGN KEY(ID_TIPO) REFERENCES TIPO_USUARIO(ID)
);

CREATE TABLE VENTA(
	ID_VENTA INT auto_increment PRIMARY KEY,
    ID_UsuarioVenta INT,
    ID_PRO INT,
    TOTAL FLOAT,
    IVA FLOAT,
    FECHA_VENTA datetime,
	ID_CLI INT,
    FOREIGN KEY(ID_CLI) REFERENCES CLIENTE (ID_CLIENTE),
    FOREIGN KEY(ID_PRO) REFERENCES PRODUCTO (ID_PRODUCTO),
    FOREIGN KEY (ID_UsuarioVenta) REFERENCES USUARIO (ID_USUARIO)
);

DROP TABLE VENTA;

CREATE TABLE PRODUCTO(
	ID_PRODUCTO INT auto_increment PRIMARY KEY,
    PESOxUNIDAD FLOAT,
    PRECIOxUNIDAD FLOAT
);


CREATE TABLE CLIENTE(
	ID_CLIENTE INT auto_increment PRIMARY KEY,
    NOMBRE VARCHAR(100),
    APELLIDO VARCHAR(100),
    DIRECCION VARCHAR(100),
    ESTADO VARCHAR(100)
);
