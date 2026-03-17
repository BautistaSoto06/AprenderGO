package tests

import (
	app "practicaGo/src/app"
	"testing"
)

func TestPersona(t *testing.T) {

	p := app.Persona{Nombre: "Bautista", Edad: 20, Vocacion: "Ing en sistemas"}

	if p.GetNombre() != "Bautista" {
		t.Error("El nombre es incorrecto")
	}

	if p.GetVocacion() != "Ing en sistemas" {
		t.Error("La Vocacion es Incorrecta")
	}

	if p.GetEdad() != 20 {
		t.Error("El edad es incorrecto")
	}
}
