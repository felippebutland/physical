package main

import (
	"fmt"
)

func main() {
	comprimento := 1.50
	largura := 1.80
	profundidade := 0.40

	densidadeAgua := 1000.0

	g := 9.8

	volumeAgua := comprimento * largura * profundidade

	massaAgua := volumeAgua * densidadeAgua

	forca := massaAgua * g

	areaBase := comprimento * largura

	pressao := forca / areaBase

	fmt.Printf("A pressão do colchão sobre o assoalho é: %.2f Pa\n", pressao)
}
