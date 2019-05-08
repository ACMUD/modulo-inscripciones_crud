package models

type Correo struct {
	TipoCorreo *TipoCorreo `orm:"column(tipo_correo);rel(fk)"`
	Persona    *Persona    `orm:"column(persona);rel(fk)"`
	Correo     string      `orm:"column(correo)"`
}
