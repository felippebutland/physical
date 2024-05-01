package main

import (
	"fmt"
	"math"
)

func main() {
	massaCorda := 0.400

	comprimentoCorda := 4.00

	massaObjeto := 6.00

	g := 9.81

	tensao := massaObjeto * g

	densidadeLinear := massaCorda / comprimentoCorda

	velocidade := math.Sqrt(tensao / densidadeLinear)

	fmt.Printf("A velocidade do pulso na corda é: %.2f m/s\n", velocidade)
}
