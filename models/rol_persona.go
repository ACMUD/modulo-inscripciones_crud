package models

type RolPersona struct {
	Persona *Persona `orm:"column(persona);rel(fk)"`
	Rol     *Rol     `orm:"column(rol);rel(fk)"`
}
