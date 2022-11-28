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
	IdCliente int    `json:"id_cliente"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Direccion string `json:"direccion"`
	Estado    string `json:"estado"`
}

type Producto struct {
	IdProducto      int     `json:"id_producto"`
	PesoProducto    float32 `json:"peso_producto"`
	PrecioProduccto float32 `json:"precio_produccto"`
}

type Venta struct {
	IdVenta        int     `json:"id_venta"`
	IdUsuarioVenta string  `json:"id_usuario_venta"`
	IdPro          int     `json:"id_pro"`
	Total          float32 `json:"total"`
	Iva            float32 `json:"iva"`
	FechaVenta     string  `json:"fecha_venta"`
	IdCli          int     `json:"id_cli"`
}
