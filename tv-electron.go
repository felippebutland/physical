package main

import (
	"fmt"
	"math"
)

func main() {
	velocidadeFracao := 0.250

	velocidadeLuz := 299792458.0

	massaRepouso := 0.511e6 * 1.602e-19 / math.Pow(velocidadeLuz, 2)

	velocidade := velocidadeFracao * velocidadeLuz

	fatorLorentz := 1 / math.Sqrt(1-math.Pow(velocidade/velocidadeLuz, 2))

	energiaTotal := fatorLorentz * massaRepouso * math.Pow(velocidadeLuz, 2)

	energiaTotalMeV := energiaTotal / (1.602e-13)

	fmt.Printf("A energia total do elétron é: %.3f MeV\n", energiaTotalMeV)
}
