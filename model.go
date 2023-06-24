package main

type Student struct {
	ID               int    `json:"id"`
	Nombre           string `json:"nombre"`
	Dni              string `json:"dni"`
	Direccion        string `json:"direccion"`
	Fecha_nacimiento string `json:"fecha_nacimiento"`
}

type TipoCliente struct {
	IdTipoCliente int    `json:"idTipoCliente"`
	Nombre        string `json:"nombre"`
	Detalle       string `json:"detalle"`
}

type Cliente struct {
	IdCliente     int    `json:"idCliente"`
	Nombre        string `json:"nombre"`
	Dni           string `json:"dni"`
	Telefono      string `json:"telefono"`
	Correo        string `json:"correo"`
	Estado        string `json:"estado"`
	IdTipoCliente int    `json:"idTipoCliente"`
}
