package main

import (
	"fmt"
	"strconv"
)

func main() {
	numero := 58800.00

	notacaoCientifica := strconv.FormatFloat(numero, 'e', -1, 64)

	fmt.Println("Notação científica:", notacaoCientifica)
}
