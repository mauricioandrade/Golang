package main

import "fmt"
func main() {

	// por padrão o go sempre coloca zerovalues quando não atribuimos valores.
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("Inteiro: %v\n", i)
	fmt.Printf("Float: %v\n", f)
	fmt.Printf("Bool: %v\n", b)
	fmt.Printf("String: %q\n", s)


}