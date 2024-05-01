package main

import "fmt"

func main() {
	calorQuente := 2.00e3

	calorFrio := 1.50e3

	trabalho := calorQuente - calorFrio

	eficiencia := trabalho / calorQuente

	fmt.Printf("A eficiência da máquina térmica é: %.2f%%\n", eficiencia*100)
}
