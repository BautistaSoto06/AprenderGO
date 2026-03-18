package app

type Agenda struct {
	Contactos []Contacto
}

func (a *Agenda) AgregarContacto(c Contacto) {
	a.Contactos = append(a.Contactos, c)
}

func (a Agenda) GetContactos() []Contacto {
	return a.Contactos
}
