package app

type Persona struct {
	Nombre   string
	Edad     int
	Vocacion string
}

func NewPersona(Nombre string, Edad int, vocacion string) *Persona {
	return &Persona{Nombre, Edad, vocacion}
}

func (p Persona) GetNombre() string {
	return p.Nombre
}

func (p Persona) GetEdad() int {
	return p.Edad
}

func (p Persona) GetVocacion() string {
	return p.Vocacion
}
