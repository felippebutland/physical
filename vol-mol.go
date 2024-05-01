package main

import "fmt"

func main() {
	pressao := 101000.0

	temperatura := 273.0

	R := 8.314

	// Como temos 1 mol de gás (n=1), a fórmula fica V = RT/P
	volume := R * temperatura / pressao

	fmt.Printf("O volume ocupado por 1 mol do gás é: %.2f L\n", volume*1000)
}
