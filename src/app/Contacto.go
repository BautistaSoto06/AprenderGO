package app

type Contacto struct {
	Email    string
	Nombre   string
	Telefono string
}

func NuevoContacto(Email string, Nombre string, Telefono string) Contacto {
	return Contacto{
		Email:    Email,
		Nombre:   Nombre,
		Telefono: Telefono,
	}
}

func (c Contacto) GetNombre() string {
	return c.Nombre
}

func (c Contacto) GetTelefono() string {
	return c.Telefono
}

func (c Contacto) GetEmail() string {
	return c.Email
}
