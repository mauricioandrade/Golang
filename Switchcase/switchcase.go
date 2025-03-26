package main

import "fmt"

func main() {

	posicao := 2
	switch posicao {
	case 1:
		fmt.Println("Primeiro Lugar")
	case 2:
		fmt.Println("Segundo lugar")
	case 3:
		fmt.Println("Terceiro lugar")
	}

	nome := "Naruto"
	switch nome {
	case "Naruto":
		fmt.Println("É um ninja!")
	case "Mauricio":
		fmt.Println("É um dev!")
	}
}
