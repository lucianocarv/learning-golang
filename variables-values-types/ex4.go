package main

import "fmt"

type qualquer int

var x qualquer

func main() {
	fmt.Println(x)
	fmt.Printf("Tipo %T\n", x)

	x = 42
	fmt.Println(x)
}
