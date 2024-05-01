package main

import (
	"fmt"
	"math"
)

func main() {
	intensidadeIndividual := 4.0e-7

	intensidadeTotal := 2 * intensidadeIndividual

	intensidadeReferencia := 1.0e-12

	nivelSonoro := 10 * math.Log10(intensidadeTotal/intensidadeReferencia)

	fmt.Printf("O nível sonoro ouvido pelo trabalhador é: %.2f dB\n", nivelSonoro)
}
