package main

import "fmt"

func main() {
	//variaveis
	// var + nome da variavel + tipo
	var nome string
	nome = "Mauricio"
	fmt.Println(nome)

	nome = "Andrade"
	fmt.Println(nome)

	var idade int
	idade = 32;
	fmt.Println(idade)

	var a = "Mauricio"
	fmt.Println(a)

	var b, c int = 1,2
	fmt.Println(b)
	fmt.Println(c)

	var d = true
	fmt.Println(d)

	f := "apple"
	fmt.Println(f)

	//constante, geralmente para campos que não vão mudar o valor
	const cpf = 11111111111
	fmt.Println(cpf)
}