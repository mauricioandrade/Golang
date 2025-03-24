package main

import "fmt"

func main() {
	// boolean (true/false)
	fmt.Printf("Type: %T - Value: %v\n", true, true)
	// string (sequencia de bytes geralmente strings vem dentro de aspas "")
	fmt.Printf("Type: %T - Value: %v\n", "Mauricio", "Mauricio")
	// int qualquer numero que esteja sem as aspas
	fmt.Printf("Type: %T - Value: %v\n", 1, 1)
	// float decimal
	fmt.Printf("Type: %T - Value: %v\n", 1.233, 1.233)

}

//Tipos:
// boolean (true/false)
//string (sequencia de bytes em go)
//int (numeros inteiros)
//float (float64/float32) - decimal