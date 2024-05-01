package main

import "fmt"

func main() {
	// Pontos fixos da escala S
	pontoGeloS := -25.0
	pontoVaporS := 63.0

	tempCelsius := 30.0

	razaoConversao := (pontoVaporS - pontoGeloS) / 100.0

	tempS := pontoGeloS + razaoConversao*tempCelsius

	fmt.Printf("A temperatura em escala S é: %.1f°C\n", tempS)
}
