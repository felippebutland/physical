package main

import (
	"fmt"
)

func main() {
	potencia := 200.0

	comprimentoOnda := 590e-9

	velocidadeLuz := 3e8

	constantePlanck := 6.63e-34

	energiaFoton := constantePlanck * velocidadeLuz / comprimentoOnda

	numeroFotons := potencia / energiaFoton

	fmt.Printf("O número de fótons emitidos por segundo é: %.9e fótons/s\n", numeroFotons)
}
