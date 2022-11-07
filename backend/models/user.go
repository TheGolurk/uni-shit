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
	HoraInicio string `json:"hora_inicio,omitempty"`
	HoraFinal  string `json:"hora_final,omitempty"`
}
