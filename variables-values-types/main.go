package main

import "fmt"

var name = "Luciano"
var lastname string // ""
var age int // 0
var weight float64 // 0.0
var b bool // false

type myType int

func main() {	

	age = 20
	num := 200
	numf := 12.42
	word := "string"
	boo := true

	fmt.Printf("num type is: %T\n", num)
	fmt.Printf("numf type is: %T\n", numf)
	fmt.Printf("word type is: %T\n", word)
	fmt.Printf("boo type is: %T\n", boo)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(weight)
	fmt.Println(b)
}