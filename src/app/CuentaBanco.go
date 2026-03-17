package app

type CuentaBanco struct {
	Saldo   float64
	Titular string
}

func NuevaCuentaBanco(Saldo float64, Titular string) CuentaBanco {
	return CuentaBanco{Saldo, Titular}
}

func (c CuentaBanco) GetSaldo() float64 {
	return c.Saldo
}

func (c CuentaBanco) GetTitular() string {
	return c.Titular
}

func (c *CuentaBanco) CargarSaldo(monto float64) {
	if monto > 0 {
		c.Saldo += monto
	}
}

func (c *CuentaBanco) RetirarSaldo(monto float64) bool {
	if monto > 0 && c.GetSaldo() >= monto {
		c.Saldo -= monto
		return true
	}
	return false
}
