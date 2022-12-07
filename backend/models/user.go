package models

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Nombre   string `json:"nombre,omitempty"`
	Apellido string `json:"apellido,omitempty"`
	IdTipo   int    `json:"id_tipo,omitempty"`
}

type TYpeUser struct {
	ID          int    `json:"id,omitempty"`
	TipoUsuario string `json:"tipo_usuario,omitempty"`
}

type UserAccess struct {
	ID         int    `json:"id,omitempty"`
	IDTipo     int    `json:"id_tipo,omitempty"`
	HoraInicio string `json:"hora_inicio,omitempty"`
	HoraFinal  string `json:"hora_final,omitempty"`
}

type Cliente struct {
	IdCliente int    `json:"id_cliente,omitempty"`
	Nombre    string `json:"nombre,omitempty"`
	Apellido  string `json:"apellido,omitempty"`
	Direccion string `json:"direccion,omitempty"`
	Estado    string `json:"estado,omitempty"`
}

type Producto struct {
	IdProducto      int     `json:"id_producto,omitempty"`
	PesoProducto    float32 `json:"peso_producto,omitempty"`
	PrecioProduccto float32 `json:"precio_produccto,omitempty"`
}

type Venta struct {
	IdVenta        int     `json:"id_venta,omitempty"`
	IdUsuarioVenta int     `json:"id_usuario_venta,omitempty"`
	IdPro          int     `json:"id_pro,omitempty"`
	Total          float32 `json:"total,omitempty"`
	Iva            float32 `json:"iva,omitempty"`
	FechaVenta     string  `json:"fecha_venta,omitempty"`
	IdCli          int     `json:"id_cli,omitempty"`
}

type Acceso struct {
	ID     int    `json:"id,omitempty"`
	IdTipo int    `json:"id_tipo,omitempty"`
	Tablas string `json:"tablas,omitempty"`
}

type R1 struct {
	Cliente     int `json:"cliente,omitempty"`
	TipoAccesso int `json:"tipo_accesso,omitempty"`
	Acceso      int `json:"acceso,omitempty"`
}

type R2 struct {
	Total float64 `json:"total"`
}

type R3 struct {
	Total int `json:"total,omitempty"`
}
