package main

import (
	"fmt"
	"math"
)

func main() {
	pesoVeiculo := 23000.0

	raioPequeno := 0.075
	raioGrande := 0.100

	areaPequeno := math.Pi * math.Pow(raioPequeno, 2)
	areaGrande := math.Pi * math.Pow(raioGrande, 2)

	forcaGrande := pesoVeiculo

	pressao := forcaGrande / areaGrande

	forcaPequeno := pressao * areaPequeno

	pressaoAr := forcaPequeno / areaPequeno

	fmt.Printf("A pressão de ar necessária no pistão pequeno é: %.2f N\n", pressaoAr)
}
