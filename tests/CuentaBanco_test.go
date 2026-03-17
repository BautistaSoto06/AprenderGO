package tests

import (
	app "practicaGo/src/app"
	"testing"
)

func TestCuentaBanco(t *testing.T) {

	//se crea el objeto
	c := app.CuentaBanco{Saldo: 1000, Titular: "Bautista"}

	//se testear el saldo
	if c.GetSaldo() != 1000 {
		t.Error("El saldo es incorrecto")
	}
	//Testear titular de la cuenta
	if c.GetTitular() != "Bautista" {
		t.Error("El titular es incorrecto")
	}
	//Test de cargar saldo y que realmente se cargue
	c.CargarSaldo(500)

	if c.GetSaldo() != 1500 {
		t.Error("El saldo no se cargo")
	}

	//Test de retirar saldo y que se reste al saldo total
	c.RetirarSaldo(1000)

	if c.GetSaldo() != 500 {
		t.Error("El saldo no se pudo retirar")
	}
	//testear de que en el caso de querer sacar mas de lo que tiene de error
	ok := c.RetirarSaldo(3000)

	if !ok {
		t.Error("El saldo no se pudo retirar")
	}

}
