package main

import "fmt"

func main() {
	massaAgua := 0.200

	tempInicialAgua, tempFinalAgua := 20.0, 50.0

	tempInicialVapor := 130.0

	calorEspecificoAgua, calorEspecificoVapor := 4190.0, 2010.0

	calorLatenteVaporizacao := 2.26e6

	calorParaAquecerAgua := massaAgua * calorEspecificoAgua * (tempFinalAgua - tempInicialAgua)

	massaVapor := calorParaAquecerAgua / (calorLatenteVaporizacao + calorEspecificoVapor*(tempInicialVapor-tempFinalAgua))

	fmt.Printf("A massa de vapor necessária é: %.4f kg\n", massaVapor*1000)
}
