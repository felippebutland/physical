package main

import (
	"fmt"
	"math"
)

func main() {
	// Constante α (atm/m^6)
	alfa := 5.0

	volumeInicial := 1.0

	volumeFinal := 2 * volumeInicial

	alfaSI := alfa * 101325 // 1 atm = 101325 Pa

	trabalho := alfaSI / 3 * (math.Pow(volumeFinal, 3) - math.Pow(volumeInicial, 3))

	fmt.Printf("O trabalho realizado no processo é: %e J\n", trabalho)
}
