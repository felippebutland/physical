package main

import (
	"fmt"
	"math"
)

func main() {
	potenciaAcustica := 80.0

	distancia := 3.00

	areaEsfera := 4 * math.Pi * math.Pow(distancia, 2)

	intensidade := potenciaAcustica / areaEsfera

	fmt.Printf("A intensidade sonora a 3,00 m da fonte é: %.2f W/m²\n", intensidade)
}
