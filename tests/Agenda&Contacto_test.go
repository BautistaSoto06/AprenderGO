package tests

import (
	app "practicaGo/src/app"
	"testing"
)

func TestCreateContacto(t *testing.T) {

	b := app.Contacto{
		Email:    "bautistasoto918@gmail.com",
		Nombre:   "Bautista",
		Telefono: "123456789",
	}

	c := app.Contacto{
		Email:    "bautistasoto918@gmail.com",
		Nombre:   "Bautista",
		Telefono: "123456789",
	}

	a := app.Agenda{}

	if b.GetEmail() != b.Email {
		t.Error("El email no es correcto")
	}

	if b.GetNombre() != b.Nombre {
		t.Error("El nombre no es correcto")
	}

	if b.GetTelefono() != b.Telefono {
		t.Error("El telefono no es correcto")
	}

	if c.GetEmail() != c.Email {
		t.Error("El email no es correcto")
	}

	a.AgregarContacto(b)
	//a.AgregarContacto(c)

	if len(a.Contactos) != 2 {
		t.Error("El contacto no se agrego")
	}
}
