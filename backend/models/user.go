package models

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Nombre   string `json:"nombre,omitempty"`
	Apellido string `json:"apellido,omitempty"`
	IdTipo   int    `json:"id_tipo,omitempty"`
}

type UserAccess struct {
	IDTipo     int    `json:"id_tipo,omitempty"`
	Tablas     string `json:"tablas,omitempty"`
	HoraInicio string `json:"hora_inicio,omitempty"`
	HoraFinal  string `json:"hora_final,omitempty"`
}

type Cliente struct {
	IdCliente int
	Nombre    string
	Apellido  string
	Direccion string
	Estado    string
}

type Producto struct {
	PesoProducto    float32
	PrecioProduccto float32
}

type Venta struct {
	IdUsuarioVenta string
	IdPro          int
	Total          float32
	Iva            float32
	FechaVenta     string
	IdCli          int
}
