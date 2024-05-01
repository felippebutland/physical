package main

import (
	"fmt"
	"math"
)

func main() {
	comprimentoRepouso := 140.0
	diametroRepouso := 30.0

	velocidadeFracao := 0.98

	fatorLorentz := 1 / math.Sqrt(1-math.Pow(velocidadeFracao, 2))

	comprimentoMovimento := comprimentoRepouso / fatorLorentz

	diametroMovimento := diametroRepouso

	fmt.Printf("Comprimento em movimento: %.2f m\n", comprimentoMovimento)
	fmt.Printf("Diâmetro em movimento: %.2f m\n", diametroMovimento)
}
