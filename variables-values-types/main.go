package main

import "fmt"

var name = "Luciano"

func main() {
	num := 200
	numf := 12.42
	word := "string"
	boo := true

	fmt.Printf("num type is: %T\n", num)
	fmt.Printf("numf type is: %T\n", numf)
	fmt.Printf("word type is: %T\n", word)
	fmt.Printf("boo type is: %T\n", boo)
	fmt.Println(name)
}