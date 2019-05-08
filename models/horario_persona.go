package models

type HorarioPersona struct {
	HorarioId *Horario `orm:"column(horario_id);rel(fk)"`
	PersonaId *Persona `orm:"column(persona_id);rel(fk)"`
}
