package main

import "fmt"

func main() {
	altura := 10.0
	largura := 120.0

	densidadeAgua := 1000.0

	g := 9.8

	pressaoMedia := 0.5 * densidadeAgua * g * altura // Pressão no ponto médio da altura

	area := altura * largura

	forcaTotal := pressaoMedia * area

	fmt.Printf("A força total exercida na parede é: %.2f N\n", forcaTotal)
}
