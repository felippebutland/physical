package main

import "fmt"

func main() {
	tempFria := -3.00 + 273.15 // Converter Celsius para Kelvin
	tempQuente := 27.0 + 273.15

	COP := tempFria / (tempQuente - tempFria)

	fmt.Printf("O COP do refrigerador é: %.0f\n", COP)
}
