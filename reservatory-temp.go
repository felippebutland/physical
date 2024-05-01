package main

import "fmt"

func main() {
	eficienciaMaxima := 0.30

	tempFria := 300.0

	tempQuente := tempFria / (1 - eficienciaMaxima)

	fmt.Printf("A temperatura do reservatório quente deve ser: %.0f K\n", tempQuente)
}
