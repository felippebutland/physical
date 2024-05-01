package main

import (
	"fmt"
	"math"
)

func main() {
	frequencia := 35.00

	velocidadeOnda := 30.0

	frequenciaAngular := 2 * math.Pi * frequencia

	comprimentoOnda := velocidadeOnda / frequencia

	numeroOnda := 2 * math.Pi / comprimentoOnda

	fmt.Printf("Frequência angular (ω): %.2f rad/s\n", frequenciaAngular)
	fmt.Printf("Número de onda (k): %.2f rad/m\n", numeroOnda)
}
