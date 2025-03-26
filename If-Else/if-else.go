package main

import "fmt"

// IF / ELSE
// SE / SENAO


func main() {
	
	// numero := 1
	// // if (condição) { <ação> }
	// if numero == 1 { // true
	// 	fmt.Println("Valor é igual a 1")
	// }else{
	// 	fmt.Println("Valor é diferente de 1")
	// }

	// if numero == 1 {
	// 	fmt.Println("Valor é igual a 1")
	// }else if numero == 3 {
	// 	fmt.Println("Valor é igual a 2")
	// }else{
	// 	fmt.Println("Valor é diferente de 1 e 2")
	// }

	// operadores
	// fmt.Println(1+1)
	// fmt.Println(10-1)
	// fmt.Println(2*2)
	// fmt.Println(10/2)
	// fmt.Println(11%2)

	numero := 8
	if numero%2 == 0 {
		fmt.Printf("%d é par\n", numero)
	}else{
		fmt.Printf("%d é impar\n", numero)
	}

	// nome := "Mauricio"
	// if nome == "Mauricio" {
	// 	fmt.Println("Acertou!\n")
	// }else{
	// 	fmt.Println("Errou!\n")
	// }

	idade := 32
	if idade > 30{
		fmt.Println("Você tem mais de 30 anos")
	}else{
		fmt.Println("Você tem menos de 30 anos")
	}
}
