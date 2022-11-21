package main

import "fmt"

type qualquer int

var y qualquer

func main() {
	fmt.Printf("%v\n", y)
	fmt.Printf("Tipo %T\n", y)

	x := int(y)
	fmt.Printf("Valor de x: %v\n", x)
	fmt.Printf("Tipo de x: %T", x)
}
