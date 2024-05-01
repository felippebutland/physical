package main

import "fmt"

func main() {
	profundidade := 6.00

	densidadeAgua := 1000.0

	g := 9.8

	pressaoAgua := densidadeAgua * g * profundidade

	fmt.Printf("A pressão exercida no tímpano é de aproximadamente: %.2f Pa\n", pressaoAgua)
}
